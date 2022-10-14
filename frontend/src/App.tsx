import React from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Navbar from "./components/Navbar";
//import Users from "./components/Users";
import UserCreate from "./components/UserCreate";
export default function App() {
return (
  <Router>
   <div>
   <Navbar />
   <UserCreate />
   </div>
  </Router>
);
}