import React from "react";
import "./style/PatientQueueStyle.css";

const PatientQueue = ({ queue, callPatient, callNextPatient }) => {
  return (
    <div className="col-lg-4">
      <div className="d-flex justify-content-end mb-3">
        <button className="btn btn-green w-100" onClick={callNextPatient}>
          Следующий
        </button>
      </div>

      <div>
        {queue.map((item) => (
          <div
            key={item.id}
            className="next-patient-card d-flex justify-content-between align-items-center"
          >
            <span>
              {item.patient.last_name} {item.patient.first_name}
            </span>
            <button
              className="btn btn-sm btn-green"
              onClick={() => callPatient(item)}
            >
              Вызвать
            </button>
          </div>
        ))}
      </div>
    </div>
  );
};

export default PatientQueue;
