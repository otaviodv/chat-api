import { useEffect, useState, Fragment } from 'react';
import { Message, Room } from '../types';
import { createMessage, getRoom } from '../api';
import { useParams } from 'react-router';


export default function JoinRoom() {
  const [msg, setMsg] = useState('');
  const [room, setRoom] = useState<Partial<Room>>({});
  const [messages, setMessages] = useState<Message[]>([]);
  const [username, setUsername] = useState('');
  const { roomId } = useParams();




  useEffect(() => {
    let username = sessionStorage.getItem('username');
    if (!username) {
      username = prompt('type your username');
      sessionStorage.setItem('username', username);
    }
    setUsername(username);
  }, []);

  useEffect(() => {
    const getCurrentRoom = async () => {
      const data = await getRoom(roomId);
      setRoom({
        title: data.title,
        createdAt: data.createdAt,
        id: data.id,
      });
      setMessages(data.messages);
    }
    getCurrentRoom();
  }, [roomId]);

  useEffect(() => {
    const ws = new WebSocket(`${import.meta.env.VITE_APP_WS_URL}/ws/subscribe/room/${roomId}`);

    ws.onopen = () => console.log('ws connected');
    ws.onclose = () => console.log('ws disconnected');
    ws.onmessage = event => {
      const data: Message = JSON.parse(event.data);
      console.log(data);
      setMessages(m => [...m, data]);
    }
    return () => ws.close();
  }, [roomId]);

  async function sendMessage(event: any) {
    event.preventDefault();
    const resp = await createMessage(roomId, {
      username: username,
      text: msg,
    });
    console.log(resp);
    setMsg('');
  }


  return (
    <>
      <h2>{room.title}</h2>
      <div className="chatBox">
        {messages.map((m) =>
          <Fragment key={m.id}>
            <p><b>{m.username}</b>: {m.text}</p>
          </Fragment>
        )}
      </div>
      <form onSubmit={sendMessage}>

        <input value={msg}
          onChange={(ev) => setMsg(ev.target.value)}
          style={{ width: '100%', marginTop: 20 }} />
        <br />
        <br />
        <button type='submit'>
          Send Message
        </button>
      </form>

    </>
  )
}

