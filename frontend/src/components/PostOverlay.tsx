import React, { useState } from 'react';
import './css/PostOverlay.css'; 

type PostOverlayProps = {
  onClose: () => void; // オーバーレイを閉じるハンドラ
  onSubmit: (content: string) => void; // 投稿内容を送信するハンドラ
};

const MAX_CHAR_LIMIT = 140; // 最大文字数

const PostOverlay: React.FC<PostOverlayProps> = ({ onClose, onSubmit }) => {
  const [content, setContent] = useState('');

  const handleContentChange = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
    // 文字数制限のバリデーション
    if (e.target.value.length <= MAX_CHAR_LIMIT) {
      setContent(e.target.value);
    }
  };

  const handleSubmit = () => {
    if (content.length > 0) {
      onSubmit(content);
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
        <div className="overlay-buttons">
          <button onClick={onClose}>キャンセル</button>
          <button onClick={handleSubmit}>投稿する</button>
        </div>
      </div>
    </div>
  );
};

export default PostOverlay;
