import { useState, useEffect } from "react";
import pinIcon from "../assets/pin.svg"; // SVG のパスを指定
import './css/Pin.css'; 
import { generateRandomColor } from "../util/ColorUtils";

export type PinProps = {
  x: number;
  y: number;
  id: string;
  comment: string;
  onClick: (comment: string) => void; // クリック時のハンドラ
  isVisible: boolean; // ピンが表示されるかどうか
  color?: string; // ピンの色を指定するプロパティ（オプショナル）
   maxHorizontalPercentage?: number; // 最大の水平方向の割合
  maxVerticalPercentage?: number; // 最大の垂直方向の割合
};


const Pin: React.FC<PinProps> = ({ x, y, comment, onClick, isVisible, color, maxHorizontalPercentage=100, maxVerticalPercentage=100}) => {
  const [shouldRender, setShouldRender] = useState(isVisible);
  const [pinColor] = useState(color || generateRandomColor()); 

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

  // x, y の値を 0〜1 の範囲でクランプする
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
        left: `${clampedX * maxHorizontalPercentage}%`,
        top: `${(1 - clampedY) * maxVerticalPercentage}%`, // Y 座標は 1 - y で計算
        transform: 'translate(-50%, -50%)', // ピンの中心を基準に配置
        cursor: 'pointer',
        pointerEvents: 'auto',
        userSelect: 'none',
        width: '24px',
        height: '24px',
        backgroundColor: pinColor, // 背景色を設定
        maskImage: `url(${pinIcon})`, // SVG をマスクとして使用
        maskSize: 'contain', // マスクのサイズを調整
        maskPosition: 'center', // マスクの位置を中央に設定
        maskRepeat: 'no-repeat', // マスクの繰り返しを防止
        WebkitMaskImage: `url(${pinIcon})`, // Webkit 対応のため
        WebkitMaskSize: 'contain', // Webkit 対応のため
        WebkitMaskPosition: 'center', // Webkit 対応のため
        WebkitMaskRepeat: 'no-repeat', // Webkit 対応のため
      }}
      onClick={() => onClick(comment)}
      onAnimationEnd={handleAnimationEnd} // アニメーション終了時に処理を実行
    />
  );
};

export default Pin;
