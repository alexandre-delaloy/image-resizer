import React, { Component } from 'react';

import UserCardComponent from './components/user-card/user-card.component';

import './App.css';

class App extends Component {

  render() {
    return (
      <div className='App'>
        <UserCardComponent />
      </div>
    );
  }
}

export default App;
