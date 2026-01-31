<script setup lang="ts">
interface Props {
  modelValue: string
  label?: string
  type?: string
  placeholder?: string
  error?: string
  name?: string
  autocomplete?: string
  required?: boolean
}

const props = withDefaults(defineProps<Props>(), {
  type: 'text',
  modelValue: '',
  required: false
})

const emit = defineEmits(['update:modelValue'])

const updateValue = (event: Event) => {
  const target = event.target as HTMLInputElement
  emit('update:modelValue', target.value)
}
</script>

<template>
  <div class="flex flex-col gap-1.5">
    <label v-if="label" :for="name" class="text-sm font-medium text-[color:var(--q-heading)]">
      {{ label }}
    </label>
    
    <input
      :id="name"
      :name="name"
      :type="type"
      :value="modelValue"
      :placeholder="placeholder"
      :autocomplete="autocomplete"
      :required="required"
      class="w-full rounded-lg border bg-[color:var(--q-surface)] px-4 py-2.5 text-sm outline-none transition-all duration-200 placeholder:text-[color:var(--q-muted)]"
      :class="[
        error 
          ? 'border-[color:var(--q-danger)] focus:border-[color:var(--q-danger)] focus:ring-[color:var(--q-danger)]/20 focus:ring-4' 
          : 'border-[color:var(--q-border)] focus:border-[color:var(--q-primary)] focus:ring-[color:var(--q-ring)] focus:ring-4'
      ]"
      @input="updateValue"
    />
    
    <span v-if="error" class="text-xs text-[color:var(--q-danger)]">
      {{ error }}
    </span>
  </div>
</template>
