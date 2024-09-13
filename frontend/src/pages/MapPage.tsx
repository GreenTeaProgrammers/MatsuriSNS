import React from 'react';
import Map from '../components/Map';
import webpImage from '../assets/sample_map.webp';
import mockData from '../mock/mock_post.json';

const MapPage: React.FC = () => {

   const pins = mockData.map((item, index) => ({
    id: `pin${index + 1}`, //TODO: スキーマに連番が含まれる      
    x: item.location_x,         
    y: item.location_y          
  }));

  return (
    <div>
      <h1>MapPage</h1>
      <Map src={webpImage} pins={pins}></Map>
      <p>これは新しいページです。</p>
    </div>
  );
};

export default MapPage;
