import { useState } from 'react';
import { useNavigate } from 'react-router';
import { createRoom } from '../api';

export default function JoinRoom() {
  const [roomName, setRoomName] = useState('');
  const [username, setUsername] = useState('');

  const navigate = useNavigate();

  async function joinRoom(event: any) {
    event.preventDefault();
    sessionStorage.setItem('username', username);
    const room = await createRoom({ title: roomName });
    navigate(`room/${room.id}`);
  }

  return (
    <>
      <div style={{ margin: '0 auto' }}>
        <h1>Chat Api</h1>
        <form onSubmit={joinRoom}>

          <input value={username} placeholder='username' onChange={(ev) => setUsername(ev.target.value)} />
          <br />
          <input value={roomName} placeholder='Room Name' onChange={(ev) => setRoomName(ev.target.value)} />
          <br />
          <br />
          <button type='submit'>
            Join Room!
          </button>
        </form>
        <p className="read-the-docs">
          Join a Room.
        </p>
      </div>
    </>
  )
}

