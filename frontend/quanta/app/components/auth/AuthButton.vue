<script setup lang="ts">
interface Props {
  loading?: boolean
  disabled?: boolean
  type?: 'button' | 'submit' | 'reset'
  variant?: 'primary'
}

withDefaults(defineProps<Props>(), {
  type: 'button',
  variant: 'primary',
  loading: false,
  disabled: false
})
</script>

<template>
  <button
    :type="type"
    :disabled="disabled || loading"
    class="relative flex w-full items-center justify-center rounded-xl py-2.5 text-sm font-semibold text-white transition-all duration-200 disabled:cursor-not-allowed disabled:opacity-70"
    :class="[
      variant === 'primary' 
        ? 'bg-[color:var(--q-primary)] hover:bg-[color:var(--q-primary-hover)] shadow-sm' 
        : ''
    ]"
  >
    <span v-if="loading" class="absolute left-1/2 top-1/2 -translate-x-1/2 -translate-y-1/2">
      <svg class="h-5 w-5 animate-spin text-white" fill="none" viewBox="0 0 24 24">
        <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
        <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
      </svg>
    </span>
    
    <span :class="{ 'invisible': loading }">
      <slot />
    </span>
  </button>
</template>
