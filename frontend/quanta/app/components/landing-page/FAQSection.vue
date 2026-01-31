<script setup lang="ts">
const faqs = [
  {
    question: "Apa yang bisa saya coba selama Free Trial Quanta (7 hari)?",
    answer: "Anda mendapatkan akses penuh ke seluruh fitur premium Quanta tanpa batasan, sehingga bisa merasakan pengalaman mengelola project dari awal hingga selesai."
  },
  {
    question: "Bagaimana sistem setelah Free Trial berakhir dan kenapa tidak otomatis ditagih?",
    answer: "Setelah 7 hari, akun Anda akan menjadi read-only kecuali Anda memilih untuk berlangganan. Kami tidak meminta kartu kredit di awal, jadi tidak ada tagihan kejutan."
  },
  {
    question: "Kenapa Quanta cocok untuk project kecil?",
    answer: "Karena Quanta menghilangkan kompleksitas spreadsheet yang berantakan, memberikan struktur yang jelas bahkan untuk tim kecil sekalipun."
  },
  {
    question: "Bagaimana Quanta membantu mengelola banyak project sekaligus?",
    answer: "Quanta memiliki dashboard portofolio yang memungkinkan Anda melihat status kesehatan semua project dalam satu tampilan ringkas."
  },
  {
    question: "Bagaimana Quanta membantu mengurangi scope creep?",
    answer: "Fitur Scope Control kami mewajibkan persetujuan formal untuk setiap perubahan scope, sehingga klien atau stakeholder tidak bisa menambah fitur sembarangan."
  },
  {
    question: "Bagaimana Quanta membantu menjaga timeline tetap realistis?",
    answer: "Dengan visualisasi timeline dan milestone yang jelas, Anda bisa melihat dampak keterlambatan satu task terhadap keseluruhan jadwal project."
  },
  {
    question: "Kenapa risk & issue tracking penting, dan bagaimana Quanta membantu?",
    answer: "Menangani risiko sejak dini mencegah kebakaran di kemudian hari. Quanta menyediakan log khusus untuk mencatat dan memitigasi risiko sebelum menjadi masalah besar."
  }
]
const openIndex = ref<number | null>(null)

const toggleFaq = (index: number) => {
  if (openIndex.value === index) {
    openIndex.value = null
  } else {
    openIndex.value = index
  }
}
</script>

<template>
  <section id="faq" class="bg-white py-16 sm:py-24">
    <div class="mx-auto max-w-3xl px-4 sm:px-6 lg:px-8">
      <div class="text-center mb-16">
        <h2 class="text-3xl font-bold tracking-tight text-[color:var(--q-heading)] sm:text-4xl">
          Frequently Asked Questions
        </h2>
      </div>

      <div class="space-y-4">
        <div 
          v-for="(faq, index) in faqs" 
          :key="index" 
          class="rounded-xl border border-[color:var(--q-border)] bg-[color:var(--q-bg)] p-6 overflow-hidden"
        >
          <button 
            @click="toggleFaq(index)"
            class="flex w-full cursor-pointer items-center justify-between gap-1.5 text-left text-[color:var(--q-heading)] focus:outline-none"
            :aria-expanded="openIndex === index"
          >
            <h3 class="text-base font-semibold">
              {{ faq.question }}
            </h3>

            <span class="relative h-5 w-5 shrink-0 ml-4">
              <svg
                class="absolute inset-0 h-5 w-5 transition-opacity duration-300"
                :class="openIndex === index ? 'opacity-0' : 'opacity-100'"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
                stroke-width="2"
              >
                <path stroke-linecap="round" stroke-linejoin="round" d="M12 9v3m0 0v3m0-3h3m-3 0H9m12 0a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg>

              <svg
                class="absolute inset-0 h-5 w-5 transition-opacity duration-300"
                :class="openIndex === index ? 'opacity-100' : 'opacity-0'"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
                stroke-width="2"
              >
                <path stroke-linecap="round" stroke-linejoin="round" d="M15 12H9m12 0a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg>
            </span>
          </button>

          <div 
            class="grid transition-all duration-300 ease-in-out"
            :class="openIndex === index ? 'grid-rows-[1fr] opacity-100 mt-4' : 'grid-rows-[0fr] opacity-0 mt-0'"
          >
            <div class="overflow-hidden">
              <p class="leading-relaxed text-[color:var(--q-text)] text-sm">
                {{ faq.answer }}
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>
