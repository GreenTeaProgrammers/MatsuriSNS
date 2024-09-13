import React from "react";
import pinIcon from "../assets/pin.svg"


export type PinProps = {
  x: number;   // ピンの X 座標（0 から 1 の値）
  y: number;   // ピンの Y 座標（0 から 1 の値）
  id: string;  
};

const Pin: React.FC<PinProps> = ({ x, y, id }) => {
  const clampedX = Math.max(0, Math.min(1, x));
  const clampedY = Math.max(0, Math.min(1, y));

  return (
    <div
      key={id}
      style={{
        position: 'absolute',
        left: `${clampedX * 100}%`,
        bottom: `${clampedY * 100}%`,
        transform: 'translate(-50%, 50%)',
        pointerEvents: 'none',
      }}
    >
		<img src={pinIcon} alt="ピン" 
			style={{ width: '30px', height: '30px' }} 
		/>
    </div>
  );
};

export default Pin;
