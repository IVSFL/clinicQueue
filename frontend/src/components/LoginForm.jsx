import React, { useState } from "react";
import DivForm from "./DivForm";
import SuccessButton from "./UI/SuccessButton";
import axios from 'axios'
import { useNavigate } from "react-router-dom";

const LoginForm = () => {
  const navigate = useNavigate();

  // локальное состояние для email и password
  const [credentials, setCredentials] = useState({
    email: "",
    password: "",
  });

  // обновление состояния при изменении поля
  const handleChange = (e) => {
    const { name, value } = e.target;
    setCredentials((prev) => ({ ...prev, [name]: value }));
  };

  // обработчик входа
  const handleSubmit = async (e) => {
    e.preventDefault();

    if (credentials.email && credentials.password) {
      try{
        const res = await axios.post('http://localhost:8000/register/login',
          {
            email: credentials.email,
            password: credentials.password,
          }
        )
        localStorage.setItem("token", res.data.token);
        localStorage.setItem("user", JSON.stringify(res.data.user));
        console.log(res.data.token);
        console.log(res.data.user)

        navigate("/account")
      } catch(err) {
        console.error('ERROR!', err);
      }
    } else {
      alert("Введите логин и пароль");
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <DivForm
        labelFor="email"
        labelText="Логин"
        inputType="email"
        id="email"
        name="email"
        placeholder="Введите логин"
        value={credentials.email}
        onChange={handleChange}
      />

      <DivForm
        labelFor="password"
        labelText="Пароль"
        inputType="password"
        id="password"
        name="password"
        placeholder="Введите пароль"
        value={credentials.password}
        onChange={handleChange}
      />

      <SuccessButton>Войти</SuccessButton>
    </form>
  );
};

export default LoginForm;
