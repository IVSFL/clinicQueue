import React from "react";

const SuccessButton = ({ children }) => {
  return (
    <button style={{background: "#3aa8a8"}} type="submit" className="btn btn-success w-100 mt-3">
      {children}
    </button>
  );
};

export default SuccessButton;
