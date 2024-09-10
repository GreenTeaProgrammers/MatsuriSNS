import { User } from './user';
import { Event } from './event';
import { PostImage } from './postImage';

export interface Post {
    user: User;
    event: Event;
    comment: string;
    videoUrl?: string;
    locationX: number;
    locationY: number;
    images: PostImage[]; 
    createdAt: Date;
    updatedAt: Date;
}

