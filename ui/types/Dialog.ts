interface DialogOptions {
    color: string;
    width: number;
    zIndex: number;
}
  
interface Dialog {
    title: string;
    message: string;
    errors: null | Record<string, any>;
    options: DialogOptions;
}

export interface ConfirmState {
    dialog: Dialog;
    open: boolean;
    result: boolean;
}

export interface MessageState {
    dialog: Dialog;
    open: boolean;
}