import React from "react";

const FormTextarea = ({ label, value, placeholder, onChange, name }) => {
  return (
    <div className="mb-3">
      <label className="form-label">{label}</label>
      <textarea
        rows="4"
        value={value}
        placeholder={placeholder}
        name={name}
        onChange={onChange}
      ></textarea>
    </div>
  );
};

export default FormTextarea;
