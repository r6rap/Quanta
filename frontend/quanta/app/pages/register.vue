<script setup lang="ts">
import { useAuth } from '@/composables/useAuth'

definePageMeta({
  layout: 'auth'
})

const { register } = useAuth()

const auth = reactive({
  name: '',
  email: '',
  password: '',
  confirmPassword: '',
  loading: false,
  error: null as string | null
})

const errors = reactive({
  password: '',
  confirmPassword: ''
})

const handleRegister = async () => {
  // Reset errors
  errors.password = ''
  errors.confirmPassword = ''
  
  if (auth.password !== auth.confirmPassword) {
    errors.confirmPassword = 'Password tidak cocok'
    return
  }

  auth.loading = true
  
  try {
    await register({
      name: auth.name,
      email: auth.email,
      password: auth.password
    })
    navigateTo('/login')
  } catch (e: any) {
    auth.error = e.message
  } finally {
    auth.loading = false
  }
}
</script>

<template>
  <div>
    <div class="mb-6 text-center">
      <h2 class="text-2xl font-bold text-[color:var(--q-heading)]">
        Buat Akun Quanta
      </h2>
      <p class="mt-2 text-sm text-[color:var(--q-text)]">
        Mulai coba workflow project management yang lebih jelas.
      </p>
    </div>

    <form @submit.prevent="handleRegister" class="space-y-5">
      <AuthInput
        v-model="auth.name"
        label="Nama Lengkap"
        name="name"
        type="text"
        placeholder="John Doe"
        autocomplete="name"
        required
      />

      <AuthInput
        v-model="auth.email"
        label="Email Address"
        name="email"
        type="email"
        placeholder="nama@perusahaan.com"
        autocomplete="email"
        required
      />

      <AuthInput
        v-model="auth.password"
        label="Password"
        name="password"
        type="password"
        placeholder="••••••••"
        autocomplete="new-password"
        :error="errors.password"
        required
      />

      <AuthInput
        v-model="auth.confirmPassword"
        label="Konfirmasi Password"
        name="confirm-password"
        type="password"
        placeholder="••••••••"
        autocomplete="new-password"
        :error="errors.confirmPassword"
        required
      />

      <AuthButton type="submit" :loading="auth.loading" class="mt-2">
        Buat Akun
      </AuthButton>
    </form>
    
    <div class="mt-6 text-center text-sm text-[color:var(--q-muted)]">
      Sudah punya akun?
      <NuxtLink to="/login" class="font-semibold text-[color:var(--q-primary)] hover:text-[color:var(--q-primary-hover)]">
        Masuk di sini
      </NuxtLink>
    </div>
  </div>
</template>
