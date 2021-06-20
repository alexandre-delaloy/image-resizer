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
          <div className='edit-profile-img'>
            <img src={ uploadIcon } alt='Edit picture' onClick={ () => this.fileInput.click() } />
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
