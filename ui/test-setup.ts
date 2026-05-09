import { ref, computed, watch, onMounted, onBeforeUnmount, reactive, toRef } from 'vue'

Object.assign(globalThis, {
  ref,
  computed,
  watch,
  onMounted,
  onBeforeUnmount,
  reactive,
  toRef,
  useRuntimeConfig: () => ({
    public: { wsBase: 'ws://localhost:3000' },
  }),
})
