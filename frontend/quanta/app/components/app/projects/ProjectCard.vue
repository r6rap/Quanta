<template>
  <div
    class="group relative overflow-hidden rounded-xl bg-white shadow-sm ring-1 ring-slate-100 transition-all duration-300 hover:shadow-lg hover:ring-q-primary/30"
    :class="{ 'ring-2 ring-q-primary/50': project.workflow.state === 'Running' }"
  >
    <!-- Top Section -->
    <div class="p-5 pb-3">
      <div class="flex items-start justify-between gap-3">
        <div class="min-w-0 flex-1">
          <h3 class="truncate text-base font-semibold text-q-heading group-hover:text-q-primary transition-colors">
            {{ project.name }}
          </h3>
          <div class="mt-1 flex items-center gap-2">
            <span
              class="inline-flex items-center rounded-full px-2 py-0.5 text-xs font-medium"
              :class="statusStyles[project.status]"
            >
              {{ project.status }}
            </span>
            <span
              v-if="project.priority === 'high'"
              class="inline-flex items-center gap-1 text-xs font-medium text-rose-600"
            >
              <svg class="h-3 w-3" fill="currentColor" viewBox="0 0 20 20">
                <path fill-rule="evenodd" d="M5.293 9.707a1 1 0 010-1.414l4-4a1 1 0 011.414 0l4 4a1 1 0 01-1.414 1.414L11 7.414V15a1 1 0 11-2 0V7.414L6.707 9.707a1 1 0 01-1.414 0z" clip-rule="evenodd" />
              </svg>
              High
            </span>
          </div>
        </div>
        <!-- Favorite Toggle -->
        <button 
          @click="$emit('toggle-favorite', project.id)"
          class="text-slate-300 hover:text-amber-400 transition-colors"
        >
          <svg class="h-5 w-5" :fill="project.favorite ? 'currentColor' : 'none'" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2" :class="{ 'text-amber-400': project.favorite }">
            <path stroke-linecap="round" stroke-linejoin="round" d="M11.049 2.927c.3-.921 1.603-.921 1.902 0l1.519 4.674a1 1 0 00.95.69h4.915c.969 0 1.371 1.24.588 1.81l-3.976 2.888a1 1 0 00-.363 1.118l1.518 4.674c.3.922-.755 1.688-1.538 1.118l-3.976-2.888a1 1 0 00-1.176 0l-3.976 2.888c-.783.57-1.838-.197-1.538-1.118l1.518-4.674a1 1 0 00-.363-1.118l-3.976-2.888c-.784-.57-.38-1.81.588-1.81h4.914a1 1 0 00.951-.69l1.519-4.674z" />
          </svg>
        </button>
      </div>

      <!-- Description -->
      <p class="mt-3 line-clamp-2 text-sm text-slate-500">
        {{ project.description }}
      </p>

      <!-- Workflow Summary -->
      <div class="mt-4 flex items-center gap-3 rounded-lg bg-slate-50 px-3 py-2">
        <div class="flex items-center gap-1.5">
          <svg class="h-4 w-4 text-q-primary" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M13 10V3L4 14h7v7l9-11h-7z" />
          </svg>
          <span class="text-xs font-medium text-slate-600">{{ project.workflow.automations }} automations</span>
        </div>
        <span class="text-slate-300">·</span>
        <span
          class="inline-flex items-center gap-1 text-xs font-medium"
          :class="workflowStateStyles[project.workflow.state]"
        >
          <span class="relative flex h-2 w-2">
            <span v-if="project.workflow.state === 'Running'" class="absolute inline-flex h-full w-full animate-ping rounded-full bg-emerald-400 opacity-75"></span>
            <span class="relative inline-flex h-2 w-2 rounded-full" :class="workflowDotStyles[project.workflow.state]"></span>
          </span>
          {{ project.workflow.state }}
        </span>
      </div>

      <!-- Last Run Info -->
      <p class="mt-2 text-xs text-slate-400">
        Last run: {{ project.workflow.lastRun }} · 
        <span :class="project.workflow.lastResult === 'Success' ? 'text-emerald-600' : 'text-rose-600'">
          {{ project.workflow.lastResult }}
        </span>
      </p>
    </div>

    <!-- Bottom Section -->
    <div class="border-t border-slate-100 bg-slate-50/50 px-5 py-3">
      <div class="flex items-center justify-between">
        <!-- Progress -->
        <div class="flex-1">
          <div class="flex items-center justify-between text-xs">
            <span class="font-medium text-slate-600">{{ project.tasksCompleted }} / {{ project.tasksTotal }} tasks</span>
            <span class="text-slate-400">{{ progressPercentage }}%</span>
          </div>
          <div class="mt-1.5 h-1.5 w-full overflow-hidden rounded-full bg-slate-200">
            <div
              class="h-full rounded-full bg-q-primary transition-all duration-500"
              :style="{ width: `${progressPercentage}%` }"
            ></div>
          </div>
        </div>
      </div>

      <div class="mt-3 flex items-center justify-between">
        <!-- Team Avatars -->
        <div class="flex -space-x-2">
          <img
            v-for="(member, idx) in project.team.slice(0, 4)"
            :key="idx"
            :src="member.avatar"
            :alt="member.name"
            class="h-7 w-7 rounded-full border-2 border-white object-cover"
          />
          <div
            v-if="project.team.length > 4"
            class="flex h-7 w-7 items-center justify-center rounded-full border-2 border-white bg-slate-100 text-xs font-medium text-slate-600"
          >
            +{{ project.team.length - 4 }}
          </div>
        </div>

        <!-- Updated Time -->
        <span class="text-xs text-slate-400">{{ project.updatedAt }}</span>
      </div>
    </div>

    <!-- Hover Quick Actions -->
    <div 
      class="absolute inset-x-0 bottom-0 flex items-center justify-center gap-3 bg-white px-4 py-2.5 border-t border-slate-200 shadow-lg translate-y-full transition-transform duration-200 ease-out group-hover:translate-y-0"
    >
      <button 
        @click="$emit('open', project.id)"
        class="flex h-8 w-8 items-center justify-center rounded-lg bg-q-primary text-white transition-colors hover:bg-q-primary-hover"
        title="Open"
      >
        <svg class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
          <path stroke-linecap="round" stroke-linejoin="round" d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14" />
        </svg>
      </button>
      <button 
        @click="$emit('settings', project.id)"
        class="flex h-8 w-8 items-center justify-center rounded-lg bg-slate-100 text-slate-600 transition-colors hover:bg-slate-200 hover:text-slate-900"
        title="Settings"
      >
        <svg class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
          <path stroke-linecap="round" stroke-linejoin="round" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" />
          <path stroke-linecap="round" stroke-linejoin="round" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
        </svg>
      </button>
      <button 
        @click="$emit('duplicate', project.id)"
        class="flex h-8 w-8 items-center justify-center rounded-lg bg-slate-100 text-slate-600 transition-colors hover:bg-slate-200 hover:text-slate-900"
        title="Duplicate"
      >
        <svg class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
          <path stroke-linecap="round" stroke-linejoin="round" d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z" />
        </svg>
      </button>
      <button 
        @click="$emit('archive', project.id)"
        class="flex h-8 w-8 items-center justify-center rounded-lg bg-slate-100 text-slate-600 transition-colors hover:bg-slate-200 hover:text-slate-900"
        title="Archive"
      >
        <svg class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
          <path stroke-linecap="round" stroke-linejoin="round" d="M5 8h14M5 8a2 2 0 110-4h14a2 2 0 110 4M5 8v10a2 2 0 002 2h10a2 2 0 002-2V8m-9 4h4" />
        </svg>
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import type { Project } from './types';

