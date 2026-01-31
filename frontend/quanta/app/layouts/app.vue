<script setup lang="ts">
import Sidebar from '~/components/app/Sidebar.vue'
</script>

<template>
  <div class="flex h-screen overflow-hidden bg-slate-50 font-sans text-slate-900">
    <!-- Sidebar -->
    <Sidebar class="hidden md:flex flex-shrink-0" />

    <!-- Main Content -->
    <div class="flex flex-1 flex-col overflow-hidden">
      <!-- Mobile Header placeholder (can be expanded later) -->
      <header class="md:hidden flex h-16 items-center justify-between border-b border-slate-200 bg-white px-4">
        <span class="text-lg font-bold text-slate-900">Quanta</span>
        <button class="text-slate-500">
          <svg class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16m-7 6h7" />
          </svg>
        </button>
      </header>

      <!-- Page Content -->
      <main class="flex-1 overflow-y-auto p-4 sm:p-6 lg:p-8">
        <slot />
      </main>
    </div>
  </div>
</template>
