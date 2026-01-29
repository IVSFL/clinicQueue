import React from "react";

const ReceptionAccount = ({ children }) => {
  return (
    <div className="container py-5" style={{ maxWidth: "800px" }}>
      <div className="card shadow border-0">
        <div
          className="card-header text-white text-center"
          style={{ background: "#3aa8a8" }}
        >
          <h4>Данные пациента</h4>
        </div>
        <div className="card-body">{children}</div>
      </div>
    </div>
  );
};

export default ReceptionAccount;