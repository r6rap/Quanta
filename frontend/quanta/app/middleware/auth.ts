import { refreshAuthToken } from '@/composables/useAuthRefresh'

export default defineNuxtRouteMiddleware(async () => {
  try {
    await refreshAuthToken()
  } catch {
    return navigateTo('/login')
  }
})
