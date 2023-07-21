import React from "react";
import styled from 'styled-components'

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

const Register = ({history}) => {
    const handleLoginClick = () => {
        history.push('/login')
    }

    return (
        <RegisterPage>
            <RegisterForm>
                <label htmlFor="username">Username</label>
                <input type="text" id="username" name="username"/>

                <label htmlFor="password">Password</label>
                <input type="password" id="password" name="password" />

                <button type="submit">Submit</button>
            </RegisterForm>

            <p>Already have a account? <span onClick={handleLoginClick}>Login here</span> </p>
        </RegisterPage>
    )
}

export default Register