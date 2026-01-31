<template>
  <section id="pricing" class="bg-[color:var(--q-bg)] py-16 sm:py-24">
    <div class="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
      <div class="text-center mb-16">
        <h2 class="text-4xl font-bold tracking-tight text-[color:var(--q-heading)] sm:text-5xl">
          Pricing
        </h2>
        <p class="mt-4 text-lg text-slate-900">
          Pilih paket yang sesuai untuk kebutuhan tim kamu.
        </p>
      </div>

      <div class="grid gap-8 lg:grid-cols-2 max-w-4xl mx-auto">
        <!-- Free Trial -->
        <div class="bg-white rounded-2xl shadow-sm border border-[color:var(--q-border)] p-8 flex flex-col hover:border-[color:var(--q-primary)] transition-colors">
          <h3 class="text-2xl font-bold text-[color:var(--q-heading)]">Free Trial (7 Days)</h3>
          <p class="mt-4 text-slate-800 flex-grow">
            Akses penuh ke semua fitur inti Quanta. Cocok untuk mencoba workflow end-to-end tanpa risiko.
          </p>
          <ul class="mt-6 space-y-4 mb-8">
            <li class="flex items-center text-sm text-slate-800">
              <svg class="mr-3 h-5 w-5 text-[color:var(--q-primary)]" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                 <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
              </svg>
              Akses Fitur Inti
            </li>
            <li class="flex items-center text-sm text-slate-800">
              <svg class="mr-3 h-5 w-5 text-[color:var(--q-primary)]" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                 <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
              </svg>
              Cancel Kapan Saja
            </li>
          </ul>
          <NuxtLink 
            to="/register" 
            class="block w-full text-center rounded-xl bg-[color:var(--q-surface)] border-2 border-[color:var(--q-primary)] px-6 py-3 text-base font-bold text-[color:var(--q-primary)] hover:bg-[color:var(--q-primary)] hover:text-white transition-all"
          >
            Mulai Free Trial
          </NuxtLink>
        </div>

        <!-- Subscription -->
        <div class="bg-white rounded-2xl shadow-lg border border-[color:var(--q-primary)] p-8 flex flex-col relative overflow-hidden">
          <div class="absolute top-0 right-0 bg-[color:var(--q-primary)] text-white text-xs font-bold px-3 py-1 rounded-bl-lg">
            POPULAR
          </div>
          <h3 class="text-2xl font-bold text-[color:var(--q-heading)]">Subscription</h3>
          <p class="mt-4 text-slate-800 flex-grow">
            Solusi jangka panjang untuk tim yang menangani project secara rutin.
          </p>
           <ul class="mt-6 space-y-4 mb-8">
            <li class="flex items-center text-sm text-slate-800">
              <svg class="mr-3 h-5 w-5 text-[color:var(--q-primary)]" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                 <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
              </svg>
              Penggunaan Berkelanjutan
            </li>
            <li class="flex items-center text-sm text-slate-800">
              <svg class="mr-3 h-5 w-5 text-[color:var(--q-primary)]" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                 <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
              </svg>
              Priority Support & Updates
            </li>
          </ul>
          <NuxtLink 
            to="/register" 
            class="block w-full text-center rounded-xl bg-[color:var(--q-primary)] px-6 py-3 text-base font-bold text-white shadow-lg shadow-[color:var(--q-primary)]/25 hover:bg-[color:var(--q-primary-hover)] transition-all"
          >
            Langganan Sekarang
          </NuxtLink>
        </div>
      </div>
    </div>
  </section>
</template>
