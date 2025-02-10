<template>
  <v-app dark>
    <v-app-bar
      :clipped-left="clipped"
      fixed
      app
    >
      <v-toolbar-title>{{ title }}</v-toolbar-title>
      <v-btn text to="/">Back</v-btn>
    </v-app-bar>
    <v-main>
      <v-container>
        <NuxtPage/>
      </v-container>
    </v-main>
    <v-footer
      :absolute="!fixed"
      app
    >
    </v-footer>
  </v-app>
</template>

<script setup>
import { ref } from 'vue';
import { useRoute } from 'vue-router';

const clipped = ref(false);
const fixed = ref(false);

const route = useRoute();

const title = ref('Classified Listings - Mini Application');
</script>
