<template>
  <div id="offcanvas-overlay" uk-offcanvas="overlay: true">
    <div class="uk-offcanvas-bar uk-flex uk-flex-column">

      <ul class="uk-nav uk-nav-primary">
          <li>
            <RouterLink to="/">Dashboard</RouterLink>
          </li>
          <li class="uk-nav-divider"></li>
          <li>
            <RouterLink to="/strategies">Strategies</RouterLink>
          </li>
          <li class="uk-nav-divider"></li>
          <li>
            <RouterLink to="/settings">Settings</RouterLink>
          </li>
          <li class="uk-nav-divider"></li>
          <li>
            <RouterLink to="/tasks">Tasks</RouterLink>
          </li>
      </ul>

    </div>
  </div>
</template>

<script setup lang="ts">
import { RouterLink } from "vue-router";
</script>
