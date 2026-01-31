<script setup lang="ts">
const benefits = [
  { 
    solution: "Scope Control", 
    problem: "Scope creep tanpa kendali",
    description: "Kunci ruang lingkup di awal. Setiap perubahan butuh approval, cegah request 'gratisan' yang membengkak."
  },
  { 
    solution: "Timeline & Monitoring", 
    problem: "Project molor terus menerus",
    description: "Visualisasikan deadline yang realistis. Deteksi potensi keterlambatan lebih dini sebelum menjadi bencana."
  },
  { 
    solution: "Task Management", 
    problem: "Komunikasi tim tercecer",
    description: "Pusatkan semua diskusi di satu tempat. Tidak ada lagi tugas yang terlewat atau miskomunikasi di chat app."
  },
  { 
    solution: "Risk Tracking", 
    problem: "Masalah muncul dadakan",
    description: "Identifikasi risiko sejak hari pertama. Siapkan mitigasi sehingga tim tidak panik saat masalah terjadi."
  },
  { 
    solution: "Budget Tracking", 
    problem: "Biaya operasional bocor",
    description: "Pantau burn rate budget secara real-time. Pastikan profit margin project tetap aman terjaga."
  },
  { 
    solution: "Automated Reports", 
    problem: "Stakeholder bingung progress",
    description: "Generate laporan status instan yang mudah dibaca klien. Kurangi meeting status update yang tidak perlu."
  }
]
</script>

<template>
  <section class="bg-[color:var(--q-bg)] py-20 sm:py-28">
    <div class="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
      <div class="text-center mb-16">
        <h2 class="text-4xl font-bold tracking-tight text-[color:var(--q-heading)] sm:text-4xl">
          Solusi untuk Masalah Klasik
        </h2>
        <p class="mt-4 text-xl text-slate-800 max-w-2xl mx-auto">
          Quanta mengubah kekacauan manajemen proyek menjadi sistem yang terprediksi.
        </p>
      </div>

      <div class="grid gap-8 sm:grid-cols-2 lg:grid-cols-3">
        <div 
          v-for="(item, index) in benefits" 
          :key="index"
          class="group relative overflow-hidden rounded-2xl bg-white p-8 shadow-sm ring-1 ring-black/5 transition-all hover:shadow-md hover:ring-[color:var(--q-primary)]/20"
        >
          <!-- Decoration -->
          <div class="absolute top-0 right-0 -mt-4 -mr-4 h-24 w-24 rounded-full bg-gradient-to-br from-[color:var(--q-primary)]/10 to-transparent blur-2xl transition-all group-hover:scale-150"></div>
          
          <div class="relative">
            <!-- Problem Header (Subtle) -->
            <div class="flex items-center gap-2 text-sm font-semibold uppercase tracking-wider text-red-600 opacity-100 mb-3">
              <svg class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
              </svg>
              {{ item.problem }}
            </div>

            <!-- Solution Title -->
            <h3 class="text-2xl font-bold text-[color:var(--q-heading)]">
              {{ item.solution }}
            </h3>

            <!-- Description -->
            <p class="mt-3 text-base leading-relaxed text-[color:var(--q-text)]">
              {{ item.description }}
            </p>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>
