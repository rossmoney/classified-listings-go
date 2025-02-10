<script setup>
const error = useError();

const handleError = () => {
  clearError({
    redirect: '/',
  });
};
</script>

<template>
    <NuxtLayout>
      <div class="prose">
        <h1>Dang</h1>
              <p>
                <strong>{{ error.message }}</strong>
              </p>
        <p>It looks like something broke.</p>
        <p>Sorry about that.</p>
        <p>
              Go back to 
              <a @click="handleError">
                listings.
              </a>
            </p>
      </div>
    </NuxtLayout>
</template>
  