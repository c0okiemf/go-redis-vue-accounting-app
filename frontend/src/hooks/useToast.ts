import ToastNotification from '@/components/ToastNotification.vue';
import type { ToastProps } from '@/types/toast';
import { createApp, h } from 'vue';

export function useToast() {
  function createToast(props: ToastProps) {
    const div = document.createElement('div');
    document.body.appendChild(div);

    const app = createApp({
      render() {
        return h(ToastNotification, props);
      }
    });

    app.mount(div);

    setTimeout(
      () => {
        app.unmount();
        document.body.removeChild(div);
      },
      (props.duration || 5000) + 500
    ); // 500ms buffer for fade out animation
  }

  return { createToast };
}
