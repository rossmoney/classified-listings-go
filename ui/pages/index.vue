<template>
  <v-row align="center" class="list px-3 mx-auto">

    <v-col cols="12" md="8">
      <v-card-title>Listings</v-card-title>
    </v-col>

    <v-col cols="12" md="4">
      <v-btn small @click="createListing">
        Create New Listing
      </v-btn>
    </v-col>

    <v-col cols="12" md="8">
      <v-text-field v-model="title" label="Search by Title"></v-text-field>
    </v-col>

    <v-col cols="12" md="4">
      <v-btn small @click="searchTitle">
        Search
      </v-btn>
      <v-btn small @click="resetSearch">
        Reset
      </v-btn>
    </v-col>

    <v-col cols="12" md="8">
      <v-select
        v-model="category"
        label="Category"
        :items="['Property', 'Vehicle', 'Electronics', 'Furniture', 'Clothing', 'Services']"
      ></v-select>
    </v-col>

    <v-col cols="12" md="4">
      <v-btn small @click="filterCategory">
        Filter by Category
      </v-btn>
      <v-btn small @click="resetSearch">
        Reset
      </v-btn>
    </v-col>

    <v-col cols="12" sm="12">
      <v-card class="mx-auto" tile>
        <v-data-table
          :headers="headers"
          :items="listings"
          class="elevation-1"
        >
          <template v-slot:[`item.actions`]="{ item }">
            <v-icon small class="mr-2" @click="editListing(item.id)">mdi-pencil</v-icon>
            <v-icon small class="mr-2" @click="deleteListing(item.id)">mdi-delete</v-icon>
          </template>
        </v-data-table>
      </v-card>
    </v-col>
  </v-row>
</template>
  
<script setup lang="ts">
  import { ref, onMounted } from 'vue';
  import { useRouter } from 'vue-router';
  import ListingDataService from '../services/ListingDataService';
  import { useAllListingsState } from '~/composables/states';
  import type { Listing } from '~/types/Listing';

  const listings = useAllListingsState();
  const currentListing = useCurrentListingState();
  const messageOpen = useDialogStore(state => state.openMessage);
  const confirmOpen = useDialogStore(state => state.openConfirm);

  const title = ref('');
  const category = ref('');
  const headers = ref([
    { text: 'Title', align: 'left', value: 'title' },
    { text: 'Description', value: 'description' },
    { text: 'Price', value: 'price' },
    { text: 'Category', value: 'category' },
    { text: 'Date Posted', value: 'date_posted_human' },
    { text: 'Status', value: 'status' },
    { text: '', value: 'actions', sortable: false },
  ]);

  const router = useRouter();

  const getDisplayListing = (listing: Listing) => {
    const formatter = new Intl.DateTimeFormat('en-GB', {
      year: 'numeric',
      month: '2-digit',
      day: '2-digit',
      hour12: false,
    });
    return {
      id: listing.id,
      title: listing.title.length > 30 ? listing.title.substr(0, 30) + "..." : listing.title,
      description: listing.description.length > 30 ? listing.description.substr(0, 30) + "..." : listing.description,
      status: listing.active ? "Active" : "Not active",
      date_posted: new Date(listing.date_posted.toString()),
      date_posted_human: formatter.format(Date.parse(listing.date_posted.toString())),
      price: '£' + listing.price,
      category: listing.category,
      active: listing.active,
    };
  };

  const loadListings = async () => {
    const { data, status, error } = await ListingDataService.getAll();
    //console.log(data?.value, status?.value, error?.value?.statusCode, error?.value?.data);
    if (data?.value !== null) {
      listings.value = data?.value?.data.map(getDisplayListing);
    }
  };

  const searchTitle = async () => {
    const { data, status, error } = await ListingDataService.getAllByTitle(title.value);
    if (data?.value !== null) {
      listings.value = data?.value?.data.map(getDisplayListing);
    }
  };

  const filterCategory = async () => {
    const { data, status, error } = await ListingDataService.getAllByCategory(category.value);
    if (data?.value !== null) {
      listings.value = data?.value?.data.map(getDisplayListing);
    }
  };

  const resetSearch = async () => {
    title.value = '';
    await loadListings();
  };

  const createListing = () => {
    router.push({ name: "create" });
  };

  const editListing = (id: number) => {
    router.push({ name: "id", params: { id: id } });
  };

  const deleteListing = (id: number) => {
    currentListing.value.id = id;
    let deleteListingMethod = async () => {
      const { status, error } = await ListingDataService.delete(currentListing.value.id)
      await loadListings();
      if (error?.value && error.value.statusCode !== 200) {
        messageOpen('Error', 'The listing was not deleted!' + (' - ' + (error?.value?.data?.message ?? '')), [], { color: 'red' })
      }
      if (status?.value == 'success') {
        messageOpen('', 'The listing was deleted successfully!', [], { color: 'green' })
      }
    };
    confirmOpen('Delete', 'Are you sure?', [], { color: 'red' }, deleteListingMethod);
  };

  onMounted(async () => {
    await loadListings();
  });
</script>