import React, { useState } from 'react';
import PostOverlay from './PostOverlay';
import './css/PostButton.css'; 
import { createPost } from '../services/postService'; // createPost 関数をインポート
import { useSelector } from 'react-redux';
import { RootState } from '../store';

const PostButton: React.FC = () => {
  const [isOverlayVisible, setIsOverlayVisible] = useState(false);

  // Redux ストアからユーザー情報を取得
  const user = useSelector((state: RootState) => state.auth.user);

  const handleOpenOverlay = () => {
    setIsOverlayVisible(true);
  };

  const handleCloseOverlay = () => {
    setIsOverlayVisible(false);
  };

  // 投稿処理を追加
  const handlePostSubmit = async (content: string, pinPosition?: { x: string; y: string }) => {
    try {
      // ユーザーがログインしているか確認
      if (!user) {
        alert('投稿するにはログインが必要です');
        return;
      }

      // 投稿データを準備
      const postData = {
        user_id: user.id,          // ユーザーID
        event_id: 1,               // 実際のイベントIDに置き換えてください
        content: content,
        location_x: pinPosition ? parseFloat(pinPosition.x) : 0,
        location_y: pinPosition ? parseFloat(pinPosition.y) : 0,
        video_url: '',             // 必要に応じてビデオURLを追加
      };

      // createPost 関数を呼び出してバックエンドに投稿を送信
      const newPost = await createPost(postData);

      alert('投稿が完了しました');

      // 投稿成功後にオーバーレイを閉じる
      handleCloseOverlay();

      // 必要に応じて、投稿一覧の更新やUIの更新を行う
    } catch (error) {
      console.error('投稿に失敗しました:', error);
      alert('投稿に失敗しました。再度お試しください。');
    }
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
