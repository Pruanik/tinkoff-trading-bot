<script setup>
import { RouterLink } from "vue-router";
</script>

<template>
  <header class="uk-background-muted">
    <nav class="uk-navbar-container uk-container" uk-navbar>
      <div class="uk-navbar-left">
        <ul class="uk-navbar-nav">
          <li>
            <RouterLink to="/">
              <img src="@/assets/img/logo.png" width="125" height="125" />
            </RouterLink>
          </li>
        </ul>
      </div>
      <div class="uk-navbar-right">
        <ul class="uk-navbar-nav">
          <li>
            <a uk-navbar-toggle-icon href="#"></a>
            <div class="uk-navbar-dropdown">
              <ul class="uk-nav uk-navbar-dropdown-nav">
                <li>
                  <RouterLink to="/">Дашборд</RouterLink>
                </li>
                <li>
                  <RouterLink to="/profile">Профиль</RouterLink>
                </li>
                <li>
                  <RouterLink to="/logs">Логи</RouterLink>
                </li>
              </ul>
            </div>
          </li>
        </ul>
      </div>
    </nav>
  </header>
</template>
