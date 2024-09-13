import React, { useState } from 'react';
import Map from '../components/Map';
import webpImage from '../assets/sample_map.webp';
import data from '../mock/mock_post.json'; 

const MapPage: React.FC = () => {
  const [selectedComment, setSelectedComment] = useState<string | null>(null);

  // JSONデータからピンデータを生成
  const pins = data.map((item, index) => ({
    id: `pin${index + 1}`,
    x: item.location_x,
    y: item.location_y,
    comment: item.comment,
  }));

  const handlePinClick = (comment: string) => {
    setSelectedComment(comment); 
  };

  return (
    <div>
      <h1>MapPage</h1>
      <Map
        src={webpImage}
        pins={pins.map(pin => ({
          ...pin,
          onClick: handlePinClick, 
        }))}
      />
      {selectedComment && (
        <div style={{ marginTop: '20px', padding: '10px', border: '1px solid #ccc' }}>
          <h2>投稿の内容</h2>
          <p>{selectedComment}</p>
        </div>
      )}
      <p>これは新しいページです。</p>
    </div>
  );
};

export default MapPage;
