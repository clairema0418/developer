import React from 'react';
import PropTypes from 'prop-types';
import './InputField.scss';

const InputField = ({ id, value, onChange, placeholder }) => {
  return (
    <input
      id={id}
      className="input-field"
      type="text"
      value={value}
      onChange={onChange}
      placeholder={placeholder}
    />
  );
};

InputField.propTypes = {
  id: PropTypes.string.isRequired,
  value: PropTypes.string.isRequired,
  onChange: PropTypes.func.isRequired,
  placeholder: PropTypes.string,
};

InputField.defaultProps = {
  placeholder: '',
};

export default InputField;