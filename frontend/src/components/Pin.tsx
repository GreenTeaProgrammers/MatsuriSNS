import React, { useState, useEffect } from "react";
import pinIcon from "../assets/pin.svg";
import './Pin.css'; 

export type PinProps = {
  x: number;
  y: number;
  id: string;
  comment: string;
  onClick: (comment: string) => void; // クリック時のハンドラ
  isVisible: boolean; // ピンが表示されるかどうか
};

const Pin: React.FC<PinProps> = ({ x, y, comment, onClick, isVisible }) => {
  const [shouldRender, setShouldRender] = useState(isVisible);

  useEffect(() => {
    if (isVisible) {
      setShouldRender(true);
    }
  }, [isVisible]);

  const handleAnimationEnd = () => {
    if (!isVisible) {
      setShouldRender(false);
    }
  };

  const clampedX = Math.max(0, Math.min(1, x));
  const clampedY = Math.max(0, Math.min(1, y));

  if (!shouldRender) {
    return null; // ピンをレンダリングしない
  }

  return (
    <div
      className={`pin ${isVisible ? '' : 'disappear'}`} // 消える際にアニメーションを適用
      style={{
        position: 'absolute',
        left: `${clampedX * 100}%`,
        top: `${(1 - clampedY) * 100}%`,
        transform: 'translate(-50%, -50%)',
        cursor: 'pointer',
        pointerEvents: 'auto',
        userSelect: 'none',
      }}
      onClick={() => onClick(comment)}
      onAnimationEnd={handleAnimationEnd} // アニメーション終了時に処理を実行
    >
      <img
        src={pinIcon}
        alt="ピン"
        style={{
          width: '24px',
          height: '24px',
          userSelect: 'none',
          pointerEvents: 'none',
        }}
        draggable="false"
      />
    </div>
  );
};

export default Pin;
