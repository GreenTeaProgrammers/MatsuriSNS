import { User } from './user';
import { Event } from './event';


export interface EventAdmin {
    event: Event;
    user: User;
    createdAt: Date;
}

