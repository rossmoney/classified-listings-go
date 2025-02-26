<template>
  <v-dialog v-model="dialog" :max-width="width" :style="{ zIndex: zIndex }" @keydown.esc="cancel">
    <v-card>
      <v-toolbar dark :color="color" dense flat>
        <v-toolbar-title class="white--text">{{ title }}</v-toolbar-title>
      </v-toolbar>
      <v-card-text v-show="!!message" class="pa-4">{{ message }}</v-card-text>
      <ul v-show="errorsDisplay">
        <li v-for="item in errorsDisplay">
          {{ item.message }}
        </li>
      </ul>
      <v-card-actions class="pt-0">
        <v-spacer></v-spacer>
        <v-btn v-if="type === 'message'" color="primary darken-1" text @click="agree">OK</v-btn>
        <v-btn v-if="type === 'confirm'" color="primary darken-1" text @click="agree">Yes</v-btn>
        <v-btn v-if="type === 'confirm'" color="grey" text @click="cancel">Cancel</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script setup lang="ts">
import {ref, watch} from 'vue';
import { useDialogStore } from '~/composables/dialog'

const state = useDialogStore();
const dialog = useDialogStore(state => state.dialog)
const type = useDialogStore(state => state.type)
const message = useDialogStore(state => state.message)
const errorsDisplay = ref<{field: string, message: string}[]>([])
const title = useDialogStore(state => state.title)
const width = useDialogStore(state => state.options.width)
const color = useDialogStore(state => state.options.color)
const zIndex = useDialogStore(state => state.options.zIndex)

watch(() => state.errors.value, (newErrors) => {
  errorsDisplay.value = [];
  newErrors?.forEach((error: {message: string, field_name: string}) => {
    errorsDisplay.value.push({ field: error.field_name, message: error.message })
  })
});

const agree = useDialogStore(state => state.agree)
const cancel = useDialogStore(state => state.cancel)
</script>