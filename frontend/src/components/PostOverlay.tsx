import React, { useState } from 'react';
import './css/PostOverlay.css'; 

type PostOverlayProps = {
  onClose: () => void; // オーバーレイを閉じるハンドラ
  onSubmit: (content: string) => void; // 投稿内容を送信するハンドラ
};

const PostOverlay: React.FC<PostOverlayProps> = ({ onClose, onSubmit }) => {
  const [content, setContent] = useState('');

  const handleContentChange = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
    setContent(e.target.value);
  };

  const handleSubmit = () => {
    onSubmit(content);
    onClose(); // 投稿後にオーバーレイを閉じる
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
        <div className="overlay-buttons">
          <button onClick={onClose}>キャンセル</button>
          <button onClick={handleSubmit}>投稿する</button>
        </div>
      </div>
    </div>
  );
};

export default PostOverlay;
