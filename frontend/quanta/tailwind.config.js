/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
        "./app/components/**/*.{js,vue,ts}",
        "./app/layouts/**/*.vue",
        "./app/pages/**/*.vue",
        "./app/plugins/**/*.{js,ts}",
        "./app/app.vue",
        "./nuxt.config.{js,ts}",
        "./app/**/*.{js,vue,ts}", // Fallback for any other files in app/
    ],
    theme: {
        extend: {
            colors: {
                // Define colors mapping to CSS variables
                q: {
                    primary: 'var(--q-primary)',
                    'primary-hover': 'var(--q-primary-hover)',
                    heading: 'var(--q-heading)',
                    bg: 'var(--q-bg)',
                    surface: 'var(--q-surface)',
                    border: 'var(--q-border)',
                    text: 'var(--q-text)',
                    muted: 'var(--q-muted)',
                    danger: 'var(--q-danger)',
                    ring: 'var(--q-ring)',
                }
            }
        },
    },
    plugins: [],
}
