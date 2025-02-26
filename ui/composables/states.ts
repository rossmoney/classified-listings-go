import type { Listing } from "~/types/Listing";
import type { ConfirmState, MessageState } from '~/types/Dialog';

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

export const useConfirmState = () => useState<ConfirmState>('confirm', () => (
    {
        open: false,
        result: false,
        dialog: {
            title: '',
            message: '',
            errors: null,
            options: {
                color: 'primary',
                width: 290,
                zIndex: 200
            },
        },
    }
))