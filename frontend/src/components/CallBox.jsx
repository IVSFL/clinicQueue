import React, { useEffect, useState } from "react";
import "./style/CallBoxStyle.css"

const CallBox = () => {
  const [calls, setCalls] = useState([]);
  const [showModal, setShowModal] = useState(false);
  const [currentCall, setCurrentCall] = useState(null);

  // Функция для простого beep звука
  const playBeep = () => {
    try {
      const audioContext = new (window.AudioContext || window.webkitAudioContext)();
      const oscillator = audioContext.createOscillator();
      const gainNode = audioContext.createGain();
      
      oscillator.connect(gainNode);
      gainNode.connect(audioContext.destination);
      
      oscillator.frequency.value = 800; // Частота в герцах
      oscillator.type = 'sine'; // Тип волны
      
      gainNode.gain.setValueAtTime(0.3, audioContext.currentTime);
      gainNode.gain.exponentialRampToValueAtTime(0.01, audioContext.currentTime + 0.5);
      
      oscillator.start(audioContext.currentTime);
      oscillator.stop(audioContext.currentTime + 0.5);
    } catch (e) {
      console.log("Web Audio API not supported:", e);
    }
  };

  useEffect(() => {
    const ws = new WebSocket("ws://127.0.0.1:8000/ws");
    ws.onopen = () => console.log("WEBSOCKET CONNECTED");
    ws.onmessage = (event) => {
      const data = JSON.parse(event.data);
      console.log("new: ", data);

      // Удаляем предыдущие вызовы этого врача
      setCalls((prev) => {
        const filteredCalls = prev.filter(call => call.doctor.id !== data.doctor.id);
        return [...filteredCalls, data];
      });

      // Показываем модальное окно и воспроизводим звук
      setCurrentCall(data);
      setShowModal(true);
      playBeep();

      // Автоматически закрываем модальное окно через 5 секунд
      setTimeout(() => {
        setShowModal(false);
      }, 5000);
    };
    ws.onclose = () => console.log("WEBSOCKET DISCONNECT");
    return () => ws.close();
  }, []);

  const closeModal = () => {
    setShowModal(false);
  };

  return (
    <>
      {/* Модальное окно */}
      {showModal && currentCall && (
        <div className="modal-overlay" onClick={closeModal}>
          <div className="modal-content" onClick={(e) => e.stopPropagation()}>
            <div className="modal-header">
              <h3>Вызов пациента</h3>
              <button className="close-btn" onClick={closeModal}>×</button>
            </div>
            <div className="modal-body">
              <div className="call-info">
                <div className="patient-name">
                  {currentCall.patient.last_name} {currentCall.patient.first_name} {currentCall.patient.middle_name}
                </div>
                <div className="ticket-number">
                  Талон: <strong>{currentCall.ticketNumber}</strong>
                </div>
                <div className="office">
                  Кабинет: <strong>{currentCall.office}</strong>
                </div>
                <div className="doctor">
                  Врач: {currentCall.doctor.last_name} {currentCall.doctor.first_name} {currentCall.doctor.middle_name}
                </div>
              </div>
            </div>
            <div className="modal-footer">
              <button className="btn-ok" onClick={closeModal}>OK</button>
            </div>
          </div>
        </div>
      )}

      {/* Основное окно вызовов */}
      <div className="wrap">
        <div className="columnsHead" aria-hidden="true">
          <div>Пациент</div>
          <div>Кабинет</div>
          <div>Специалист</div>
        </div>

        <section className="rows" aria-live="polite">
          {calls.map((call, idx) => (
            <div className="row" key={`${call.doctor.id}-${idx}`}>
              <div className="ticket">{call.ticketNumber}</div>
              <div className="middle">
                <span className="arrow">→</span>
                <span className="room">{call.office}</span>
              </div>
              <div className="doctor">
                {call.doctor.last_name} {call.doctor.first_name} <br />
                {call.doctor.middle_name}
              </div>
            </div>
          ))}
        </section>
      </div>
    </>
  );
};

export default CallBox;