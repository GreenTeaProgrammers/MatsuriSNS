import React from 'react';

type SelectedPostProps = {
  comment: string;
  createdAt: string;
};

const SelectedPost: React.FC<SelectedPostProps> = ({ comment, createdAt }) => {
  return (
    <div style={{ marginTop: '20px', padding: '10px', border: '1px solid #ccc' }}>
      <h2>投稿の内容</h2>
      <p>{comment}</p>
      <p>
        <strong>投稿時刻:</strong> {new Date(createdAt).toLocaleString()}
      </p>
    </div>
  );
};

export default SelectedPost;
