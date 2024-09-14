import React, { useState } from 'react';
import Map from './Map'; 
import './css/PostOverlay.css'; 
import webpImage from '../assets/sample_map.webp'; // TODO: 画像はフェッチしてくる
import { PinProps } from './Pin';

type PostOverlayProps = {
  onClose: () => void; // オーバーレイを閉じるハンドラ
  onSubmit: (content: string, pinPosition?: { x: string; y: string }) => void; // 投稿内容とピンの座標を送信するハンドラ
};

const MAX_CHAR_LIMIT = 140; // 最大文字数

const PostOverlay: React.FC<PostOverlayProps> = ({ onClose, onSubmit }) => {
  const [content, setContent] = useState('');
  const [pins, setPins] = useState<PinProps[]>([]); // ピンの状態を管理するためのステート

  const handleContentChange = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
    if (e.target.value.length <= MAX_CHAR_LIMIT) {
      setContent(e.target.value);
    }
  };

  const handleMapClick = (e: React.MouseEvent<HTMLDivElement, MouseEvent>) => {
    const mapElement = e.currentTarget.getBoundingClientRect();
    const scrollLeft = e.currentTarget.scrollLeft;
    const scrollTop = e.currentTarget.scrollTop;

    // スクロールや隠れた部分を含めた全体の幅と高さ
    const totalWidth = e.currentTarget.scrollWidth;
    const totalHeight = e.currentTarget.scrollHeight;
    const clientWidth = e.currentTarget.clientWidth;
    const clientHeight = e.currentTarget.clientHeight;
    const scrollWidth = e.currentTarget.scrollWidth;
    const scrollHeight = e.currentTarget.scrollHeight;
    const hiddenWidth = scrollWidth - clientWidth; // 隠れている幅
    const hiddenHeight = scrollHeight - clientHeight; // 隠れている高さ

    // クリック位置を左下(0, 0)、右上(1, 1)の座標系に変換
    const x = (e.clientX - mapElement.left + scrollLeft) / totalWidth;
    const y = 1 - (e.clientY - mapElement.top + scrollTop) / totalHeight; 

    console.log(`x: ${x}, y: ${y}`);

    const newPin: PinProps = {
      id: `pin${pins.length + 1}`,
      x,
      y,
      comment: '', // コメントを追加
      onClick: () => {},
      isVisible: true, 
      maxHorizontalPercentage: (clientWidth + hiddenWidth) / clientWidth * 95,
      maxVerticalPercentage: (clientHeight + hiddenHeight) / clientHeight * 95,
    };

    // 既存のピンをクリアして新しいピンだけをセットする
    setPins([newPin]);
  };

  const handleSubmit = () => {
    if (content.length > 0) {
      if (pins.length > 0) {
        const pin = pins[0];
        const roundedX = pin.x.toFixed(2); // x座標を小数点以下2桁に丸める
        const roundedY = pin.y.toFixed(2); // y座標を小数点以下2桁に丸める
        onSubmit(content, { x: roundedX, y: roundedY }); // 座標と共に投稿内容を送信
      } else {
        onSubmit(content); // ピンがない場合は投稿内容のみを送信
      }
      onClose(); // 投稿後にオーバーレイを閉じる
    } else {
      alert('投稿内容を入力してください。');
    }
  };

  return (
    <div className="overlay">
      <div className="overlay-content">
        <h2>新しい投稿</h2>
        <textarea
          value={content}
          onChange={handleContentChange}
          placeholder="投稿内容を入力してください..."
        />
        <div className="char-counter">
          {content.length} / {MAX_CHAR_LIMIT} 文字
        </div>
        <Map
          src={webpImage}
          pins={pins}
          width={400}
          height={300}
          onClick={handleMapClick} // マップのクリックイベントハンドラ
        />
        <div className="overlay-buttons">
          <button onClick={onClose}>キャンセル</button>
          <button onClick={handleSubmit}>投稿する</button>
        </div>
      </div>
    </div>
  );
};

export default PostOverlay;
