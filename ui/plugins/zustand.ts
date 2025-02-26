import { useDialogStore } from '../composables/dialog'

export default defineNuxtPlugin((nuxtApp) => {
    if (process.server) {
      nuxtApp.hooks.hook('app:rendered', () => {
        const initialState = JSON.parse(JSON.stringify(useDialogStore.getState()))
        nuxtApp.payload.zustand = initialState
      })
    }
  
    if (process.client) {
      nuxtApp.hooks.hook('app:created', () => {
        useDialogStore.setState({
          ...useDialogStore.getState(),
          ...(typeof nuxtApp.payload.zustand === 'object' ? nuxtApp.payload.zustand : {}),
        })
      })
    }
})