import React, { useState, useEffect } from 'react';
import Map from '../components/Map';
import TimeSlider from '../components/TimeSlider'; 
import SelectedPost from '../components/SelectedPost';
import webpImage from '../assets/sample_map.webp';
import PostButton from '../components/PostButton';
import { getAllPosts, Post } from '../services/postService'; // Import the Post Service

type SelectedPostType = {
  comment: string;
  createdAt: string;
} | null;

const MapPage: React.FC = () => {
  const [selectedPost, setSelectedPost] = useState<SelectedPostType>(null);
  const [sliderValue, setSliderValue] = useState<number>(0);
  const [posts, setPosts] = useState<Post[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);

  // Fetch posts data from the API
  useEffect(() => {
    const fetchPosts = async () => {
      try {
        const data = await getAllPosts();
        setPosts(data);
        setLoading(false);

        // Set initial slider value to the latest post's time
        if (data.length > 0) {
          const times = data.map(post => new Date(post.created_at).getTime());
          const maxTime = Math.max(...times);
          setSliderValue(maxTime);
        }
      } catch (error) {
        setError((error as Error).message);
        setLoading(false);
      }
    };

    fetchPosts();
  }, []);

  // Handle loading and error states
  if (loading) {
    return <div>読み込み中...</div>;
  }

  if (error) {
    return <div>エラーが発生しました: {error}</div>;
  }

  // Calculate the minimum and maximum times from the posts
  const times = posts.map(post => new Date(post.created_at).getTime());
  const minTime = Math.min(...times);
  const maxTime = Math.max(...times);

  // Generate pins from the fetched posts
  const pins = posts.map((post, index) => ({
    id: `pin${index + 1}`,
    x: post.location_x,
    y: post.location_y,
    comment: post.content,
    createdAt: post.created_at,
    isVisible: new Date(post.created_at).getTime() <= sliderValue, // Show pins up to the slider's time
  }));

  // Slider change handler
  const handleSliderChange = (value: number) => {
    setSliderValue(value);
  };

  // Pin click handler
  const handlePinClick = (comment: string, createdAt: string) => {
    setSelectedPost({ comment, createdAt });
  };

  return (
    <div>
      <h1>MapPage</h1>
      <PostButton />
      <Map
        src={webpImage}
        pins={pins.map(pin => ({
          ...pin,
          onClick: () => handlePinClick(pin.comment, pin.createdAt),
        }))}
      />
      <TimeSlider
        minTime={minTime}
        maxTime={maxTime}
        value={sliderValue}
        onChange={handleSliderChange}
      />
      {selectedPost && (
        <SelectedPost comment={selectedPost.comment} createdAt={selectedPost.createdAt} />
      )}
    </div>
  );
};

export default MapPage;