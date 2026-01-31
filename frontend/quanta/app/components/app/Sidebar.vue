<template>
  <aside class="flex w-64 flex-col border-r border-slate-200 bg-white shadow-sm transition-all duration-300">
    <!-- Brand -->
    <div class="flex h-16 items-center px-6 border-b border-slate-100/50">
      <div class="flex items-center gap-2">
        <!-- Logo Icon (Simple Geometric Shape for Quanta) -->
        <div class="flex h-8 w-8 items-center justify-center rounded-lg bg-q-primary text-white">
          <svg class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M13 10V3L4 14h7v7l9-11h-7z" />
          </svg>
        </div>
        <span class="text-xl font-bold tracking-tight text-q-heading">Quanta</span>
      </div>
    </div>

    <!-- Main Navigation -->
    <nav class="flex-1 space-y-1 px-3 py-6 overflow-y-auto">
      <div class="mb-2 px-3 text-xs font-semibold uppercase tracking-wider text-slate-400">
        Overview
      </div>
      <NuxtLink v-for="item in mainNav" :key="item.name" :to="item.to"
        class="group flex items-center rounded-lg px-3 py-2 text-sm font-medium transition-colors duration-200"
        :class="[
          $route.path === item.to 
            ? 'bg-teal-50 text-q-primary' 
            : 'text-slate-600 hover:bg-slate-50 hover:text-slate-900'
        ]"
      >
        <component :is="item.icon" 
          class="mr-3 h-5 w-5 flex-shrink-0 transition-colors"
          :class="[$route.path === item.to ? 'text-q-primary' : 'text-slate-400 group-hover:text-slate-500']"
        />
        {{ item.name }}
      </NuxtLink>

      <div class="mt-8 mb-2 px-3 text-xs font-semibold uppercase tracking-wider text-slate-400">
        Workspace
      </div>
      <NuxtLink v-for="item in secondaryNav" :key="item.name" :to="item.to"
        class="group flex items-center rounded-lg px-3 py-2 text-sm font-medium transition-colors duration-200"
        :class="[
          $route.path === item.to 
            ? 'bg-teal-50 text-q-primary' 
            : 'text-slate-600 hover:bg-slate-50 hover:text-slate-900'
        ]"
      >
        <component :is="item.icon" 
          class="mr-3 h-5 w-5 flex-shrink-0 transition-colors"
          :class="[$route.path === item.to ? 'text-q-primary' : 'text-slate-400 group-hover:text-slate-500']"
        />
        {{ item.name }}
      </NuxtLink>
    </nav>

    <!-- Footer / User Profile -->
    <div class="border-t border-slate-200 p-4">
      <div class="flex items-center gap-3 rounded-xl border border-transparent p-2 transition-colors hover:border-slate-200 hover:bg-slate-50 cursor-pointer group">
        <img
          class="h-9 w-9 rounded-full object-cover ring-2 ring-white"
          src="https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80"
          alt="User Profile"
        />
        <div class="flex-1 overflow-hidden">
          <p class="truncate text-sm font-medium text-slate-700 group-hover:text-slate-900">Alex Morgan</p>
          <p class="truncate text-xs text-slate-500">Project Manager</p>
        </div>
        <button class="text-slate-400 hover:text-slate-600">
            <!-- Logout Icon -->
             <svg class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1" />
              </svg>
        </button>
      </div>
    </div>
  </aside>
</template>

<script setup lang="ts">
import { h } from 'vue';

// Minimal Icon Components (Inline for simplicity, typically would be imported)
const IconDashboard = h('svg', { fill: 'none', viewBox: '0 0 24 24', stroke: 'currentColor', 'stroke-width': 2 }, [
  h('path', { 'stroke-linecap': 'round', 'stroke-linejoin': 'round', d: 'M4 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2V6zM14 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2V6zM4 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2v-2zM14 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2v-2z' })
]);
const IconProjects = h('svg', { fill: 'none', viewBox: '0 0 24 24', stroke: 'currentColor', 'stroke-width': 2 }, [
  h('path', { 'stroke-linecap': 'round', 'stroke-linejoin': 'round', d: 'M3 7v10a2 2 0 002 2h14a2 2 0 002-2V9a2 2 0 00-2-2h-6l-2-2H5a2 2 0 00-2 2z' })
]);
const IconTasks = h('svg', { fill: 'none', viewBox: '0 0 24 24', stroke: 'currentColor', 'stroke-width': 2 }, [
  h('path', { 'stroke-linecap': 'round', 'stroke-linejoin': 'round', d: 'M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2m-6 9l2 2 4-4' })
]);
const IconRisks = h('svg', { fill: 'none', viewBox: '0 0 24 24', stroke: 'currentColor', 'stroke-width': 2 }, [
  h('path', { 'stroke-linecap': 'round', 'stroke-linejoin': 'round', d: 'M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z' })
]);
const IconReports = h('svg', { fill: 'none', viewBox: '0 0 24 24', stroke: 'currentColor', 'stroke-width': 2 }, [
  h('path', { 'stroke-linecap': 'round', 'stroke-linejoin': 'round', d: 'M9 17v-2m3 2v-4m3 4v-6m2 10H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z' })
]);
const IconTeam = h('svg', { fill: 'none', viewBox: '0 0 24 24', stroke: 'currentColor', 'stroke-width': 2 }, [
  h('path', { 'stroke-linecap': 'round', 'stroke-linejoin': 'round', d: 'M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z' })
]);
const IconSettings = h('svg', { fill: 'none', viewBox: '0 0 24 24', stroke: 'currentColor', 'stroke-width': 2 }, [
  h('path', { 'stroke-linecap': 'round', 'stroke-linejoin': 'round', d: 'M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z' }),
  h('path', { 'stroke-linecap': 'round', 'stroke-linejoin': 'round', d: 'M15 12a3 3 0 11-6 0 3 3 0 016 0z' })
]);

const mainNav = [
  { name: 'Dashboard', to: '/app/dashboard', icon: IconDashboard },
  { name: 'Projects', to: '/app/projects', icon: IconProjects },
  { name: 'Tasks', to: '/app/tasks', icon: IconTasks },
  { name: 'Risks', to: '/app/risks', icon: IconRisks },
  { name: 'Reports', to: '/app/reports', icon: IconReports },
];

const secondaryNav = [
  { name: 'Team', to: '/app/team', icon: IconTeam },
  { name: 'Settings', to: '/app/settings', icon: IconSettings },
];
</script>
