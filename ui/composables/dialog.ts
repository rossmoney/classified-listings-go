import create from 'vue-zustand'

interface DialogState {
    dialog: boolean,
    type: string,
    title: string,
    message: string,
    errors: [],
    options: {
        color: string,
        width: number,
        zIndex: number
    },
    openMessage: (title: string, message: string, errors?: [], 
      options?: {}, resolve?: () => void) => void,
    openConfirm: (title: string, message: string, errors?: [], 
      options?: {}, resolve?: () => void, reject?: () => void) => void,
    agree: () => void,
    cancel: () => void,
    resolve: () => void,
    reject: () => void,
}

export const useDialogStore = create<DialogState>(set => ({
  dialog: false,
  type: 'message',
  title: '',
  message: '',
  errors: [],
  options: {
    color: 'primary',
    width: 290,
    zIndex: 200
  },
  resolve: () => {},
  reject: () => {},
  agree: () => {
    set(state => {
      if (state.resolve !== undefined) {
        state.resolve();
      }
      return {
        dialog: false,
      };
    });
  },
  cancel: () => {
    set(state => {
      if (state.reject !== undefined) {
        state.reject();
      }
      return {
        dialog: false,
      };
    });
  },
  openConfirm: (title, message, errors?: [], options?: {}, resolve?: () => void, reject?: () => void) => {
    set(state => ({ 
      dialog: true,
      type: 'confirm',
      title: title,
      message: message,
      errors: errors,
      options: { ...state.options, ...options },
      resolve: resolve,
      reject: reject,
    }));
  },
  openMessage: (title, message, errors?: [], options?: {}, resolve?: () => void) => {
    set(state => ({ 
      dialog: true,
      type: 'message',
      title: title,
      message: message,
      errors: errors,
      options: { ...state.options, ...options },
      resolve: resolve,
      reject: () => {},
    }));
  }
}));