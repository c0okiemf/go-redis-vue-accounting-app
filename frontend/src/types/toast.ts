export enum ToastTheme {
  Success = 'success',
  Error = 'error'
}

export interface ToastProps {
  message: string;
  duration?: number;
  theme?: ToastTheme;
}
