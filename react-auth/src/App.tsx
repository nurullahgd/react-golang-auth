import React, { useEffect, useState } from "react";
import "./App.css";
import Login from "./pages/login";
import Register from "./pages/register";
import Home from "./pages/home";

import Nav from "./components/Nav";
import { BrowserRouter, Routes, Route } from "react-router-dom";

function App() {
  const [name, setName] = useState("");
  useEffect(() => {
    (
      async () => {
        const response = await fetch("http://localhost:8000/api/user", {
          method: "GET",
          credentials: "include",
          headers: {
            "Content-Type": "application/json",
          },
        });
        if (response.status === 200) {
          const data = await response.json();
          setName(data.name);
        }
      }
    )();
  });
  return (
    <div className="App">
      <BrowserRouter>
        <Nav name={name} setName={setName} />

        <main className="form-signin w-100 m-auto">
          <Routes>
            <Route path="/" element={<Home name={name} />} />
            <Route path="/login" element={<Login setName={setName} />} />
            <Route path="/register" element={<Register />} />
          </Routes>
        </main>
      </BrowserRouter>
    </div>
  );
}

export default App;