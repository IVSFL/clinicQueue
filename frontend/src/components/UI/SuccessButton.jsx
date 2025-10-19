import React from "react";

const SuccessButton = ({ children }) => {
  return (
    <button type="submit" className="btn btn-success w-100 mt-3">
      {children}
    </button>
  );
};

export default SuccessButton;
