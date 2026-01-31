<script setup lang="ts">
import HeroSection from '~/components/landing-page/HeroSection.vue'
import ProblemSection from '~/components/landing-page/ProblemSection.vue'
import SolutionOverviewSection from '~/components/landing-page/SolutionOverviewSection.vue'
import HowItWorksSection from '~/components/landing-page/HowItWorksSection.vue'
import FeaturesSection from '~/components/landing-page/FeaturesSection.vue'
import WhyThisHelpsSection from '~/components/landing-page/WhyThisHelpsSection.vue'
import PricingSection from '~/components/landing-page/PricingSection.vue'
import FAQSection from '~/components/landing-page/FAQSection.vue'
import CTASection from '~/components/landing-page/CTASection.vue'

definePageMeta({
  layout: 'landing'
})
</script>

<template>
  <div>
    <HeroSection />
    <ProblemSection />
    <SolutionOverviewSection />
    <HowItWorksSection />
    <FeaturesSection />
    <WhyThisHelpsSection />
    <PricingSection />
    <FAQSection />
    <CTASection />
  </div>
</template>
