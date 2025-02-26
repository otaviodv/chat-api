import { Room } from "./types";

export const getRoom = async (roomId: string) => {
  const resp = await fetch(`http://localhost:8080/rooms/${roomId}`);
  return resp.json() as Promise<Room>;
}


export const createRoom = async (data: Partial<Room>) => {
  const resp = await fetch(`http://localhost:8080/rooms/${data}`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
  });
  return resp.json() as Promise<{ id: string }>;
}