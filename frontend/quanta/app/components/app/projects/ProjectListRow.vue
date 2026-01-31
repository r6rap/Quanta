<template>
  <li class="group px-6 py-4 transition-colors hover:bg-slate-50/50">
    <div class="grid grid-cols-12 gap-4 items-center">
      <!-- Project Name -->
      <div class="col-span-4 flex items-center gap-3">
        <button 
          @click="$emit('toggle-favorite', project.id)"
          class="text-slate-300 hover:text-amber-400 transition-colors"
        >
          <svg class="h-4 w-4" :fill="project.favorite ? 'currentColor' : 'none'" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2" :class="{ 'text-amber-400': project.favorite }">
            <path stroke-linecap="round" stroke-linejoin="round" d="M11.049 2.927c.3-.921 1.603-.921 1.902 0l1.519 4.674a1 1 0 00.95.69h4.915c.969 0 1.371 1.24.588 1.81l-3.976 2.888a1 1 0 00-.363 1.118l1.518 4.674c.3.922-.755 1.688-1.538 1.118l-3.976-2.888a1 1 0 00-1.176 0l-3.976 2.888c-.783.57-1.838-.197-1.538-1.118l1.518-4.674a1 1 0 00-.363-1.118l-3.976-2.888c-.784-.57-.38-1.81.588-1.81h4.914a1 1 0 00.951-.69l1.519-4.674z" />
          </svg>
        </button>
        <div>
          <p class="text-sm font-medium text-q-heading group-hover:text-q-primary transition-colors">{{ project.name }}</p>
          <p class="text-xs text-slate-400 mt-0.5 line-clamp-1">{{ project.description }}</p>
        </div>
      </div>

      <!-- Status -->
      <div class="col-span-1">
        <span
          class="inline-flex items-center rounded-full px-2 py-0.5 text-xs font-medium"
          :class="statusStyles[project.status]"
        >
          {{ project.status }}
        </span>
      </div>

      <!-- Workflow State -->
      <div class="col-span-2">
        <div class="flex items-center gap-2">
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
          <span class="text-xs text-slate-400">· {{ project.workflow.automations }} auto</span>
        </div>
      </div>

      <!-- Progress -->
      <div class="col-span-2">
        <div class="flex items-center gap-3">
          <div class="h-1.5 flex-1 overflow-hidden rounded-full bg-slate-200">
            <div
              class="h-full rounded-full bg-q-primary transition-all duration-500"
              :style="{ width: `${progressPercentage}%` }"
            ></div>
          </div>
          <span class="text-xs text-slate-500 whitespace-nowrap">{{ project.tasksCompleted }}/{{ project.tasksTotal }}</span>
        </div>
      </div>

      <!-- Last Activity -->
      <div class="col-span-2">
        <div class="flex items-center gap-2">
          <div class="flex -space-x-1">
            <img
              v-for="(member, idx) in project.team.slice(0, 2)"
              :key="idx"
              :src="member.avatar"
              :alt="member.name"
              class="h-5 w-5 rounded-full border border-white object-cover"
            />
          </div>
          <span class="text-xs text-slate-500">{{ project.updatedAt }}</span>
        </div>
      </div>

      <!-- Actions -->
      <div class="col-span-1 text-right">
        <div class="flex items-center justify-end gap-1 opacity-0 group-hover:opacity-100 transition-opacity">
          <button 
            @click="$emit('open', project.id)"
            class="rounded-lg p-1.5 text-slate-400 hover:bg-slate-100 hover:text-slate-600 transition-colors"
            title="Open project"
          >
            <svg class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14" />
            </svg>
          </button>
          <button 
            @click="showMenu = !showMenu"
            class="rounded-lg p-1.5 text-slate-400 hover:bg-slate-100 hover:text-slate-600 transition-colors"
            title="More actions"
          >
            <svg class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M12 5v.01M12 12v.01M12 19v.01M12 6a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2z" />
            </svg>
          </button>
        </div>
      </div>
    </div>
  </li>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue';
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

const showMenu = ref(false);

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
