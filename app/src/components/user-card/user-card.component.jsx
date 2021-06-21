import React, { Component } from 'react';
import axios from 'axios';

import avatar from '../../assets/photo.jpg';
import uploadIcon from '../../assets/upload-icon.svg';

import './user-card.styles.scss';

class UserCardComponent extends Component {
  state = {
    selectedFile: null
  };

  fileSelectedHandler = event => {
    this.setState({
      selectedFile: event.target.files[0]
    });
  };

  fileUploadHandler = () => {
    const fd = new FormData();
    fd.append('image', this.state.selectedFile, this.state.selectedFile.name);
    axios.post('', fd, {
      onUploadProgress: progressEvent => {
        console.log(`Upload Progress: ${ Math.round(progressEvent.loaded / progressEvent.total * 100) } %`);
      }
    })
      .then(res => {
        console.log(res);
      });
  };

  render() {
    return (
      <div className='card'>
        <div className='profile-pic'>
          <img className='picture' src={ avatar } alt='User Picture' />
          <div className='edit-profile-img' onClick={ () => this.fileInput.click() }>
            <svg className='edit-icon' width="40" height="40" viewBox="0 0 40 40" xmlns="http://www.w3.org/2000/svg">
              <path d="M20 0C31.02 0 40 8.96 40 20C40 31.02 31.02 40 20 40C8.96 40 0 31.02 0 20C0 8.96 8.96 0 20 0ZM14.12 16.06C13.54 15.46 12.58 15.46 12 16.04C11.4 16.64 11.4 17.58 12 18.16L18.94 25.14C19.22 25.42 19.6 25.58 20 25.58C20.4 25.58 20.78 25.42 21.06 25.14L28 18.16C28.3 17.88 28.44 17.5 28.44 17.12C28.44 16.72 28.3 16.34 28 16.04C27.4 15.46 26.46 15.46 25.88 16.06L20 21.96L14.12 16.06Z"/>
            </svg>
          </div>
        </div>
        <div className='infos'>
          <div className='username'>Username</div>
          <div className='created-at'>Creation Date</div>
          <input type='file' name='myImage' accept='image/*'
                 style={{display: 'none'}}
                 onChange={ this.fileSelectedHandler }
                 ref={ fileInput => this.fileInput = fileInput } />
        </div>
      </div>
    );
  }
}


export default UserCardComponent;
