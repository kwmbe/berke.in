<script setup>
import TimelineSection from './TimelineSection.vue'
</script>

<template>
  <TimelineSection>
    <template #heading>Today</template>
  </TimelineSection>

  <TimelineSection>
    <template #heading>Junior Software Developer</template>

    <a href="https://easypost.eu/">EasyPost</a>
    <p>July 2023 - Now</p>
  </TimelineSection>

  <TimelineSection>
    <template #heading>Intern Frontend Developer</template>

    <a href="https://www.deliverect.com/en-us/">Deliverect</a>
    <p>September 2022 - January 2023</p>
  </TimelineSection>

  <TimelineSection>
    <template #heading>Finalist CSCBE Hackaton</template>

    <a href="https://www.cybersecuritychallenge.be/">Cyber Security Challenge Belgium</a>
    <p>March 2022</p>
  </TimelineSection>

  <TimelineSection>
    <template #heading>Bachelor Electronics and ICT</template>

    <a href="https://odisee.be/">Odisee</a>
    <p>September 2019 - January 2023</p>
  </TimelineSection>

  <TimelineSection>
    <template #heading>Born</template>

    <p>17 May 2001</p>
  </TimelineSection>
</template>
