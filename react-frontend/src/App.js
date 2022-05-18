import { BrowserRouter as Router, Routes, Switch, Route } from "react-router-dom";
import React from "react";
import Dashboard from "./screens/Dashboard";
import EditArticle from "./components/EditArticle";


function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Dashboard />} />
        <Route path="/editarticle/:id" element={<EditArticle />} />
      </Routes>
    </Router>
  );
}

export default App;
