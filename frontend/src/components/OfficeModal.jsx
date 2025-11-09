import React, { useState, useEffect } from "react";

const OfficeModal = ({ doctorId, onOfficeSet, show, onClose }) => {
  const [office, setOffice] = useState("");
  const [currentOffice, setCurrentOffice] = useState("");

  useEffect(() => {
    if (show) {
      const loadCurrentOffice = async () => {
        try {
          const res = await fetch(`http://localhost:8000/doctors/${doctorId}`);
          const doctor = await res.json();
          setCurrentOffice(doctor.office || "");
          setOffice(doctor.office || "");
        } catch (err) {
          console.error("Ошибка загрузки данных врача:", err);
        }
      };
      loadCurrentOffice();
    }
  }, [show, doctorId]);

  const handleSubmit = async (e) => {
    e.preventDefault();
    
    if (!office.trim()) {
      alert("Введите номер кабинета");
      return;
    }

    try {
      const response = await fetch(`http://localhost:8000/doctors/${doctorId}/office`, {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ office: office.trim() }),
      });

      if (response.ok) {
        onOfficeSet(office.trim());
      } else {
        alert("Ошибка при обновлении кабинета");
      }
    } catch (err) {
      console.error("Ошибка:", err);
      alert("Не удалось обновить кабинет");
    }
  };

  if (!show) return null;

  return (
    <div className="modal-overlay office-modal-overlay">
      <div className="modal-content office-modal-content">
        <div className="modal-header">
          <h3>Выберите кабинет</h3>
          <button type="button" className="close-btn" onClick={onClose}>
            ×
          </button>
        </div>
        <form onSubmit={handleSubmit}>
          <div className="modal-body">
            {currentOffice && (
              <div className="alert alert-info mb-3">
                <strong>Предыдущий кабинет:</strong> {currentOffice}
              </div>
            )}
            <div className="mb-3">
              <label htmlFor="office" className="form-label">
                Номер кабинета:
              </label>
              <input
                type="text"
                className="form-control"
                id="office"
                value={office}
                onChange={(e) => setOffice(e.target.value)}
                placeholder="Например: 101, 205, 3А"
                required
                autoFocus
              />
              <div className="form-text">
                Укажите кабинет, в котором вы работаете сегодня
              </div>
            </div>
          </div>
          <div className="modal-footer">
            <button type="button" className="btn btn-secondary" onClick={onClose}>
              Отмена
            </button>
            <button type="submit" className="btn btn-primary">
              Сохранить
            </button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default OfficeModal;