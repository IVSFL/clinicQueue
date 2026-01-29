import React, { useEffect, useState } from "react";
import ReceptionAccount from "./ReceptionAccount";
import DivForm from "./DivForm";
import SuccessButton from "./UI/SuccessButton";
import FormSelect from "./UI/FormSelect";
import FormTextarea from "./UI/FormTextaera";
import ModalWindow from "./UI/ModalWindow/ModalWindow";
import DoctorAccount from "./DoctorAccount";
import axios from "axios";

const Account = () => {
  const [role, setRole] = useState(""); // для хранения роли
  const [ticket, setTicket] = useState('');
  const [specialization, setSpecialization] = useState([]);
  const [selectedSpecialization, setSelectedSpecialization] = useState("");
  const [modal, setModal] = useState(false);

  const [credentials, setCredentials] = useState({
    lastName: "",
    firstName: "",
    middleName: "",
    birthDate: "",
    phoneNumber: "",
    passportNumber: "",
    policyOMS: "",
    content: "",
  });

  useEffect(() => {
    // Получаем роль из localStorage
    const storedRole = localStorage.getItem("role");
    setRole(storedRole);

    if (storedRole === "admin") {
      fetchSpecialization();
      fetchPatient();
    }
  }, []);

  async function fetchSpecialization() {
    const res = await axios.get("http://localhost:8000/specialization");
    setSpecialization(res.data);
  }

  async function fetchPatient() {
    const res = await axios.get("http://localhost:8000/patients");
    console.log(res.data);
  }

  const handleChange = (e) => {
    const { name, value } = e.target;
    setCredentials((prev) => ({ ...prev, [name]: value }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const res = await axios.post("http://localhost:8000/patients", {
        last_name: credentials.lastName,
        first_name: credentials.firstName,
        middle_name: credentials.middleName,
        birth_date: credentials.birthDate,
        phone_number: credentials.phoneNumber,
        passport_number: credentials.passportNumber,
        policy_oms: credentials.policyOMS,
        content: credentials.content,
      });

      const resTicket = await axios.post("http://localhost:8000/tickets", {
        patient_id: res.data.id,
        specialization: selectedSpecialization,
      });

      setTicket(resTicket.data.ticket_number);
      setModal(true);

      // Сброс формы
      setCredentials({
        lastName: "",
        firstName: "",
        middleName: "",
        birthDate: "",
        phoneNumber: "",
        passportNumber: "",
        policyOMS: "",
        content: "",
      });
      setSelectedSpecialization("");
    } catch (err) {
      alert("Ошибка при регистрации");
      console.error("ERROR: ", err);
    }
  };

  return (
    <div style={{ padding: "2rem" }}>
      {role === "admin" && (
        <>
          <ModalWindow visible={modal} setVisible={setModal}>
            <h1>Талон номер: {ticket}</h1>
          </ModalWindow>

          <ReceptionAccount>
            <div className="mb-4 d-flex justify-content-end">
              <SuccessButton>Найти пациента</SuccessButton>
            </div>

            <form onSubmit={handleSubmit}>
              <div className="row g-3">
                <div className="col-md-4">
                  <DivForm
                    labelFor="lastName"
                    labelText="Фамилия"
                    inputType="text"
                    name="lastName"
                    placeholder="Введите фамилию"
                    value={credentials.lastName}
                    onChange={handleChange}
                  />
                </div>
                <div className="col-md-4">
                  <DivForm
                    labelFor="firstName"
                    labelText="Имя"
                    inputType="text"
                    name="firstName"
                    placeholder="Введите имя"
                    value={credentials.firstName}
                    onChange={handleChange}
                  />
                </div>
                <div className="col-md-4">
                  <DivForm
                    labelFor="middleName"
                    labelText="Отчество"
                    inputType="text"
                    name="middleName"
                    placeholder="Введите отчество"
                    value={credentials.middleName}
                    onChange={handleChange}
                  />
                </div>
              </div>

              <div className="row g-3 mt-2">
                <div className="col-md-4">
                  <DivForm
                    labelFor="birthDate"
                    labelText="Дата рождения"
                    inputType="date"
                    name="birthDate"
                    value={credentials.birthDate}
                    onChange={handleChange}
                  />
                </div>
                <div className="col-md-4">
                  <DivForm
                    labelFor="phoneNumber"
                    labelText="Номер телефона"
                    inputType="text"
                    name="phoneNumber"
                    placeholder="+7(999)999-99-99"
                    value={credentials.phoneNumber}
                    onChange={handleChange}
                  />
                </div>
                <div className="col-md-4">
                  <DivForm
                    labelFor="passportNumber"
                    labelText="Паспортные данные"
                    inputType="text"
                    name="passportNumber"
                    placeholder="1111 111111"
                    value={credentials.passportNumber}
                    onChange={handleChange}
                  />
                </div>
              </div>

              <div className="row g-3 mt-2">
                <div className="col-md-6">
                  <DivForm
                    labelFor="policyOMS"
                    labelText="Полис ОМС"
                    inputType="text"
                    name="policyOMS"
                    placeholder="Введите номер полиса"
                    value={credentials.policyOMS}
                    onChange={handleChange}
                  />
                </div>
                <div className="col-md-6">
                  <FormSelect
                    label="Специалист"
                    specialization={specialization}
                    selected={selectedSpecialization}
                    onChange={setSelectedSpecialization}
                  />
                </div>
                <div>
                  <FormTextarea
                    label="Жалобы"
                    name="content"
                    value={credentials.content}
                    placeholder="Введите жалобы"
                    onChange={handleChange}
                  />
                </div>
              </div>

              <div className="d-flex justify-content-end mt-4">
                <SuccessButton>Выдать талон</SuccessButton>
              </div>
            </form>
          </ReceptionAccount>
        </>
      )}

      {role === "doctor" && <DoctorAccount />}
    </div>
  );
};

export default Account;
