import{ useRef, useEffect, useState } from "react";
import Pin, { PinProps } from "./Pin";
import './css/Map.css';

type MapProps = {
  src: string;
  pins: PinProps[];
  width?: number;
  height?: number;
  onClick?: (e: React.MouseEvent<HTMLDivElement, MouseEvent>) => void; // クリックハンドラを追加
};

const Map: React.FC<MapProps> = ({ src: path, pins, width = 500, height = 400, onClick }) => {
  const mapRef = useRef<HTMLDivElement>(null);
  const [hiddenPixels, setHiddenPixels] = useState({ hiddenWidth: 0, hiddenHeight: 0 });

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

    const calculateHiddenPixels = () => {
      if (mapElement) {
        const hiddenWidth = mapElement.scrollWidth - mapElement.clientWidth; // 隠れている幅
        const hiddenHeight = mapElement.scrollHeight - mapElement.clientHeight; // 隠れている高さ

        setHiddenPixels({ hiddenWidth, hiddenHeight });
      }
    };

    calculateHiddenPixels();
    window.addEventListener('resize', calculateHiddenPixels);

    const onMouseUpOrLeave = () => {
      isDragging = false;
      mapElement.style.cursor = 'grab'; // ドラッグ終了時のポインター
    };

    mapElement.addEventListener('mousedown', onMouseDown);
    mapElement.addEventListener('mousemove', onMouseMove);
    mapElement.addEventListener('mouseup', onMouseUpOrLeave);
    mapElement.addEventListener('mouseleave', onMouseUpOrLeave);

    return () => {
      mapElement.removeEventListener('mousedown', onMouseDown);
      mapElement.removeEventListener('mousemove', onMouseMove);
      mapElement.removeEventListener('mouseup', onMouseUpOrLeave);
      mapElement.removeEventListener('mouseleave', onMouseUpOrLeave);
      window.removeEventListener('resize', calculateHiddenPixels);
    };
  }, [width, height]);

  return (
    <div
      ref={mapRef}
      className="map-container"
      style={{
        width: `${width}px`,
        height: `${height}px`,
        overflow: 'auto',
        cursor: 'grab',
        position: 'relative',
      }}
      onClick={onClick} // マップクリック時にイベントを発火
    >
      <div style={{ position: 'relative', width: '100%', height: '100%' }}>
        <img
          className="map-image"
          src={path}
          alt="会場地図"
          style={{
            display: 'block',
            width: '100%',
            height: 'auto',
            objectFit: 'cover',
            userSelect: 'none',
            pointerEvents: 'none',
          }}
          draggable="false" // 画像のドラッグを無効化
        />
        {pins.map((pin) => (
          <Pin key={pin.id} {...pin} />
        ))}
      </div>
      <div className="no-select">
        <p>隠れている幅: {hiddenPixels.hiddenWidth}px</p>
        <p>隠れている高さ: {hiddenPixels.hiddenHeight}px</p>
      </div>
    </div>
  );
};

export default Map;
