import { defineStore } from 'pinia'

export const useAuthStore = defineStore('auth', {
  state: () => ({
    user: null as any | null,
    accessToken: null as string | null,
    res: null as any | null
  }),

  getters: {
    isAuthenticated: (state) => !!state.accessToken
  },

  actions: {
    setAuth(payload: { user: any; accessToken: string, res: any }) {
      this.user = payload.user
      this.accessToken = payload.accessToken
      this.res = payload.res
    },

    clearAuth() {
      this.user = null
      this.accessToken = null
      this.res = null
    }
  }
})
