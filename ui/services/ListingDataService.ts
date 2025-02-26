import type { Listing } from "~/types/Listing";
import type { AsyncData } from 'nuxt/app';
import type { FetchError } from 'ofetch';

class ListingDataService {

    getAll(baseUrl?: string) {
        let url = (baseUrl ?? '') + '/api/listings';
        return useAsyncData('listings', () => $fetch(url)) as 
            AsyncData<{ message: string, status: number, data: Listing[] }, FetchError>;
    }

    getAllByTitle(title: string) {
        return useFetch(`/api/listings?title=${title}`) as
            AsyncData<{ message: string, status: number, data: Listing[] }, FetchError>;
    }

    getAllByCategory(category: string) {
        return useFetch(`/api/listings?category=${category}`) as 
            AsyncData<{ message: string, status: number, data: Listing[] }, FetchError>;
    }

    get(id: number) {
        let url = `/api/listings/${id}`
        return useAsyncData('currentListing', () => $fetch(url)) as 
            AsyncData<{ message: string, status: number, data: Listing }, FetchError>;
    }

    create(data: Listing) {
        return useFetch( `/api/listings`, { 
            method : "post", 
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)
        }) as AsyncData<{ message: string, status: number, data: Listing }, FetchError>;
    }

    update(id: number, data: Listing) {
        return useFetch( `/api/listings/${id}`, { method : "put", body: JSON.stringify(data)}) as 
            AsyncData<{ message: string, status: number, data: Listing }, FetchError>;
    }

    delete(id: number) {
        return useFetch( `/api/listings/${id}`, { method : "delete"}) as 
            AsyncData<{ message: string, status: number, data: Listing }, FetchError>;
    }
}

export default new ListingDataService();