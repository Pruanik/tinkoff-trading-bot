<template>
  <router-link
    v-bind="$props"
    custom
    v-slot="{ isActive, href, navigate }">
      <li :class="{[activeClass]: isActive}">
        <a
          v-bind="$attrs"
          :href="href"
          @click="navigate">
          <slot />
        </a>
      </li>
  </router-link>
</template>

<script lang="ts">
import {RouterLink} from 'vue-router';

export default {
  name: 'DashboardCommonMenuLink',
  inheritAttrs: false,
  props:{ ...RouterLink.props }
}
</script>
