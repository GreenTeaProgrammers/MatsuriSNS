import React, { useState } from 'react';
import PostOverlay from './PostOverlay';
import './css/PostButton.css'; 

const PostButton: React.FC = () => {
  const [isOverlayVisible, setIsOverlayVisible] = useState(false);

  const handleOpenOverlay = () => {
    setIsOverlayVisible(true);
  };

  const handleCloseOverlay = () => {
    setIsOverlayVisible(false);
  };

  // 座標も受け取るように修正
  const handlePostSubmit = (content: string, pinPosition?: { x: string; y: string }) => {
    if (pinPosition) {
      alert(`投稿しました: ${content}\nピンの位置: x=${pinPosition.x}, y=${pinPosition.y}`);
    } else {
      alert(`投稿しました: ${content}`);
    }
    // TODO: 実際の投稿処理をここに追加
  };

  return (
    <div>
      <button className="post-button" onClick={handleOpenOverlay}>
        投稿する
      </button>
      {isOverlayVisible && (
        <PostOverlay
          onClose={handleCloseOverlay}
          onSubmit={handlePostSubmit} // handlePostSubmitを渡す
        />
      )}
    </div>
  );
};

export default PostButton;
