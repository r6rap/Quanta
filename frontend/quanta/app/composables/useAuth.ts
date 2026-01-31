import { useAuthStore } from '@/stores/auth'

interface User {
    id: number
    email: string
    name: string
}

interface LoginPayload {
    email: string
    password: string
}

interface RegisterPayload {
    name: string
    email: string
    password: string
}

export interface ApiError {
  code: string
  message: string
  details?: any
}

function extractApiError(err: any): ApiError {
  if (err?.data?.error) {
    return err.data.error
  }

  return {
    code: 'UNKNOWN_ERROR',
    message: 'Unexpected error occurred',
    details: err
  }
}



export function useAuth() {

  const login = async (payload: LoginPayload) => {
    const config = useRuntimeConfig()
    const authStore = useAuthStore()

    try {
      const res: any = await $fetch('/quanta/auth/login', {
      method: 'POST',
      baseURL: config.public.apiBase,
      credentials: 'include',
      body: payload
    })

    const data = res.data[0]

    authStore.setAuth({
      user: data.user,
      accessToken: data.access_token,
      res: res
    })
    } catch (error) {
      const apiError = extractApiError(error)
      throw apiError
    }
  }

  const register = async (payload: RegisterPayload) => {
    const config = useRuntimeConfig()
    const authStore = useAuthStore()

    const res: any = await $fetch('/quanta/auth/register', {
      method: 'POST',
      baseURL: config.public.apiBase,
      credentials: 'include',
      body: payload
    })

    const data = res.data[0]

    authStore.setAuth({
      user: data.user,
      accessToken: data.access_token,
      res: res
    })
  }

  const logout = async () => {
    const config = useRuntimeConfig()
    const authStore = useAuthStore()

    await $fetch('/quanta/auth/logout', {
      method: 'POST',
      baseURL: config.public.apiBase,
      credentials: 'include'
    })

    authStore.clearAuth()
    navigateTo('/login')
  }

  return { login, register, logout }
}
