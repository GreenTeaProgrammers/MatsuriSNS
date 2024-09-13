import React from 'react';
import Map from '../components/Map';
import webpImage from '../assets/sample_map.webp';

const MapPage: React.FC = () => {
  return (
    <div>
      <h1>MapPage</h1>
      <Map src={webpImage}></Map>
      <p>これは新しいページです。</p>
    </div>
  );
};

export default MapPage;
