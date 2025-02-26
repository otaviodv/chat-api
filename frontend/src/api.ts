import { Message, Room } from "./types";


const baseUrl = import.meta.env.VITE_APP_API_URL;
export const getRoom = async (roomId: string) => {
  const resp = await fetch(`${baseUrl}/rooms/${roomId}`);
  return resp.json() as Promise<Room>;
}


export const createRoom = async (data: Partial<Room>) => {
  const resp = await fetch(`${baseUrl}/rooms`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(data),
  });
  return resp.json() as Promise<{ id: string }>;
}

export async function createMessage(roomId: string, data: Partial<Message>) {
  const resp = await fetch(`${baseUrl}/rooms/${roomId}/messages`,
    {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(data)
    });

  return resp.json();
}

