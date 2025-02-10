class ListingDataService {

    getAll(baseUrl) {
        return useAsyncData('listings', () => $fetch((baseUrl ?? '') + '/api/listings'), {server: false});
    }

    getAllByTitle(title) {
        return useFetch(`/api/listings?title=${title}`);
    }

    getAllByCategory(category) {
        return useFetch(`/api/listings?category=${category}`);
    }

    get(id) {
        return useAsyncData('currentListing', () => $fetch(`/api/listings/${id}`), {server: false});
    }

    create(data) {
        return useFetch( `/api/listings`, { 
            method : "post", 
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)
        });
    }

    update(id, data) {
        return useFetch( `/api/listings/${id}`, { method : "put", body: JSON.stringify(data)});
    }

    delete(id) {
        return useFetch( `/api/listings/${id}`, { method : "delete"});
    }
}

export default new ListingDataService();