<script setup lang="ts">
import { computed } from 'vue';

const MAX_PAGES = 5;

const props = defineProps<{
  total: number;
  itemsPerPage: number;
  currentPage: number;
}>();

const emit = defineEmits<{
  (event: 'onSetPage', payload: number): void;
}>();

const displayPages = computed(() => {
  const pages = [];
  const start = Math.max(1, props.currentPage - Math.floor(MAX_PAGES / 2));
  const end = start + MAX_PAGES - 1;

  for (let i = start; i <= end; i++) {
    pages.push(i);
  }

  return pages;
});

const totalPages = computed(() => Math.ceil(props.total / props.itemsPerPage));

const onPrevPage = () => {
  if (props.currentPage > 1) {
    emit('onSetPage', props.currentPage - 1);
  }
};

const onNextPage = () => {
  if (props.currentPage * props.itemsPerPage < props.total) {
    emit('onSetPage', props.currentPage + 1);
  }
};

const onFirstPage = () => {
  emit('onSetPage', 1);
};

const onLastPage = () => {
  emit('onSetPage', Math.ceil(props.total / props.itemsPerPage));
};
</script>

<template>
  <div class="pagination">
    <button @click="onFirstPage" :disabled="currentPage === 1" class="cell arrow">
      <v-icon name="bi-chevron-bar-up" />
    </button>
    <button @click="onPrevPage" :disabled="currentPage === 1" class="cell arrow">
      <v-icon name="bi-arrow-up" />
    </button>

    <div v-for="page in displayPages" :key="page">
      <button
        v-if="page <= totalPages"
        @click="$emit('onSetPage', page)"
        :class="['cell', page === currentPage && 'selected']"
      >
        <span>{{ page }}</span>
      </button>
      <div v-else class="cell"></div>
    </div>

    <button @click="onNextPage" :disabled="currentPage * itemsPerPage >= total" class="cell arrow">
      <v-icon name="bi-arrow-down" />
    </button>
    <button @click="onLastPage" :disabled="currentPage * itemsPerPage >= total" class="cell arrow">
      <v-icon name="bi-chevron-bar-down" />
    </button>
  </div>
</template>

<style scoped lang="scss">
.pagination {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  gap: 1rem;
  margin-top: 1rem;
  padding-right: calc(var(--section-gap) / 4);

  @media (max-width: $desktop-breakpoint) {
    flex-direction: row;
    gap: 0.5rem;
    margin: 0 auto;
    padding: 0;
  }
}

button {
  background: none;
  border: none;
  cursor: pointer;
  color: unset;

  svg {
    width: 100%;
    height: 100%;
    fill: var(--icon-active);
  }

  &:disabled {
    svg {
      fill: var(--icon-disabled);
    }
  }
}

.cell {
  width: 50px;
  height: 50px;

  @media (max-width: $desktop-breakpoint) {
    width: 30px;
    height: 30px;
  }
}

.arrow {
  @media (max-width: $desktop-breakpoint) {
    transform: rotate(270deg);
  }
}

.selected {
  background: var(--background-highlighted);
  color: var(--background-highlighted-text);
}
</style>
