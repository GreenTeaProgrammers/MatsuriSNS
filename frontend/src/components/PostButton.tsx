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

  const handlePostSubmit = (content: string) => {
    alert(`投稿しました: ${content}`); // TODO: 実際の投稿処理をここに追加
  };

  return (
    <div>
      <button className="post-button" onClick={handleOpenOverlay}>
        投稿する
      </button>
      {isOverlayVisible && (
        <PostOverlay
          onClose={handleCloseOverlay}
          onSubmit={handlePostSubmit}
        />
      )}
    </div>
  );
};

export default PostButton;
