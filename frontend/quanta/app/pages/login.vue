<script setup lang="ts">
import { useAuth } from '@/composables/useAuth'

definePageMeta({
  layout: 'auth'
})

const auth = reactive({
  email: '',
  password: '',
  loading: false,
  error: null as string | null
})

const { login } = useAuth()

const handleLogin = async () => {
  auth.loading = true
  auth.error = null
  
  try {
    await login({
        email: auth.email,
        password: auth.password
    })
    navigateTo('/app/dashboard')
  } catch (e: any) {
    console.log("Login error", e)
    auth.error = e.message
    console.log(auth.error)
  } finally {
    auth.loading = false
  }
}
</script>

<template>
  <div>
    <div class="mb-6 text-center">
      <h2 class="text-2xl font-bold text-[color:var(--q-heading)]">
        Masuk ke Quanta
      </h2>
      <p class="mt-2 text-sm text-[color:var(--q-text)]">
        Kelola project lebih rapi, terukur, dan on-track.
      </p>
    </div>

    <form @submit.prevent="handleLogin" class="space-y-5">
      <AuthInput
        v-model="auth.email"
        label="Email Address"
        name="email"
        type="email"
        placeholder="nama@perusahaan.com"
        autocomplete="email"
        required
      />

      <div class="space-y-1">
        <AuthInput
          v-model="auth.password"
          label="Password"
          name="password"
          type="password"
          placeholder="••••••••"
          autocomplete="current-password"
          required
        />
        <p v-if="auth.error" class="text-red-500 text-sm mt-2">
          {{ auth.error }}
        </p>  
        <div class="text-right">
          <a href="#" class="text-xs font-medium text-[color:var(--q-primary)] hover:text-[color:var(--q-primary-hover)]">
            Lupa password?
          </a>
        </div>
      </div>

      <AuthButton type="submit" :loading="auth.loading" class="mt-2">
        Masuk sekarang
      </AuthButton>
    </form>
    
    <div class="mt-6 text-center text-sm text-[color:var(--q-muted)]">
      Belum punya akun?
      <NuxtLink to="/register" class="font-semibold text-[color:var(--q-primary)] hover:text-[color:var(--q-primary-hover)]">
        Daftar gratis
      </NuxtLink>
    </div>
  </div>
</template>
