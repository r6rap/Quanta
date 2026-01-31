<script setup lang="ts">
const navItems = [
  { name: 'Product', href: '#product' },
  { name: 'How It Works', href: '#how-it-works' },
  { name: 'Features', href: '#features' },
  { name: 'Pricing', href: '#pricing' },
  { name: 'FAQ', href: '#faq' },
]

const isDarkSection = ref(false)

let observer: IntersectionObserver | null = null

  onMounted(() => {
    // Determine the offset for the navbar height (64px)
    // We want the intersection to happen when the element hits the top 64px of the viewport
    const bottomMargin = -(window.innerHeight - 64)
    
    const darkSections = document.querySelectorAll('#cta')

    observer = new IntersectionObserver(
      (entries) => {
        // Check if ANY of the observed sections are intersecting with the top bar
        const isIntersecting = entries.some((entry) => entry.isIntersecting)
        isDarkSection.value = isIntersecting
      },
      {
        root: null,
        // Shrink the root's bottom edge up so only the top 64px is the "active" capture area
        rootMargin: `0px 0px ${bottomMargin}px 0px`,
        threshold: 0,
      }
    )

    darkSections.forEach((el) => observer?.observe(el))
  })

  onBeforeUnmount(() => {
    observer?.disconnect()
  })

  const scrollToSection = (e: MouseEvent, href: string) => {
    e.preventDefault()
    const element = document.querySelector(href)
    if (element) {
      element.scrollIntoView({ behavior: 'smooth' })
    }
  }
  </script>

  <template>
    <div class="min-h-screen bg-[color:var(--q-bg)] font-sans text-[color:var(--q-text)] selection:bg-[color:var(--q-primary)] selection:text-white">
      <!-- Navbar -->
      <header 
        class="fixed top-0 left-0 right-0 z-50 backdrop-blur-md border-b shadow-sm transition-all duration-300"
        :class="[
          isDarkSection 
            ? 'bg-transparent border-white/5' 
            : 'bg-[color:var(--q-surface)]/80 border-black/5'
        ]"
      >
        <div class="mx-auto flex h-16 max-w-7xl items-center justify-between px-4 sm:px-6 lg:px-8">
          <!-- Logo -->
          <div class="flex items-center">
            <NuxtLink 
              to="/" 
              class="text-2xl font-bold transition-colors duration-300"
              :class="[ isDarkSection ? 'text-white' : 'text-[color:var(--q-heading)]' ]"
            >
              Quanta
            </NuxtLink>
          </div>

          <!-- Desktop Nav -->
          <nav class="hidden md:flex items-center gap-8">
            <a 
              v-for="item in navItems" 
              :key="item.name" 
              :href="item.href"
              @click="scrollToSection($event, item.href)"
              class="text-base font-medium px-3 py-2 rounded-lg transition-colors duration-300"
              :class="[
                isDarkSection 
                  ? 'text-white/90 hover:bg-white/10 hover:text-white' 
                  : 'text-[color:var(--q-heading)] hover:bg-black/5 hover:text-[color:var(--q-heading)]'
              ]"
            >
              {{ item.name }}
            </a>
          </nav>

          <!-- Auth Buttons -->
          <div class="flex items-center gap-4">
            <NuxtLink 
              to="/login"
              class="hidden text-base font-medium transition-colors duration-300 sm:block"
              :class="[
                isDarkSection
                  ? 'text-white hover:text-white/80'
                  : 'text-[color:var(--q-heading)] hover:text-[color:var(--q-primary)]'
              ]"
            >
              Login
            </NuxtLink>
            <NuxtLink 
              to="/register"
              class="rounded-lg bg-[color:var(--q-primary)] px-4 py-2 text-base font-semibold text-white shadow-sm hover:bg-[color:var(--q-primary-hover)] transition-all"
            >
              Start Trial
            </NuxtLink>
          </div>
        </div>
      </header>

    <!-- Main Content -->
    <main class="pt-16">
      <slot />
    </main>

    <!-- Footer -->
    <footer class="border-t border-[color:var(--q-border)] bg-[color:var(--q-surface)] pt-16 pb-8">
      <div class="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
        <div class="grid grid-cols-2 gap-8 md:grid-cols-4 lg:grid-cols-5">
          <div class="col-span-2 lg:col-span-2">
            <h3 class="text-lg font-bold text-[color:var(--q-heading)]">Quanta</h3>
            <p class="mt-4 text-sm text-[color:var(--q-muted)] max-w-xs">
              Platform manajemen proyek end-to-end yang bantu tim mengelola project lebih rapi, terukur, dan on-track.
            </p>
          </div>
          <div>
            <h4 class="text-sm font-semibold text-[color:var(--q-heading)]">Product</h4>
            <ul class="mt-4 space-y-2 text-sm text-[color:var(--q-muted)]">
              <li><a href="#features" @click="scrollToSection($event, '#features')" class="hover:text-[color:var(--q-primary)]">Features</a></li>
              <li><a href="#pricing" @click="scrollToSection($event, '#pricing')" class="hover:text-[color:var(--q-primary)]">Pricing</a></li>
              <li><a href="#faq" @click="scrollToSection($event, '#faq')" class="hover:text-[color:var(--q-primary)]">FAQ</a></li>
            </ul>
          </div>
          <div>
            <h4 class="text-sm font-semibold text-[color:var(--q-heading)]">Company</h4>
            <ul class="mt-4 space-y-2 text-sm text-[color:var(--q-muted)]">
              <li><a href="#" class="hover:text-[color:var(--q-primary)]">About</a></li>
              <li><a href="#" class="hover:text-[color:var(--q-primary)]">Blog</a></li>
              <li><a href="#" class="hover:text-[color:var(--q-primary)]">Careers</a></li>
            </ul>
          </div>
          <div>
            <h4 class="text-sm font-semibold text-[color:var(--q-heading)]">Legal</h4>
            <ul class="mt-4 space-y-2 text-sm text-[color:var(--q-muted)]">
              <li><a href="#" class="hover:text-[color:var(--q-primary)]">Privacy</a></li>
              <li><a href="#" class="hover:text-[color:var(--q-primary)]">Terms</a></li>
            </ul>
          </div>
        </div>
        <div class="mt-12 border-t border-[color:var(--q-border)] pt-8 text-center text-sm text-[color:var(--q-muted)]">
          &copy; {{ new Date().getFullYear() }} Quanta. All rights reserved.
        </div>
      </div>
    </footer>
  </div>
</template>
