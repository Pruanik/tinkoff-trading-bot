<template>
  <nav class="uk-navbar-container uk-margin">
    <div class="uk-container uk-container-expand">
      <div uk-navbar>
        <div class="uk-navbar-left">
          <a
            class="uk-navbar-toggle"
            uk-navbar-toggle-icon
            href="#offcanvas-overlay"
            uk-toggle
          ></a>
        </div>

        <div class="uk-navbar-center">
          <RouterLink class="uk-navbar-item uk-logo" to="/">
            <img src="@/assets/img/logo.svg" width="125" height="125" />
          </RouterLink>
        </div>
      </div>
    </div>
  </nav>
</template>

<script setup lang="ts">
import { RouterLink } from "vue-router";
</script>
