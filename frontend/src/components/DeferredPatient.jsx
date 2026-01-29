import React from "react";

const DeferredPatients = ({ patients, callPatient }) => {
  return (
    <div className="card">
      <div className="card-header bg-warning text-dark">
        <h5 className="mb-0">Отложенные пациенты</h5>
      </div>
      <div className="card-body">
        {!patients || patients.length === 0 ? (
          <p className="text-muted">Нет отложенных пациентов</p>
        ) : (
          <div className="list-group">
            {patients.map((ticket) => (
              <div key={ticket.ticket_number} className="list-group-item">
                <div className="d-flex justify-content-between align-items-center">
                  <div className="flex-grow-1">
                    <h6 className="mb-1">
                      {ticket.patient?.last_name} {ticket.patient?.first_name} {ticket.patient?.middle_name}
                    </h6>
                    <p className="mb-1 text-muted small">
                      Талон: <strong>{ticket.ticket_number}</strong>
                    </p>
                    {ticket.patient?.content && (
                      <p className="mb-1 small">{ticket.patient.content}</p>
                    )}
                    <p className="mb-0 text-muted small">
                      Отложен: {ticket.called_at ? new Date(ticket.called_at).toLocaleTimeString() : 'Не указано'}
                    </p>
                  </div>
                  <button
                    className="btn btn-success btn-sm ms-3"
                    onClick={() => callPatient(ticket)}
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

export default DeferredPatients;