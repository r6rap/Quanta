<script setup lang="ts">
const problems = [
  "Objectives tidak jelas di awal",
  "Scope creep (request tambahan tanpa kontrol)",
  "Komunikasi tercecer di banyak tempat",
  "Risiko muncul mendadak dan bikin panik",
  "Timeline & cost meleset dari rencana",
  "Resource terbatas dan overload"
]
</script>

<template>
  <section class="bg-white py-16 sm:py-24">
    <div class="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
      <div class="text-center">
        <h2 class="text-4xl font-bold tracking-tight text-slate-900 sm:text-5xl">
          Project sering telat bukan karena timnya,<br />
          tapi karena
          <span class="text-cyan-600">
            sistemnya tidak rapi.
          </span>
        </h2>
      </div>
      <div class="mt-12 grid gap-8 sm:grid-cols-2 lg:grid-cols-3">
        <div 
          v-for="(problem, index) in problems" 
          :key="index"
          class="flex items-start rounded-xl border border-[color:var(--q-border)] bg-[color:var(--q-bg)] p-6 transition-all hover:shadow-md"
        >
          <div class="flex-shrink-0">
            <svg class="h-6 w-6 text-[color:var(--q-danger)]" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
            </svg>
          </div>
          <div class="ml-4">
            <p class="text-base font-medium text-[color:var(--q-text)]">
              {{ problem }}
            </p>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>
