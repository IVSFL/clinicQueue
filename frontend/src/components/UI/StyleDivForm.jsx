import React from "react";

const StyleDivForm = ({ children }) => {
  return (
    <div
      className="d-flex justify-content-center align-items-center"
      style={{ height: "100vh", backgroundColor: "#f9fafb" }}
    >
      <div
        className="card shadow p-4"
        style={{ width: "400px", borderRadius: "10px" }}
      >
        <h3 className="text-center mb-4">Вход в аккаунт</h3>
        {children}
      </div>
    </div>
  );
};

export default StyleDivForm;
