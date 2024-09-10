import { User } from './user';

export interface Event {
    id: number;
    title: string;
    description?: string;
    mapUrl: string;
    qrCodeUrl?: string;
    createdBy: User;
    createdAt: Date;
    updatedAt: Date;
}

