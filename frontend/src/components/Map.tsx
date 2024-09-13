import React from "react";
import Pin, { PinProps } from "./Pin";
import './Map.css';

type MapProps = {
  src: string;
  pins: PinProps[];
  width?: number;  
  height?: number;
};

const Map: React.FC<MapProps> = ({ src: path, pins, width, height }) => {
  return (
    <div
		className="map-container"
      style={{
        width: width || '500px',     // 幅を固定
        height: height || '400px',    // 高さを固定
        overflow: 'auto',   // スクロール可能にする
      }}
    >
      <div style={{ position: 'relative', width: '100%' }}>
        <img src={path} alt="会場地図" style={{ display: 'block', width: '100%', height: 'auto' }} />
        {pins.map((pin) => (
          <Pin key={pin.id} {...pin} />
        ))}
      </div>
    </div>
  );
};

export default Map;
