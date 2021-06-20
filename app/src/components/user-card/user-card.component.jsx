import './user-card.styles.scss';

const UserCardComponent = () => {
  return (
    <div className='card'>
      <div className='profile-pic'>
        <img src='../../assets/photo.jpg' alt='User Picture' />
      </div>
    </div>
  );
};

export default UserCardComponent;
