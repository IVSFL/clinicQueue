import React from "react";

const FormLabel = ({ labelFor, labelText }) => {
  return (
    <label htmlFor={labelFor} className="form-label">
      {labelText}
    </label>
  );
};

export default FormLabel;
