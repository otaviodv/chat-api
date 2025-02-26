import { BrowserRouter, Route, Routes } from "react-router";
import JoinRoom from "./pages/JoinRoom";
import Room from "./pages/Room";

export default function Router() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<JoinRoom />} />
        <Route path="/room/:roomId" element={<Room />} />
      </Routes>
    </BrowserRouter>
  );
}