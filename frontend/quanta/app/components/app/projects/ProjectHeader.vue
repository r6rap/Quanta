<template>
  <header class="mb-8">
    <div class="flex flex-col gap-6 lg:flex-row lg:items-end lg:justify-between">
      <!-- Title Section -->
      <div>
        <h1 class="text-3xl font-bold tracking-tight text-q-heading">Projects</h1>
        <p class="mt-2 text-sm text-slate-500">Manage and automate your project workflows</p>
      </div>

      <!-- Actions Section -->
      <div class="flex flex-wrap items-center gap-3">
        <!-- Search Input -->
        <div class="relative">
          <div class="pointer-events-none absolute inset-y-0 left-0 flex items-center pl-3">
            <svg class="h-4 w-4 text-slate-400" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
            </svg>
          </div>
          <input
            :value="searchQuery"
            @input="$emit('update:searchQuery', ($event.target as HTMLInputElement).value)"
            type="text"
            placeholder="Search projects..."
            class="w-64 rounded-lg border border-slate-200 bg-white py-2 pl-10 pr-4 text-sm text-slate-700 placeholder:text-slate-400 transition-all focus:border-q-primary focus:outline-none focus:ring-2 focus:ring-q-ring"
          />
        </div>

        <!-- Filter Button -->
        <button class="flex items-center gap-2 rounded-lg border border-slate-200 bg-white px-4 py-2 text-sm font-medium text-slate-600 transition-colors hover:bg-slate-50 hover:text-slate-900 focus:outline-none focus:ring-2 focus:ring-q-ring">
          <svg class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M3 4a1 1 0 011-1h16a1 1 0 011 1v2.586a1 1 0 01-.293.707l-6.414 6.414a1 1 0 00-.293.707V17l-4 4v-6.586a1 1 0 00-.293-.707L3.293 7.293A1 1 0 013 6.586V4z" />
          </svg>
          Filter
        </button>

        <!-- Sort Dropdown -->
        <div class="relative">
          <button 
            @click="showSortDropdown = !showSortDropdown"
            class="flex items-center gap-2 rounded-lg border border-slate-200 bg-white px-4 py-2 text-sm font-medium text-slate-600 transition-colors hover:bg-slate-50 hover:text-slate-900 focus:outline-none focus:ring-2 focus:ring-q-ring"
          >
            <svg class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M3 4h13M3 8h9m-9 4h6m4 0l4-4m0 0l4 4m-4-4v12" />
            </svg>
            Sort
            <svg class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M19 9l-7 7-7-7" />
            </svg>
          </button>
          <!-- Sort Dropdown Menu -->
          <Transition
            enter-active-class="transition ease-out duration-100"
            enter-from-class="transform opacity-0 scale-95"
            enter-to-class="transform opacity-100 scale-100"
            leave-active-class="transition ease-in duration-75"
            leave-from-class="transform opacity-100 scale-100"
            leave-to-class="transform opacity-0 scale-95"
          >
            <div v-if="showSortDropdown" class="absolute right-0 z-10 mt-2 w-48 origin-top-right rounded-lg border border-slate-200 bg-white py-1 shadow-lg">
              <button 
                v-for="option in sortOptions" 
                :key="option.value"
                @click="handleSortChange(option.value)"
                class="flex w-full items-center px-4 py-2 text-sm text-slate-700 hover:bg-slate-50"
                :class="{ 'bg-teal-50 text-q-primary': sortBy === option.value }"
              >
                {{ option.label }}
              </button>
            </div>
          </Transition>
        </div>

        <!-- New Project Button -->
        <button 
          @click="$emit('new-project')"
          class="flex items-center gap-2 rounded-lg bg-q-primary px-4 py-2 text-sm font-medium text-white shadow-sm transition-all hover:bg-q-primary-hover focus:outline-none focus:ring-2 focus:ring-q-primary focus:ring-offset-2"
        >
          <svg class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2.5">
            <path stroke-linecap="round" stroke-linejoin="round" d="M12 4v16m8-8H4" />
          </svg>
          New Project
        </button>
      </div>
    </div>

    <!-- View Mode Toggle -->
    <div class="mt-6 flex items-center justify-between">
      <div class="flex items-center rounded-lg border border-slate-200 bg-white p-1">
        <button
          @click="$emit('update:viewMode', 'grid')"
          class="flex items-center gap-2 rounded-md px-3 py-1.5 text-sm font-medium transition-colors"
          :class="viewMode === 'grid' ? 'bg-q-primary text-white shadow-sm' : 'text-slate-600 hover:text-slate-900'"
        >
          <svg class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M4 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2V6zM14 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2V6zM4 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2v-2zM14 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2v-2z" />
          </svg>
          Grid
        </button>
        <button
          @click="$emit('update:viewMode', 'list')"
          class="flex items-center gap-2 rounded-md px-3 py-1.5 text-sm font-medium transition-colors"
          :class="viewMode === 'list' ? 'bg-q-primary text-white shadow-sm' : 'text-slate-600 hover:text-slate-900'"
        >
          <svg class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M4 6h16M4 12h16M4 18h16" />
          </svg>
          List
        </button>
      </div>

      <p class="text-sm text-slate-500">{{ projectCount }} projects</p>
    </div>
  </header>
</template>

<script setup lang="ts">
import { ref } from 'vue';

defineProps<{
  searchQuery: string;
  viewMode: 'grid' | 'list';
  sortBy: string;
  projectCount: number;
}>();

const emit = defineEmits<{
  (e: 'update:searchQuery', value: string): void;
  (e: 'update:viewMode', value: 'grid' | 'list'): void;
  (e: 'update:sortBy', value: string): void;
  (e: 'new-project'): void;
}>();

const showSortDropdown = ref(false);

const sortOptions = [
  { label: 'Last Updated', value: 'updated' },
  { label: 'Name (A-Z)', value: 'name' },
  { label: 'Created Date', value: 'created' },
  { label: 'Progress', value: 'progress' },
];

function handleSortChange(value: string) {
  emit('update:sortBy', value);
  showSortDropdown.value = false;
}
</script>
