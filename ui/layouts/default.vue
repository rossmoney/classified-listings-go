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
        <ModalDialog />
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

const clipped = ref(false);
const fixed = ref(false);

const title = ref('Classified Listings - Mini Application');
</script>
