import React from "react";
import pinIcon from "../assets/pin.svg";

export type PinProps = {
  x: number;
  y: number;
  id: string;
};

const Pin: React.FC<PinProps> = ({ x, y }) => {
  const clampedX = Math.max(0, Math.min(1, x));
  const clampedY = Math.max(0, Math.min(1, y));

  return (
    <div
      style={{
        position: 'absolute',
        left: `${clampedX * 100}%`,
        top: `${(1 - clampedY) * 100}%`, 
        transform: 'translate(-50%, -50%)', // 位置調整
        pointerEvents: 'none',
      }}
    >
      <img src={pinIcon} alt="ピン" style={{ width: '24px', height: '24px' }} />
    </div>
  );
};

export default Pin;
