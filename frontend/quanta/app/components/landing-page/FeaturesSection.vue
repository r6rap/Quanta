<script setup lang="ts">
const features = [
  {
    title: "Scope Control",
    description: "Mengunci ruang lingkup kerja dan mengontrol perubahan request.",
    icon: "M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"
  },
  {
    title: "Timeline & Milestones",
    description: "Atur target delivery, milestone, dan deadline agar tidak slip.",
    icon: "M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"
  },
  {
    title: "Budget Tracking",
    description: "Pantau rencana biaya vs biaya aktual secara ringkas.",
    icon: "M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
  },
  {
    title: "Risk & Issue Tracking",
    description: "Catat risiko, prioritas, dan tindakan mitigasi sebelum terlambat.",
    icon: "M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"
  },
  {
    title: "Task & Resource Management",
    description: "Bagi tugas, cek progress, dan hindari overload di tim.",
    icon: "M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z"
  },
  {
    title: "Progress Dashboard & Reports",
    description: "Status project jelas, update stakeholder lebih cepat tanpa meeting berlebihan.",
    icon: "M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"
  }
]
</script>

<template>
  <section id="features" class="bg-[color:var(--q-bg)] py-16 sm:py-24">
    <div class="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
      <div class="text-center mb-16">
        <h2 class="text-4xl font-bold tracking-tight text-[color:var(--q-heading)] sm:text-4xl">
          Features
        </h2>
        <p class="mt-4 text-xl text-slate-800 max-w-2xl mx-auto">
          Semua alat yang kamu butuhkan untuk menjaga project tetap dalam kendali.
        </p>
      </div>

      <div class="grid gap-8 sm:grid-cols-2 lg:grid-cols-3">
        <div 
          v-for="feature in features" 
          :key="feature.title"
          class="bg-white rounded-2xl p-8 shadow-sm border border-[color:var(--q-border)] hover:shadow-md transition-shadow group"
        >
          <div class="flex h-12 w-12 items-center justify-center rounded-lg bg-[color:var(--q-primary)]/10 text-[color:var(--q-primary)] group-hover:bg-[color:var(--q-primary)] group-hover:text-white transition-colors">
            <svg class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" :d="feature.icon" />
            </svg>
          </div>
          <h3 class="mt-6 text-2xl font-bold text-[color:var(--q-heading)]">
            {{ feature.title }}
          </h3>
          <p class="mt-3 text-base leading-relaxed text-[color:var(--q-text)]">
            {{ feature.description }}
          </p>
        </div>
      </div>
    </div>
  </section>
</template>
