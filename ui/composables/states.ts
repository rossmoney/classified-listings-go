import type { Listing } from "~/types/Listing";

export const useCurrentListingState = () => useState<Listing>('currentListing', () => (
    {
        id: 0,
        title: '',
        description: '',
        date_posted: new Date(),
        date_posted_human: '',
        price: '0.00',
        category: '',
        active: false,
    }
));

export const useAllListingsState = () => useState<Listing[]>('listings', () => []);
