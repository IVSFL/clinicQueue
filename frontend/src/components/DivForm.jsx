import React from "react";
import FormInput from "./UI/FormInput";
import FormLabel from "./UI/FormLabel";

const DivForm = ({
  labelFor,
  labelText,
  inputType,
  id,
  name,
  placeholder,
  value,
  onChange,
}) => {
  return (
    <div className="mb-3">
      <FormLabel labelFor={labelFor} labelText={labelText} />
      <FormInput
        inputType={inputType}
        id={id}
        name={name}
        placeholder={placeholder}
        value={value}
        onChange={onChange}
      />
    </div>
  );
};

export default DivForm;
