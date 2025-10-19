import React from "react";

const FormSelect = ({ label, specialization, selected, onChange }) => {
  return (
    <div className="mb-3">
      {label && <label className="form-label">{label}</label>}
      <select
        className="form-select"
        value={selected}
        onChange={(e) => onChange(e.target.value)}
      >
        <option value="">Выберите специалиста</option>
        {specialization.map((spec) => (
          <option key={spec.id} value={spec.name}>
            {spec.name}
          </option>
        ))}
      </select>
    </div>
  );
};

export default FormSelect;
