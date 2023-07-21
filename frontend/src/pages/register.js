import React from "react";
import styled from 'styled-components'
import { registerReq } from "../api/register";
import { useState } from "react";
import { Navigate } from "react-router-dom";
import { useNavigate } from "react-router-dom";

const RegisterPage = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100vh;
`;

const RegisterForm = styled.form`
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 1rem;
  padding: 1rem;
  border: 1px solid #ccc;
  border-radius: 5px;
  max-width: 300px;
`;

const Register = () => {
    let history = useNavigate()

    const [formData, setFormData] = useState({username: '', password: ''})
    const [error, setError] = useState('')
    const [isRegister, setIsRegister] = useState(false)

    const handleLoginClick = () => {
        history("/login", {replace: true})
    }

    const handleInputChange = (event) => {
        const { name, value } = event.target;
        setFormData({ ...formData, [name]: value });
    };

    const handleRegisterSubmit = async (e) => {
        e.preventDefault()
        try {
            await registerReq(formData)
            setError('')
            setIsRegister(true)
        } catch(error) {
            setError(error.message);
            console.error('Register error:', error);
        }
    }

    if (isRegister) {
        return <Navigate to="/login" />;
    }

    return (
        <RegisterPage>
            <h1>Register</h1>
            {error && <p style={{ color: 'red' }}>{error}</p>}
            <RegisterForm onSubmit={handleRegisterSubmit}>
                <label htmlFor="username">Username</label>
                <input 
                type="text" 
                id="username" 
                name="username"
                value={formData.username}
                onChange={handleInputChange} 
                required={true}
                />

                <label htmlFor="password">Password</label>
                <input 
                type="password" 
                id="password" 
                name="password" 
                value={formData.password}
                onChange={handleInputChange}
                required={true}
                />

                <button type="submit">Submit</button>
            </RegisterForm>

            <p>Already have a account? <span onClick={handleLoginClick}>Login here</span> </p>
        </RegisterPage>
    )
}

export default Register