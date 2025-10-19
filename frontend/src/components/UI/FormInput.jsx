import React from "react";

const FormInput = ({ inputType, id, name, placeholder, value, onChange }) => {
  return (
    <input
      className="form-control"
      type={inputType}
      id={id}
      name={name}
      placeholder={placeholder}
      value={value}
      onChange={onChange}
      autoComplete="off"
    />
  );
};

export default FormInput;
