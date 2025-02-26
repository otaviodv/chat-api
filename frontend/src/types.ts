export interface Message {
  id: string;
  roomId: string;
  username: string;
  text: string;
  createdAt: string;
}

export interface Room {
  id: string;
  name: string;

  createdAt: string;
  messages: Message[];
}

