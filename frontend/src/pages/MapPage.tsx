import React, { useState } from 'react';
import Map from '../components/Map';
import webpImage from '../assets/sample_map.webp';
import data from '../mock/mock_post.json'; 

type SelectedPost = {
  comment: string;
  createdAt: string;
} | null;

const MapPage: React.FC = () => {
  const [selectedPost, setSelectedPost] = useState<SelectedPost>(null);

  const pins = data.map((item, index) => ({
    id: `pin${index + 1}`,
    x: item.location_x,
    y: item.location_y,
    comment: item.comment,
    createdAt: item.created_at, // 投稿時刻を追加
  }));

  // ピンがクリックされた時のハンドラ
  const handlePinClick = (comment: string, createdAt: string) => {
    setSelectedPost({ comment, createdAt }); // クリックされたピンのコメントと時刻をセット
  };

  return (
    <div>
      <h1>MapPage</h1>
      <Map
        src={webpImage}
        pins={pins.map(pin => ({
          ...pin,
          onClick: () => handlePinClick(pin.comment, pin.createdAt), 
        }))}
      />
      {selectedPost && (
        <div style={{ marginTop: '20px', padding: '10px', border: '1px solid #ccc' }}>
          <h2>投稿の内容</h2>
          <p>{selectedPost.comment}</p>
          <p><strong>投稿時刻:</strong> {new Date(selectedPost.createdAt).toLocaleString()}</p> 
        </div>
      )}
    </div>
  );
};

export default MapPage;
