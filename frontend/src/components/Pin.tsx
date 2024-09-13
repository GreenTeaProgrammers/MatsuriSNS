import React from "react";
import pinIcon from "../assets/pin.svg";
import './Pin.css'; // CSSファイルをインポート

export type PinProps = {
  x: number;
  y: number;
  id: string;
  comment: string;
  onClick: (comment: string) => void; // クリック時のハンドラ
};

const Pin: React.FC<PinProps> = ({ x, y, comment, onClick }) => {
  const clampedX = Math.max(0, Math.min(1, x));
  const clampedY = Math.max(0, Math.min(1, y));

  return (
    <div
      className="pin" // アニメーション用のクラスを追加
      style={{
        position: 'absolute',
        left: `${clampedX * 100}%`,
        top: `${(1 - clampedY) * 100}%`,
        transform: 'translate(-50%, -50%)',
        cursor: 'pointer', // マウスカーソルをポインターに
        pointerEvents: 'auto', // クリックイベントを受け取るように
        userSelect: 'none', // ピンの選択を無効化
      }}
      onClick={() => onClick(comment)} // クリック時にコメントを渡す
    >
      <img
        src={pinIcon}
        alt="ピン"
        style={{
          width: '24px',
          height: '24px',
          userSelect: 'none',  // 画像の選択を無効化
          pointerEvents: 'none', // 画像自体はクリックイベントを受け取らない
        }}
        draggable="false" // 画像のドラッグを無効化
      />
    </div>
  );
};

export default Pin;
