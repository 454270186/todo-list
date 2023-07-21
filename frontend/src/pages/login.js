import React, { useState } from "react";
import {Navigate} from "react-router-dom";
import styled from 'styled-components'
import { loginReq } from "../api/login";

const LoginPage = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100vh;
`;

const LoginForm = styled.form`
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 1rem;
  padding: 1rem;
  border: 1px solid #ccc;
  border-radius: 5px;
  max-width: 300px;
`;

const Login = ({history}) => {
    const [formData, setFormData] = useState({username: '', password: ''})
    const [error, setError] = useState('')
    const [isLoggedIn, setIsLoggedIn] = useState(false)

    const handleRegisterClick = () => {
        history.push('/register')
    }

    const handleInputChange = (event) => {
        const { name, value } = event.target;
        setFormData({ ...formData, [name]: value });
    };

    const handleLoginSubmit = async (e) => {
        e.preventDefault()
        try {
            const data = await loginReq(formData)
            console.log(data)
            setError('')
            setIsLoggedIn(true)
        } catch(error) {
            setError('Invalid username or password');
            console.error('Login error:', error);
        }
    }

    if (isLoggedIn) {
        return <Navigate to="/todo" />;
    }

    return (
        <LoginPage>
            <h1>Login</h1>
            {error && <p style={{ color: 'red' }}>{error}</p>}
            <LoginForm onSubmit={handleLoginSubmit}>
                <label htmlFor="username">Username</label>
                <input 
                type="text" 
                id="username" 
                name="username"
                value={formData.username}
                onChange={handleInputChange} 
                />

                <label htmlFor="password">Password</label>
                <input 
                type="password" 
                id="password" 
                name="password" 
                value={formData.password}
                onChange={handleInputChange}
                />

                <button type="submit">Confirm</button>
            </LoginForm>

            <p>Do not have a account? <span onClick={handleRegisterClick}>Register here</span></p>
        </LoginPage>
    )
}

export default Login


