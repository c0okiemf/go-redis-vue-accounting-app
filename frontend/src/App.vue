<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router';
</script>

<template>
  <div class="container">
    <header>
      <nav>
        <RouterLink to="/">Dashboard</RouterLink>
        <RouterLink to="/create">Record New Transaction</RouterLink>
        <RouterLink to="/upload">Upload File</RouterLink>
      </nav>
    </header>
    <div class="page-wrapper">
      <RouterView />
    </div>
  </div>
</template>

<style scoped lang="scss">
.container {
  display: flex;
  flex-direction: row;
  width: 100%;
  height: 100%;
  flex-wrap: nowrap;
  overflow: hidden;

  @media (max-width: $desktop-breakpoint) {
    flex-direction: column;
  }
}

.page-wrapper {
  width: 100%;
  overflow-x: scroll;
}

header {
  line-height: 1.5;
  flex: 0 0 auto;
  display: flex;
  place-items: center;
  padding-right: calc(var(--section-gap) / 4);
  place-items: center;
  justify-content: center;

  @media (max-width: $desktop-breakpoint) {
    padding: 0.5rem 0;
    position: fixed;
    width: 100%;
    background: var(--color-background);
  }
}

nav {
  width: 100%;
  font-size: 12px;
  text-align: center;
  display: flex;
  flex-direction: column;
  border-right: 1px solid var(--color-border);
  gap: 1rem;
  text-align: left;
  margin-left: -1rem;
  font-size: 1rem;
  padding: 1rem 0;
  margin-top: 1rem;

  @media (max-width: $desktop-breakpoint) {
    border-right: none;
    flex-direction: row;
    justify-content: center;
    margin: 0;
    padding: 0.5rem 0;
  }
}

nav a.router-link-exact-active {
  color: var(--text-highlighted);
}

nav a.router-link-exact-active:hover {
  background-color: transparent;
}

nav a {
  display: inline-block;
  padding: 0 1rem;
  color: var(--color-text);
}

nav a:first-of-type {
  border: 0;
}
</style>
