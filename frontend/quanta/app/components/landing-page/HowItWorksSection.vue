<script setup lang="ts">
const steps = [
  {
    title: "Initiation",
    description: "Tentukan tujuan, output, dan stakeholder utama"
  },
  {
    title: "Planning",
    description: "Susun scope, timeline, budget, resource, dan risk plan"
  },
  {
    title: "Execution",
    description: "Jalankan task, koordinasi tim, dan progress update"
  },
  {
    title: "Monitoring",
    description: "Pantau KPI, issue, perubahan scope, dan status project"
  },
  {
    title: "Closing",
    description: "Final delivery, sign-off, dan evaluasi hasil"
  }
]
</script>

<template>
  <section id="how-it-works" class="py-16 sm:py-24 bg-white">
    <div class="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
      <div class="text-center mb-16">
        <h2 class="text-lg font-semibold uppercase tracking-wide text-[color:var(--q-primary)]">How It Works</h2>
        <p class="mt-2 text-4xl font-bold leading-8 tracking-tight text-[color:var(--q-heading)] sm:text-5xl">
          Lifecycle Project Management
        </p>
      </div>

      <div class="relative">
        <!-- Connecting Line (Desktop) -->
        <div class="absolute top-12 left-0 hidden w-full border-t-2 border-dashed border-[color:var(--q-border)] lg:block" aria-hidden="true"></div>

        <div class="grid gap-8 lg:grid-cols-5">
          <div v-for="(step, index) in steps" :key="index" class="relative flex flex-col items-center text-center">
            <!-- Number Circle -->
            <div class="flex h-24 w-24 items-center justify-center rounded-full bg-white border-4 border-[color:var(--q-primary)] text-2xl font-bold text-[color:var(--q-primary)] shadow-sm relative z-10">
              {{ index + 1 }}
            </div>
            
            <h3 class="mt-6 text-2xl font-bold text-[color:var(--q-heading)]">
              {{ step.title }}
            </h3>
            <p class="mt-2 text-base text-slate-800 px-2">
              {{ step.description }}
            </p>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>
