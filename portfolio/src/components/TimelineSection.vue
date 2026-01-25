<template>
  <div class="item">
    <div class="dot"/>
    <div class="details">
      <h3>
        <slot name="heading"></slot>
      </h3>
      <slot></slot>
    </div>
  </div>
</template>

<style scoped>
.item {
  padding: 0.4rem calc(var(--section-gap) / 2) 1rem calc(var(--section-gap) / 2);
  display: flex;
  position: relative;
}

.details {
  flex: 1;
}

.dot {
  display: flex;
  place-items: center;
  place-content: center;

  color: var(--color-text);
  top: 0.5rem;
  left: -1rem;
  position: absolute;
  border: 0.5rem solid var(--color-text);
  background: var(--color-background);
  border-radius: 100%;
  width: 2rem;
  height: 2rem;
}

h3 {
  font-size: 1.2rem;
  font-weight: 500;
  margin-bottom: 0.4rem;
  color: var(--color-heading);
}

.item:after {
  content: ' ';
  border-left: 0.5rem solid var(--color-text);
  position: absolute;
  left: -0.25rem;
  top: 2.45rem;
  height: 100%;
}

.item:first-of-type:before {
  display: none;
}

.item:last-of-type:after {
  display: none;
}

@media (min-width: 1024px) {
  .item:after{

  }
}
</style>
