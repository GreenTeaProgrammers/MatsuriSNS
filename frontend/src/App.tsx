import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';

import './App.css'
import MapPage from './pages/MapPage';
import HomePage from './pages/HomePage';
import RegisterEventPage from './pages/RegisterEventPage';
import LoginPage from './pages/LoginPage';
import RegisterPage from './pages/RegisterPage';

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<HomePage />} />
        <Route path="/map" element={<MapPage />} />
        <Route path="/register_event" element={<RegisterEventPage />} />
        <Route path="/login" element={<LoginPage />} />
        <Route path="/register" element={<RegisterPage />} />
      </Routes>
    </Router>
  )
}

export default App