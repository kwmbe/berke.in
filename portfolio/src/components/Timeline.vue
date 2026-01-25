<script setup>
import TimelineSection from './TimelineSection.vue'
</script>

<template>
  <TimelineSection>
    <template #heading>Today</template>
  </TimelineSection>

  <TimelineSection>
    <template #heading>Junior Software Developer</template>

    <ul>
      <li icon="󰁥"><a href="https://easypost.eu/">EasyPost</a></li>
      <li icon="󰃭">July 2023 - Now</li>
    </ul>
  </TimelineSection>

  <TimelineSection>
    <template #heading>Intern Frontend Developer</template>

    <ul>
      <li icon="󰁥"><a href="https://www.deliverect.com/en-us/">Deliverect</a></li>
      <li icon="󰃭">September 2022 - January 2023</li>
    </ul>
  </TimelineSection>

  <TimelineSection>
    <template #heading>Finalist CSCBE Hackaton</template>

    <ul>
      <li icon="󰁥"><a href="https://www.cybersecuritychallenge.be/">Cyber Security Challenge Belgium</a></li>
      <li icon="󰃭">March 2022</li>
    </ul>
  </TimelineSection>

  <TimelineSection>
    <template #heading>Bachelor Electronics and ICT</template>

    <ul>
      <li icon="󰁥"><a href="https://odisee.be/">Odisee</a></li>
      <li icon="󰃭">September 2019 - January 2023</li>
    </ul>
  </TimelineSection>

  <TimelineSection>
    <template #heading>Born</template>

    <ul>
      <li icon="󰃭">17 May 2001</li>
    </ul>
  </TimelineSection>
</template>

<style scoped>
[icon]:before {
  content: attr(icon);
  padding-right: .5em;
  font-size: 1.2em;
}

ul {
  list-style-type: none;
  padding: 0;
}
</style>
