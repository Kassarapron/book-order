import { BrowserRouter, Routes, Route } from "react-router-dom";


import Navbar from "./components/Navbar";

import Users from "./components/Users";

import UserCreate from "./components/UserCreate";


export default function App() {

  return (
    <BrowserRouter>
      <Navbar />
      <Routes>
        <Route path="/" element={<Users />} />
        <Route path="/create" element={<UserCreate />} />
      </Routes>
    </BrowserRouter>
  );

}