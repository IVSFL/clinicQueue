import React from "react";

const PatientInfo = ({ currentPatient, deferPatient, completePatient, transferPatient }) => {
  if (!currentPatient) {
    return (
      <div className="card" style={{ minWidth: '300px' }}>
        <div className="card-header bg-primary text-white">
          <h5 className="mb-0">Текущий пациент</h5>
        </div>
        <div className="card-body text-center">
          <p className="text-muted">Нет активного пациента</p>
        </div>
      </div>
    );
  }

  return (
    <div className="card" style={{ minWidth: '300px' }}>
      <div className="card-header bg-primary text-white">
        <h5 className="mb-0">Текущий пациент</h5>
      </div>
      <div className="card-body">
        <div className="mb-4">
          <h6 className="card-title text-dark fw-bold">
            {currentPatient.name}
          </h6>
          <div className="patient-details">
            <div className="detail-item mb-2">
              <span className="text-muted small">Талон:</span>
              <strong className="ms-2 text-dark">{currentPatient.ticketNumber}</strong>
            </div>
            {currentPatient.content && (
              <div className="detail-item mb-2">
                <span className="text-muted small">Жалобы:</span>
                <span className="ms-2 small">{currentPatient.content}</span>
              </div>
            )}
            <div className="detail-item">
              <span className="text-muted small">Кабинет:</span>
              <strong className="ms-2 text-dark">{currentPatient.office}</strong>
            </div>
          </div>
        </div>
        
        <div className="action-buttons">
          <div className="d-grid gap-2">
            <button
              className="btn btn-success btn-action"
              onClick={() => completePatient(currentPatient.ticketNumber)}
            >
              <i className="bi bi-check-circle me-2"></i>
              Завершить прием
            </button>
            
            <button
              className="btn btn-info btn-action"
              onClick={() => transferPatient(currentPatient.ticketNumber)}
            >
              <i className="bi bi-arrow-left-right me-2"></i>
              Передать врачу
            </button>
            
            <button
              className="btn btn-warning btn-action"
              onClick={() => deferPatient(currentPatient.ticketNumber)}
            >
              <i className="bi bi-clock me-2"></i>
              Отложить
            </button>
          </div>
        </div>
      </div>
    </div>
  );
};

export default PatientInfo;