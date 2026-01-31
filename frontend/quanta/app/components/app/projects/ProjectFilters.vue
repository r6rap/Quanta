<template>
  <aside class="hidden w-56 flex-shrink-0 lg:block">
    <nav class="sticky top-6 space-y-1">
      <div class="mb-3 px-3 text-xs font-semibold uppercase tracking-wider text-slate-400">Status</div>
      <button
        v-for="filter in statusFilters"
        :key="filter.id"
        @click="$emit('update:activeFilter', filter.id)"
        class="group flex w-full items-center gap-3 rounded-lg px-3 py-2 text-sm font-medium transition-all duration-200"
        :class="activeFilter === filter.id 
          ? 'bg-teal-50 text-q-primary' 
          : 'text-slate-600 hover:bg-slate-50 hover:text-slate-900'"
      >
        <component 
          :is="filter.icon" 
          class="h-5 w-5 transition-colors"
          :class="activeFilter === filter.id ? 'text-q-primary' : 'text-slate-400 group-hover:text-slate-500'"
        />
        {{ filter.label }}
        <span 
          v-if="filter.count !== undefined"
          class="ml-auto text-xs"
          :class="activeFilter === filter.id ? 'text-q-primary' : 'text-slate-400'"
        >
          {{ filter.count }}
        </span>
      </button>

      <div class="my-4 border-t border-slate-200"></div>

      <div class="mb-3 px-3 text-xs font-semibold uppercase tracking-wider text-slate-400">Ownership</div>
      <button
        v-for="filter in ownershipFilters"
        :key="filter.id"
        @click="$emit('update:activeFilter', filter.id)"
        class="group flex w-full items-center gap-3 rounded-lg px-3 py-2 text-sm font-medium transition-all duration-200"
        :class="activeFilter === filter.id 
          ? 'bg-teal-50 text-q-primary' 
          : 'text-slate-600 hover:bg-slate-50 hover:text-slate-900'"
      >
        <component 
          :is="filter.icon" 
          class="h-5 w-5 transition-colors"
          :class="activeFilter === filter.id ? 'text-q-primary' : 'text-slate-400 group-hover:text-slate-500'"
        />
        {{ filter.label }}
      </button>
    </nav>
  </aside>
</template>

<script setup lang="ts">
import { h } from 'vue';
import type { FilterItem } from './types';

defineProps<{
  activeFilter: string;
}>();

defineEmits<{
  (e: 'update:activeFilter', value: string): void;
}>();

// Icons
const IconAll = h('svg', { fill: 'none', viewBox: '0 0 24 24', stroke: 'currentColor', 'stroke-width': 2 }, [
  h('path', { 'stroke-linecap': 'round', 'stroke-linejoin': 'round', d: 'M4 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2V6zM14 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2V6zM4 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2v-2zM14 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2v-2z' })
]);
const IconActive = h('svg', { fill: 'none', viewBox: '0 0 24 24', stroke: 'currentColor', 'stroke-width': 2 }, [
  h('path', { 'stroke-linecap': 'round', 'stroke-linejoin': 'round', d: 'M14.752 11.168l-3.197-2.132A1 1 0 0010 9.87v4.263a1 1 0 001.555.832l3.197-2.132a1 1 0 000-1.664z' }),
  h('path', { 'stroke-linecap': 'round', 'stroke-linejoin': 'round', d: 'M21 12a9 9 0 11-18 0 9 9 0 0118 0z' })
]);
const IconDraft = h('svg', { fill: 'none', viewBox: '0 0 24 24', stroke: 'currentColor', 'stroke-width': 2 }, [
  h('path', { 'stroke-linecap': 'round', 'stroke-linejoin': 'round', d: 'M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z' })
]);
const IconArchive = h('svg', { fill: 'none', viewBox: '0 0 24 24', stroke: 'currentColor', 'stroke-width': 2 }, [
  h('path', { 'stroke-linecap': 'round', 'stroke-linejoin': 'round', d: 'M5 8h14M5 8a2 2 0 110-4h14a2 2 0 110 4M5 8v10a2 2 0 002 2h10a2 2 0 002-2V8m-9 4h4' })
]);
const IconTemplate = h('svg', { fill: 'none', viewBox: '0 0 24 24', stroke: 'currentColor', 'stroke-width': 2 }, [
  h('path', { 'stroke-linecap': 'round', 'stroke-linejoin': 'round', d: 'M4 5a1 1 0 011-1h14a1 1 0 011 1v2a1 1 0 01-1 1H5a1 1 0 01-1-1V5zM4 13a1 1 0 011-1h6a1 1 0 011 1v6a1 1 0 01-1 1H5a1 1 0 01-1-1v-6zM16 13a1 1 0 011-1h2a1 1 0 011 1v6a1 1 0 01-1 1h-2a1 1 0 01-1-1v-6z' })
]);
const IconStar = h('svg', { fill: 'none', viewBox: '0 0 24 24', stroke: 'currentColor', 'stroke-width': 2 }, [
  h('path', { 'stroke-linecap': 'round', 'stroke-linejoin': 'round', d: 'M11.049 2.927c.3-.921 1.603-.921 1.902 0l1.519 4.674a1 1 0 00.95.69h4.915c.969 0 1.371 1.24.588 1.81l-3.976 2.888a1 1 0 00-.363 1.118l1.518 4.674c.3.922-.755 1.688-1.538 1.118l-3.976-2.888a1 1 0 00-1.176 0l-3.976 2.888c-.783.57-1.838-.197-1.538-1.118l1.518-4.674a1 1 0 00-.363-1.118l-3.976-2.888c-.784-.57-.38-1.81.588-1.81h4.914a1 1 0 00.951-.69l1.519-4.674z' })
]);
const IconUser = h('svg', { fill: 'none', viewBox: '0 0 24 24', stroke: 'currentColor', 'stroke-width': 2 }, [
  h('path', { 'stroke-linecap': 'round', 'stroke-linejoin': 'round', d: 'M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z' })
]);
const IconUsers = h('svg', { fill: 'none', viewBox: '0 0 24 24', stroke: 'currentColor', 'stroke-width': 2 }, [
  h('path', { 'stroke-linecap': 'round', 'stroke-linejoin': 'round', d: 'M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z' })
]);

// Filters
const statusFilters: FilterItem[] = [
  { id: 'all', label: 'All Projects', icon: IconAll, count: 8 },
  { id: 'active', label: 'Active', icon: IconActive, count: 5 },
  { id: 'draft', label: 'Draft', icon: IconDraft, count: 2 },
  { id: 'archived', label: 'Archived', icon: IconArchive, count: 1 },
  { id: 'templates', label: 'Templates', icon: IconTemplate },
  { id: 'favorites', label: 'Favorites', icon: IconStar, count: 3 },
];

const ownershipFilters: FilterItem[] = [
  { id: 'my-projects', label: 'My Projects', icon: IconUser },
  { id: 'shared', label: 'Shared with Me', icon: IconUsers },
];
</script>
