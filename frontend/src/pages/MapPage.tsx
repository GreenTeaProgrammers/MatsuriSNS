import React, { useState } from 'react';
import Map from '../components/Map';
import TimeSlider from '../components/TimeSlider'; 
import SelectedPost from '../components/SelectedPost';
import webpImage from '../assets/sample_map.webp';
import data from '../mock/mock_post.json';

type SelectedPostType = {
  comment: string;
  createdAt: string;
} | null;

const MapPage: React.FC = () => {
  const [selectedPost, setSelectedPost] = useState<SelectedPostType>(null);
  const [sliderValue, setSliderValue] = useState(0);

  // データの最小・最大時刻を取得
  const minTime = new Date(data[0].created_at).getTime();
  const maxTime = new Date(data[data.length - 1].created_at).getTime();

  // JSONデータからピンデータを生成
  const pins = data.map((item, index) => ({
    id: `pin${index + 1}`,
    x: item.location_x,
    y: item.location_y,
    comment: item.comment,
    createdAt: item.created_at,
    isVisible: new Date(item.created_at).getTime() <= sliderValue, // スライダーの時刻以下のピンを表示
  }));

  // スライダーが変更された時のハンドラ
  const handleSliderChange = (value: number) => {
    setSliderValue(value);
  };

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
      <TimeSlider
        minTime={minTime}
        maxTime={maxTime}
        value={sliderValue}
        onChange={handleSliderChange}
      />
      {selectedPost && (
        <SelectedPost comment={selectedPost.comment} createdAt={selectedPost.createdAt} />
      )}
    </div>
  );
};

export default MapPage;
