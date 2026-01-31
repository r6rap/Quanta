<template>
  <div class="flex flex-col items-center justify-center rounded-xl border-2 border-dashed border-slate-200 bg-white py-16">
    <div class="flex h-16 w-16 items-center justify-center rounded-full bg-slate-100">
      <svg class="h-8 w-8 text-slate-400" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="1.5">
        <path stroke-linecap="round" stroke-linejoin="round" d="M3 7v10a2 2 0 002 2h14a2 2 0 002-2V9a2 2 0 00-2-2h-6l-2-2H5a2 2 0 00-2 2z" />
      </svg>
    </div>
    <h3 class="mt-4 text-lg font-semibold text-q-heading">No projects yet</h3>
    <p class="mt-1 text-sm text-slate-500">Create your first workflow-driven project</p>
    <button 
      @click="$emit('create')"
      class="mt-6 flex items-center gap-2 rounded-lg bg-q-primary px-4 py-2 text-sm font-medium text-white shadow-sm transition-all hover:bg-q-primary-hover focus:outline-none focus:ring-2 focus:ring-q-primary focus:ring-offset-2"
    >
      <svg class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2.5">
        <path stroke-linecap="round" stroke-linejoin="round" d="M12 4v16m8-8H4" />
      </svg>
      Create Project
    </button>
  </div>
</template>

<script setup lang="ts">
defineEmits<{
  (e: 'create'): void;
}>();
</script>
