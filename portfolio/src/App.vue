<script setup>
import Timeline from './components/Timeline.vue'
</script>

<template>
  <header>
    <h1>Kobe Berkein</h1>
  </header>

  <!-- placeholder -->
  <div />

  <main>
    <Timeline />
  </main>
</template>

<style scoped>
header {
  line-height: 1.2;
  font-size: 3rem;
  text-align: center;
}

@media (min-width: 1024px) {
  header {
    text-align: right;
    display: flex;
    place-items: center;
    padding-right: calc(var(--section-gap) / 2);
    position: fixed;
    width: calc(50% - var(--section-gap));
    max-width: calc(50vw - var(--section-gap));
    height: 100vh;
  }

  main {
    margin: 47.5vh 0 2rem 0;
  }
}
</style>
