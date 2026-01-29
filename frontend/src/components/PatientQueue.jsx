import React from "react";

const PatientQueue = ({ queue, callPatient, callNextPatient }) => {
  return (
    <div className="card h-100"> {/* Добавляем h-100 для заполнения высоты */}
      <div className="card-header bg-info text-white">
        <div className="d-flex justify-content-between align-items-center">
          <h5 className="mb-0">Очередь пациентов</h5>
          <button 
            className="btn btn-light btn-sm"
            onClick={callNextPatient}
          >
            Следующий
          </button>
        </div>
      </div>
      <div className="card-body">
        {queue.length === 0 ? (
          <p className="text-muted">Очередь пуста</p>
        ) : (
          <div className="list-group">
            {queue.map((queueItem) => (
              <div key={queueItem.id} className="list-group-item">
                <div className="d-flex justify-content-between align-items-center">
                  <div className="flex-grow-1">
                    <h6 className="mb-1">
                      {queueItem.patient?.last_name} {queueItem.patient?.first_name} {queueItem.patient?.middle_name}
                    </h6>
                    <p className="mb-1 text-muted small">
                      Талон: <strong>{queueItem.ticket?.ticket_number}</strong>
                    </p>
                    {queueItem.patient?.content && (
                      <p className="mb-1 small">{queueItem.patient.content}</p>
                    )}
                  </div>
                  <button
                    className="btn btn-primary btn-sm ms-3"
                    onClick={() => callPatient(queueItem)}
                  >
                    Вызвать
                  </button>
                </div>
              </div>
            ))}
          </div>
        )}
      </div>
    </div>
  );
};

export default PatientQueue;