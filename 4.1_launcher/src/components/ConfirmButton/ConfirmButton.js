import React from 'react';
import PropTypes from 'prop-types';
import './ConfirmButton.scss';

const ConfirmButton = ({ onClick }) => {
  return (
    <button className="confirm-button" onClick={onClick}>
      Confirm
    </button>
  );
};

ConfirmButton.propTypes = {
  onClick: PropTypes.func.isRequired,
};

export default ConfirmButton;