import React, { useState } from "react";
import axios from "axios";

const TransferPatientModal = ({ show, onClose, ticketNumber, patient }) => {
  const [selectedDoctorSpecialization, setSelectedDoctorSpecialization] = useState("");
  const [loading, setLoading] = useState(false);

  // 🔹 Список врачей только по специализации
  const doctors = [
    "Терапевт",
    "Хирург",
    "Кардиолог",
    "Невролог",
  ];

  const handleTransfer = async () => {
    if (!selectedDoctorSpecialization) {
      alert("Выберите врача для передачи");
      return;
    }

    setLoading(true);
    try {
      await axios.post(
        `http://localhost:8000/queue/transfer/${ticketNumber}`,
        {
          // сервер сам найдет врача по specialization
          new_doctor_specialization: selectedDoctorSpecialization
        }
      );

      alert("✅ Пациент успешно передан");
      onClose(true);
    } catch (err) {
      console.error(err);
      alert("❌ Ошибка при передаче пациента: " + (err.response?.data?.error || err.message));
    } finally {
      setLoading(false);
    }
  };

  if (!show) return null;

  return (
    <div className="modal-overlay transfer-modal-overlay">
      <div className="modal-content transfer-modal-content">
        <div className="modal-header bg-info text-white">
          <h5 className="modal-title">Передача пациента</h5>
          <button
            type="button"
            className="btn-close btn-close-white"
            onClick={() => onClose(false)}
          />
        </div>

        <div className="modal-body p-4">
          {/* 🔹 Информация о пациенте/талоне */}
          <div className="ticket-info mb-3">
            <strong>Пациент:</strong> {patient?.last_name} {patient?.first_name} <br/>
            <strong>Номер талона:</strong> {ticketNumber} <br/>
            <strong>Специализация талона:</strong> {patient?.specialization_name || "не указана"}
          </div>

          {/* 🔹 Выбор врача по специализации */}
          <div className="mb-3">
            <label className="form-label fw-semibold mb-2">
              Выберите специализацию врача для передачи
            </label>
            <select
              className="form-select"
              value={selectedDoctorSpecialization}
              onChange={(e) => setSelectedDoctorSpecialization(e.target.value)}
            >
              <option value="">-- Выберите специализацию --</option>
              {doctors.map((spec) => (
                <option key={spec} value={spec}>
                  {spec}
                </option>
              ))}
            </select>
          </div>
        </div>

        <div className="modal-footer bg-light">
          <button
            type="button"
            className="btn btn-outline-secondary"
            onClick={() => onClose(false)}
          >
            Отмена
          </button>
          <button
            type="button"
            className="btn btn-success"
            disabled={!selectedDoctorSpecialization || loading}
            onClick={handleTransfer}
          >
            {loading ? "Передача..." : "Передать пациента"}
          </button>
        </div>
      </div>
    </div>
  );
};

export default TransferPatientModal;
