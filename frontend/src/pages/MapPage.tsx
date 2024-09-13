import React from 'react';
import Map from '../components/Map';
import webpImage from '../assets/sample_map.webp';

const MapPage: React.FC = () => {
  const pins = [
    { id: 'pin1', x: 0.2, y: 0.3 },
    { id: 'pin2', x: 0.5, y: 0.5 },
    { id: 'pin3', x: 0.8, y: 0.7 },
  ];

  return (
    <div>
      <h1>MapPage</h1>
      <Map src={webpImage} pins={pins}></Map>
      <p>これは新しいページです。</p>
    </div>
  );
};

export default MapPage;
