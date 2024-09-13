import React, { useRef, useEffect } from "react";
import Pin, { PinProps } from "./Pin";
import './Map.css';

type MapProps = {
  src: string;
  pins: PinProps[];
  width?: number;
  height?: number;
};

const Map: React.FC<MapProps> = ({ src: path, pins, width = 500, height = 400 }) => {
  const mapRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    const mapElement = mapRef.current;
    if (!mapElement) return;

    let isDragging = false;
    let startX = 0;
    let startY = 0;
    let scrollLeft = 0;
    let scrollTop = 0;

    const onMouseDown = (e: MouseEvent) => {
      isDragging = true;
      startX = e.pageX - mapElement.offsetLeft;
      startY = e.pageY - mapElement.offsetTop;
      scrollLeft = mapElement.scrollLeft;
      scrollTop = mapElement.scrollTop;
      mapElement.style.cursor = 'grabbing'; // ドラッグ中のポインター
    };

    const onMouseMove = (e: MouseEvent) => {
      if (!isDragging) return;
      e.preventDefault();
      const x = e.pageX - mapElement.offsetLeft;
      const y = e.pageY - mapElement.offsetTop;
      const walkX = (x - startX) * 1; // 移動量
      const walkY = (y - startY) * 1;
      mapElement.scrollLeft = scrollLeft - walkX;
      mapElement.scrollTop = scrollTop - walkY;
    };

    const onMouseUpOrLeave = () => {
      isDragging = false;
      mapElement.style.cursor = 'grab'; // ドラッグ終了時のポインター
    };

    mapElement.addEventListener('mousedown', onMouseDown);
    mapElement.addEventListener('mousemove', onMouseMove);
    mapElement.addEventListener('mouseup', onMouseUpOrLeave);
    mapElement.addEventListener('mouseleave', onMouseUpOrLeave);

    // クリーンアップ
    return () => {
      mapElement.removeEventListener('mousedown', onMouseDown);
      mapElement.removeEventListener('mousemove', onMouseMove);
      mapElement.removeEventListener('mouseup', onMouseUpOrLeave);
      mapElement.removeEventListener('mouseleave', onMouseUpOrLeave);
    };
  }, []);

  return (
    <div
      ref={mapRef}
      className="map-container"
      style={{
        width: `${width}px`,
        height: `${height}px`,
        overflow: 'hidden', // スクロールバーを隠す
        cursor: 'grab', // 初期状態のポインター
        position: 'relative'
      }}
    >
      <div style={{ position: 'relative', width: '100%', height: '100%' }}>
        <img
          src={path}
          alt="会場地図"
          style={{
            display: 'block',
            width: '100%',
            height: 'auto',
            objectFit: 'cover',
            userSelect: 'none',  // 画像の選択を無効化
            pointerEvents: 'none', // ポインターイベントを無効化
          }}
          draggable="false" // 画像のドラッグを無効化
        />
        {pins.map((pin) => (
          <Pin key={pin.id} {...pin} />
        ))}
      </div>
    </div>
  );
};

export default Map;
