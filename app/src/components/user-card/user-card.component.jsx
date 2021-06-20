import React from 'react';

import avatar from '../../assets/photo.jpg';
import uploadIcon from '../../assets/Iconly/upload-icon.png';

import './user-card.styles.scss';

const UserCardComponent = () => {
  return (
    <div className='card'>
      <div className='profile-pic'>
        <img className='picture' src={ avatar } alt='User Picture' />
        <div className='edit-profile-img'>
          <img src={ uploadIcon }  alt='Edit picture'/>
        </div>
      </div>
      <div className='infos'>
        <div className='username'>Username</div>
        <div className='occupation'>Occupation</div>
      </div>
    </div>
  );
};

export default UserCardComponent;
