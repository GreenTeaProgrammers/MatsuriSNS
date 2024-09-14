import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';

import './App.css'
import MapPage from './pages/MapPage';
import HomePage from './pages/HomePage';
import RegisterEventPage from './pages/RegisterEventPage';

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<HomePage />} />
        <Route path="/map" element={<MapPage />} />
        <Route path="/register_event" element={<RegisterEventPage />} />
      </Routes>
    </Router>
  )
}

export default App
