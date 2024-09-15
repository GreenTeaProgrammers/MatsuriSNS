import api from "./api";

// Define the types for the Post data
export interface CreatePostRequest {
  user_id: number;
  event_id: number;
  content: string;
  location_x: number;
  location_y: number;
  video_url?: string;
}

export interface Post {
  id: number;
  user_id: number;
  event_id: number;
  content: string;
  location_x: number;
  location_y: number;
  video_url?: string;
  created_at: string;
  updated_at: string;
}

// Service for creating a new post
export const createPost = async (postData: CreatePostRequest): Promise<Post> => {
  try {
    const response = await api.post<Post>("/posts", postData);
    return response.data;
  } catch (error) {
    throw new Error(`Failed to create post: ${(error as any).message}`);
  }
};

// Service for fetching a post by ID
export const getPost = async (postId: number): Promise<Post> => {
  try {
    const response = await api.get<Post>(`/posts/${postId}`);
    return response.data;
  } catch (error) {
    throw new Error(`Failed to fetch post: ${(error as any).message}`);
  }
};

// Service for fetching all posts
export const getAllPosts = async (): Promise<Post[]> => {
  try {
    const response = await api.get<Post[]>("/posts");
    return response.data;
  } catch (error) {
    throw new Error(`Failed to fetch posts: ${(error as any).message}`);
  }
};

// Service for updating an existing post by ID
export const updatePost = async (postId: number, postData: Partial<CreatePostRequest>): Promise<Post> => {
  try {
    const response = await api.put<Post>(`/posts/${postId}`, postData);
    return response.data;
  } catch (error) {
    throw new Error(`Failed to update post: ${(error as any).message}`);
  }
};

// Service for deleting a post by ID
export const deletePost = async (postId: number): Promise<void> => {
  try {
    await api.delete(`/posts/${postId}`);
  } catch (error) {
    throw new Error(`Failed to delete post: ${(error as any).message}`);
  }
};
