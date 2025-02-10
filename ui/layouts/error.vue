<template>
  <v-app dark>
    <h1 v-if="error.statusCode === 404">
      {{ pageNotFound }}
    </h1>
    <h1 v-else>
      {{ otherError }}
    </h1>
    <NuxtLink to="/">
      Home page
    </NuxtLink>
  </v-app>
</template>

<script setup>
import { computed, defineProps } from 'vue';

const props = defineProps({
  error: {
    type: Object,
    default: null
  }
});

const pageNotFound = '404 Not Found';
const otherError = 'An error occurred';

const title = computed(() => 
  props.error.statusCode === 404 ? pageNotFound : otherError
);

useHead({
  title: title.value
});
</script>

<style scoped>
h1 {
  font-size: 20px;
}
</style>
