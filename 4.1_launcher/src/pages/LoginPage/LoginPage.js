import React, { useState } from 'react';
import { useDispatch } from 'react-redux';
import { useHistory } from 'react-router-dom';
import { Logo } from '../../components/Logo/Logo';
import { InputField } from '../../components/InputField/InputField';
import { ConfirmButton } from '../../components/ConfirmButton/ConfirmButton';
import { validateKey } from '../../utils/validateKey';
import { login } from '../../store/authSlice';
import './LoginPage.scss';

const LoginPage = () => {
  const [key, setKey] = useState('');
  const [error, setError] = useState(null);
  const dispatch = useDispatch();
  const history = useHistory();

  const handleKeyChange = (e) => {
    setKey(e.target.value);
  };

  const handleConfirmClick = async () => {
    try {
      const isValid = await validateKey(key);
      if (isValid) {
        dispatch(login(key));
        history.push('/game-lobby');
      } else {
        setError('Invalid key. Please try again.');
      }
    } catch (err) {
      setError('Error validating key. Please try again.');
    }
  };

  return (
    <div className="login-page">
      <Logo />
      <InputField
        id="login-input"
        type="text"
        value={key}
        onChange={handleKeyChange}
        placeholder="Enter your key"
      />
      {error && <p className="error-message">{error}</p>}
      <ConfirmButton id="confirm-button" onClick={handleConfirmClick}>
        Confirm
      </ConfirmButton>
    </div>
  );
};

export default LoginPage;