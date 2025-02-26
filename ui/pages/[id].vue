<template>
  <div v-if="currentListing" class="edit-form py-3">
    <h2>Edit Listing</h2><br>

    <v-form ref="form" lazy-validation>
      <v-text-field
        v-model="currentListing.title"
        :rules="[(v) => !!v || 'Title is required']"
        label="Title"
        required
      ></v-text-field>

      <v-text-field
        v-model="currentListing.description"
        :rules="[(v) => !!v || 'Description is required']"
        label="Description"
        required
      ></v-text-field>

      <v-select
        v-model="currentListing.category"
        :rules="[(v) => !!v || 'Category is required']"
        label="Category"
        :items="['Property', 'Vehicle', 'Electronics', 'Furniture', 'Clothing', 'Services']"
        required
      ></v-select>

      <v-date-picker
        v-model="currentListing.date_posted"
        :rules="[(v: any) => !!v || 'Date is required']"
        label="Date"
        required
      ></v-date-picker>

      <v-text-field
        v-model="currentListing.price"
        :rules="[(v) => !!v || 'Price is required']"
        label="Price"
        required
      ></v-text-field>

      <label><strong>Status:</strong></label>
      {{ currentListing.active ? "Active" : "Not active" }}

      <v-divider class="my-5"></v-divider>

      <v-btn v-if="currentListing.active"
        @click="updateActive(false)"
        color="primary" small class="mr-2"
      >
        Make not active
      </v-btn>

      <v-btn v-else
        @click="updateActive(true)"
        color="primary" small class="mr-2"
      >
        Make Active
      </v-btn>

      <v-btn color="error" small class="mr-2" @click="deleteListing">
        Delete
      </v-btn>

      <v-btn color="success" small @click="updateListing">
        Update
      </v-btn>
    </v-form>
  </div>

  <div v-else>
    <p>Please click on a Listing...</p>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import ListingDataService from '../services/ListingDataService';
import type { Listing } from '~/types/Listing';

const currentListing = useCurrentListingState();
const confirmOpen = useDialogStore(state => state.openConfirm);
const messageOpen = useDialogStore(state => state.openMessage);

const route = useRoute();
const router = useRouter();

const partialUpdate = ref<Listing>(currentListing.value);
const doPartialUpdate = ref(false);

const getEditListing = (listing: Listing) => {
  listing.date_posted = new Date(listing.date_posted.toString());
  return listing;
};

const loadListing = async (id: number) => {
  const { data, error } = await ListingDataService.get(id);
  if (error?.value?.statusCode == 404) {
    messageOpen('Oops', 'This listing does not exist.', [], { color: 'red' }, async () => { 
      await navigateTo("/");
    });
  }
  currentListing.value = getEditListing(data?.value?.data);
};

const deleteListing = async () => {
  let deleteListingMethod = async () => {
    const { error } = await ListingDataService.delete(currentListing.value.id)
    if (error?.value) {
      messageOpen('Error', 'The listing was not deleted!' + (' - ' + (error?.value?.data?.message ?? '')), [], { color: 'red' });
      return;
    }
    await navigateTo("/");
    messageOpen('', 'The listing was deleted successfully!', [], { color: 'green' });
  }
  confirmOpen('Delete', 'Are you sure?', [], { color: 'red' }, deleteListingMethod);
};

const updateActive = (status: boolean) => {
  partialUpdate.value = currentListing.value;
  currentListing.value.active = status;
  partialUpdate.value.active = currentListing.value.active;
  doPartialUpdate.value = true;
  updateListing();
};

const updateListing = async () => {
  const { error } = await ListingDataService.update(currentListing.value.id, doPartialUpdate.value ? partialUpdate.value : currentListing.value)
  doPartialUpdate.value = false;
  if (error?.value) {
    messageOpen('Error', 'The listing was not updated successfully!' + (' - ' + (error?.value?.data?.message ?? '')), error?.value?.data?.errors, { color: 'red' });
    return;
  }
  messageOpen('', 'The listing was updated successfully!', [], { color: 'green' });
};

const onMountedCallback = async () => {
  await loadListing(+route.params.id);
};

onMounted(onMountedCallback);
</script>
