<template>
  <div class="min-h-screen flex items-center justify-center p-4 bg-[color:var(--q-bg)]">
    <div class="w-full max-w-md space-y-8">
      <!-- Brand Header -->
      <div class="text-center">
        <!-- Optional Logo Placeholder -->
        <h1 class="text-3xl font-bold tracking-tight text-[color:var(--q-heading)]">
          Quanta
        </h1>
      </div>

      <!-- Main Card -->
      <div class="bg-[color:var(--q-surface)] p-6 sm:p-8 rounded-2xl shadow-md border border-[color:var(--q-border)]">
        <slot />
      </div>
      
      <!-- Footer / Back Link (Optional) -->
      <div class="text-center text-sm">
        <NuxtLink to="/" class="font-medium text-[color:var(--q-muted)] hover:text-[color:var(--q-text)] transition-colors">
          &larr; Kembali ke Beranda
        </NuxtLink>
      </div>
    </div>
  </div>
</template>
