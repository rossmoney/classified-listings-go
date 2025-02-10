<template>
  <confirm ref="confirm"></confirm>
  <message ref="message"></message>
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
        :rules="[(v) => !!v || 'Date is required']"
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
import { ref, onMounted, useTemplateRef } from 'vue';
import Confirm from '../components/confirm.vue'
import Message from '../components/message.vue'
import type { ComponentExposed } from 'vue-component-type-helpers'
import { useRoute, useRouter } from 'vue-router';
import ListingDataService from '../services/ListingDataService';

const confirm = useTemplateRef<ComponentExposed<typeof Confirm>>('confirm');
const message = useTemplateRef<ComponentExposed<typeof Message>>('message');

const route = useRoute();
const router = useRouter();

const currentListing = useState('currentListing', () => (
  {
    title: '',
    description: '',
    date_posted: new Date(),
    price: '0.00',
    category: '',
  }
));

const partialUpdate = ref({});

const getEditListing = (listing) => {
  listing.date_posted = new Date(listing.date_posted);
  return listing;
};

const loadListing = async (id) => {
  const { data, status, error } = await ListingDataService.get(id);
  //console.log(data?.value, status?.value, error?.value?.statusCode, error?.value?.data);
  currentListing.value = getEditListing(data?.value?.data);
};

const deleteListing = async () => {
  if (await confirm.value?.open('Delete', 'Are you sure?', { color: 'red' })) {
    const { error } = await ListingDataService.delete(currentListing.value.id)
    if (error?.value) {
      await message.value?.open('Error', 'The listing was not deleted!' + (' - ' + (error?.value?.data?.message ?? '')), [], { color: 'red' })
      return;
    }
    router.push({ name: "/" });
    await message.value?.open('', 'The listing was deleted successfully!', [], { color: 'green' })
  }
};

const updateActive = (status) => {
  currentListing.value.active = status;
  partialUpdate.value = { active: currentListing.value.active };
  updateListing();
};

const updateListing = async () => {
  const { error } = await ListingDataService.update(currentListing.value.id, (Object.keys(partialUpdate.value).length > 0) ? partialUpdate.value : currentListing.value)
  partialUpdate.value = {};
  if (error?.value) {
    let errorMsg = 'The listing was not updated successfully!' + (' - ' + (error?.value?.data?.message ?? ''));
    await message.value?.open('Error', errorMsg, error?.value?.data?.errors, { color: 'red' })
    return;
  }
  await message.value?.open('', 'The listing was updated successfully!', [], { color: 'green' })
};

onMounted(async () => {
  await loadListing(route.params.id);
});
</script>
