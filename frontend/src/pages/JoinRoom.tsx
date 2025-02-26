import { useState } from 'react';
import { useNavigate } from 'react-router';

export default function JoinRoom() {
  const [roomName, setRoomName] = useState('');

  const navigate = useNavigate();

  async function joinRoom(event: any) {
    event.preventDefault();
    navigate(`room/${roomName}`);
  }

  return (
    <>
      <div style={{ margin: '0 auto' }}>

        <h1>Chat Api</h1>
        <form onSubmit={joinRoom}>

          <input value={roomName} onChange={(ev) => setRoomName(ev.target.value)} />
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

