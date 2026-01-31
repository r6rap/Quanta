import { useAuthStore } from '@/stores/auth'

export async function refreshAuthToken() {
    const authStore = useAuthStore()
    const config = useRuntimeConfig()

    const res: any = await $fetch('/quanta/auth/refresh', {
        method: 'POST',
        baseURL: config.public.apiBase,
        credentials: 'include'
    })

    authStore.accessToken = res.data[0]
}
