import React from "react";
import Pin, { PinProps } from "./Pin"; 

type MapProps = {
  src: string;        // 地図画像のパス
  pins: PinProps[];   // ピンの配列
};

const Map: React.FC<MapProps> = ({ src: path, pins }) => {
  return (
    <div style={{ position: 'relative', display: 'inline-block' }}>
      <img src={path} alt="会場地図" style={{ display: 'block' }} />
      {pins.map((pin) => (
        <Pin key={pin.id} {...pin} />
      ))}
    </div>
  );
};

export default Map;
