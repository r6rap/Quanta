<script setup lang="ts">
  const scrollToSection = (e: MouseEvent, href: string) => {
    e.preventDefault()
    const element = document.querySelector(href)
    if (element) {
      element.scrollIntoView({ behavior: 'smooth' })
    }
  }
</script>

<template>
  <section id="product" class="relative overflow-hidden pt-16 pb-20 sm:pt-24 sm:pb-32">
    <div class="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8 text-center">
      <h1 class="mx-auto max-w-4xl text-4xl font-extrabold tracking-tight text-[color:var(--q-heading)] sm:text-5xl md:text-6xl lg:leading-tight">
        Kelola Project Tanpa Chaos.<br class="hidden sm:block" />
        <span class="text-[color:var(--q-primary)]">Lebih Jelas, Lebih Terstruktur.</span>
      </h1>
      
      <p class="mx-auto mt-6 max-w-2xl text-xl text-slate-900 sm:text-xl">
        Quanta membantu kamu mengatur scope, timeline, budget, dan risiko dalam satu workflow dari awal sampai selesai.
      </p>

      <div class="mt-10 flex flex-col items-center justify-center gap-4 sm:flex-row">
        <NuxtLink 
          to="/register" 
          class="inline-flex h-12 items-center justify-center rounded-xl bg-[color:var(--q-primary)] px-8 text-base font-semibold text-white shadow-lg shadow-[color:var(--q-primary)]/25 transition-all hover:translate-y-px hover:bg-[color:var(--q-primary-hover)]"
        >
          Mulai Free Trial 7 hari
        </NuxtLink>
        <a 
          @click="scrollToSection($event, '#pricing')"
          href="#pricing"
          class="inline-flex h-12 items-center justify-center rounded-xl border border-[color:var(--q-border)] bg-white px-8 text-base font-semibold text-[color:var(--q-heading)] shadow-sm transition-all hover:bg-gray-50 hover:text-[color:var(--q-primary)]"
        >
          Lihat Pricing
        </a>
      </div>

      <div class="mt-12 flex flex-wrap justify-center gap-x-8 gap-y-4 text-base font-medium text-[color:var(--q-muted)] sm:gap-x-12">
        <div class="flex items-center">
          <svg class="mr-2 h-5 w-5 text-[color:var(--q-primary)]" viewBox="0 0 20 20" fill="currentColor">
             <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd" />
          </svg>
          Kontrol Scope
        </div>
        <div class="flex items-center">
          <svg class="mr-2 h-5 w-5 text-[color:var(--q-primary)]" viewBox="0 0 20 20" fill="currentColor">
             <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd" />
          </svg>
          Pantau Timeline
        </div>
        <div class="flex items-center">
           <svg class="mr-2 h-5 w-5 text-[color:var(--q-primary)]" viewBox="0 0 20 20" fill="currentColor">
             <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd" />
          </svg>
          Track Budget
        </div>
        <div class="flex items-center">
           <svg class="mr-2 h-5 w-5 text-[color:var(--q-primary)]" viewBox="0 0 20 20" fill="currentColor">
             <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd" />
          </svg>
          Kelola Risk
        </div>
      </div>
    </div>
  </section>
</template>
