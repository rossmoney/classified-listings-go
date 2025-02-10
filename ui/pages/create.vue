
  <template>
    <confirm ref="confirm"></confirm>
    <message ref="message"></message>
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
          :rules="[(v) => !!v || 'Date is required']"
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
  import { ref, onMounted, useTemplateRef } from 'vue';
  import Confirm from '../components/confirm.vue'
  import Message from '../components/message.vue'
  import type { ComponentExposed } from 'vue-component-type-helpers'
  import { useRouter } from 'vue-router';
  import ListingDataService from '../services/ListingDataService';

  const confirm = useTemplateRef<ComponentExposed<typeof Confirm>>('confirm');
  const message = useTemplateRef<ComponentExposed<typeof Message>>('message');

  const newListing = {
    title: '',
    description: '',
    date_posted: new Date(),
    price: '0.00',
    category: '',
  };

  const currentListing = useState('currentListing', () => null);
  const router = useRouter();

  const updateActive = (status) => {
    currentListing.value.active = status;
  };

  const createListing = async () => {
    const { data, error } = await ListingDataService.create(currentListing.value);
    if (error?.value) {
      let errorMsg = 'The listing was not created successfully!' + (' - ' + (error?.value?.data?.message ?? ''));
      await message.value?.open('Error', errorMsg, error?.value?.data?.errors, { color: 'red' })
      return;
    }
    if (data?.value?.id) {
      router.push({ name: "id", params: { id: data?.value.id } })
    }
    await message.value?.open('', 'The listing was created successfully!', [], { color: 'green' })
  };

  onMounted(async () => {
    currentListing.value = newListing;
  });
  </script>