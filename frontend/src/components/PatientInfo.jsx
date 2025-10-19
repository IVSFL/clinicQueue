import React from "react";
import "./style/PatientInfoStyle.css";

const PatientInfo = ({ currentPatient }) => {
  if (!currentPatient) return <div className="co-lg-3">Пациент не вызван</div>;
  return (
    <div className="col-lg-8 mb-4">
      <div className="patient-card">
        <h3>
          Пациент: {currentPatient.name}
        </h3>
        <p><strong>Номер талона:</strong> {currentPatient.ticketNumber}</p>
        <p>
          <strong>Жалобы:</strong> {currentPatient.content}
        </p>

        <div className="d-flex justify-content-end mt-4">
          <button className="btn btn-outline-success">Отложить</button>
        </div>
      </div>
    </div>
  );
};

export default PatientInfo;
