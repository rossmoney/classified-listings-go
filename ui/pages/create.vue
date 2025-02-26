
  <template>
    <div class="edit-form py-3">
      <h2>New Listing</h2><br>
  
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
          label="Price (£)"
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
  
        <v-btn color="success" small @click="createListing">
          Create
        </v-btn>
      </v-form>
    </div>
  </template>
  
  <script setup lang="ts">
  import { onMounted } from 'vue';
  import { useRouter } from 'vue-router';
  import ListingDataService from '../services/ListingDataService';
  import { useCurrentListingState } from '~/composables/states';
  import type { Listing } from '~/types/Listing';

  const newListing = {
    id: 0,
    title: '',
    description: '',
    date_posted: new Date(),
    date_posted_human: '',
    price: '0.00',
    category: '',
    active: false,
  } as Listing;

  const currentListing = useCurrentListingState();
  const messageOpen = useDialogStore(state => state.openMessage);

  const router = useRouter();

  const updateActive = (status: boolean) => {
    currentListing.value.active = status;
  };

  const createListing = async () => {
    const { data, error } = await ListingDataService.create(currentListing.value);
    if (error?.value) {
      messageOpen('Error', 
        'The listing was not created successfully!' + (' - ' + (error?.value?.data?.message ?? '')), 
        error?.value?.data?.errors, 
        { color: 'red'},
      );
      return;
    }
    if (data?.value?.data.id) {
      router.push({ name: "id", params: { id: data?.value.data.id } })
    }
    messageOpen('', 'The listing was created successfully!', [], { color: 'green' })
  };

  onMounted(async () => {
    currentListing.value = newListing;
  });
  </script>