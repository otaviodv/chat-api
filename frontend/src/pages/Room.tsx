import { useEffect, useState, Fragment } from 'react';
import { Message } from '../types';
import { getRoom } from '../service';
import { useParams } from 'react-router';


export default function JoinRoom() {
  const [msg, setMsg] = useState('');
  const [messages, setMessages] = useState<Message[]>([]);

  const { roomId } = useParams();

  useEffect(() => {
    getRoom(roomId).then((data) => {
      setMessages(data.messages);
    });
  }, [roomId]);

  async function sendMessage(event: any) {
    event.preventDefault();

    setMsg('');
  }


  return (
    <>
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

