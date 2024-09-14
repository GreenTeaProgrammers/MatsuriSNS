import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';

import './App.css'
import MapPage from './pages/MapPage';
import HomePage from './pages/HomePage';

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<HomePage />} />
        <Route path="/map" element={<MapPage />} />
      </Routes>
    </Router>
  )
}

export default App
