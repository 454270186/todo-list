import React from "react";
import styled from 'styled-components'


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
    const handleRegisterClick = () => {
        history.push('/register')
    }

    return (
        <LoginPage>
            <h1>Login</h1>
            <LoginForm>
                <label htmlFor="username">Username</label>
                <input type="text" id="username" name="username" />

                <label htmlFor="password">Password</label>
                <input type="password" id="password" name="password" />

                <button type="submit">Confirm</button>
            </LoginForm>

            <p>Do not have a account? <span onClick={handleRegisterClick}>Register here</span></p>
        </LoginPage>
    )
}

export default Login


