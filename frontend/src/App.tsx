import React from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Navbar from "./components/Navbar";
//import Users from "./components/Users";
import PatientCreate from "./components/PatientCreate";
export default function App() {
return (
  <Router>
   <div>
   <Navbar />
   <PatientCreate />
   </div>
  </Router>
);
}