const props = defineProps<{
  project: Project;
}>();

defineEmits<{
  (e: 'open', id: number): void;
  (e: 'settings', id: number): void;
  (e: 'duplicate', id: number): void;
  (e: 'archive', id: number): void;
  (e: 'toggle-favorite', id: number): void;
}>();

const progressPercentage = computed(() => 
  Math.round((props.project.tasksCompleted / props.project.tasksTotal) * 100)
);

// Status styles
const statusStyles: Record<string, string> = {
  'Active': 'bg-emerald-100 text-emerald-800',
  'Draft': 'bg-slate-100 text-slate-600',
  'Paused': 'bg-amber-100 text-amber-800',
  'Completed': 'bg-blue-100 text-blue-800',
};

const workflowStateStyles: Record<string, string> = {
  'Running': 'text-emerald-600',
  'Idle': 'text-slate-500',
  'Error': 'text-rose-600',
};

const workflowDotStyles: Record<string, string> = {
  'Running': 'bg-emerald-500',
  'Idle': 'bg-slate-400',
  'Error': 'bg-rose-500',
};
</script>

<style scoped>
/* Glow effect for running workflows */
.ring-q-primary\/50 {
  box-shadow: 0 0 0 2px rgba(20, 184, 166, 0.1), 0 0 20px rgba(20, 184, 166, 0.1);
}
</style>
