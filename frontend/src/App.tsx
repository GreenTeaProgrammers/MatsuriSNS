import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';

import './App.css'
import MapPage from './pages/MapPage';
import HomePage from './pages/HomePage';
import RegisterEventPage from './pages/RegisterEventPage';
import LoginPage from './pages/LoginPage';

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<HomePage />} />
        <Route path="/map" element={<MapPage />} />
        <Route path="/register_event" element={<RegisterEventPage />} />
        <Route path="/login" element={<LoginPage />} />
      </Routes>
    </Router>
  )
}

export default Ap