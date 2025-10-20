import React, { useEffect, useState } from "react";
import "./style/CallBoxStyle.css";

const CallBox = () => {
  const [calls, setCalls] = useState([]);

  useEffect(() => {
    const ws = new WebSocket("ws://127.0.0.1/ws");
    ws.onopen = () => console.log("WEBSOKET CONNECTED");
    ws.onmessage = (event) => {
      const data = JSON.parse(event.data);
      console.log("new: ", data);

      setCalls((prev) => [...prev, data]);

      alert(
        `Пациент ${data.patient.last_name}, вызван в кабинет ${data.office}`
      );
    };
    ws.onclose = () => console.log("WEBSOKEN DISCONNECT");
    return () => ws.close();
  }, []);

  return (
    <div className="wrap">
      <div className="columns-head" aria-hidden="true">
        <div>Пациент</div>
        <div>Кабинет</div>
        <div>Специалист</div>
      </div>

      <section className="rows" aria-live="polite">
        {calls.map((call, idx) => (
          <div className="row" key={idx}>
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
  );
};

export default CallBox;
