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
        <p>Mod: {{ mod }}</p>
      </div>
      <div class="uk-navbar-right">
        <ul class="uk-navbar-nav">
          <li>
            <a uk-navbar-toggle-icon href="#"></a>
            <div class="uk-navbar-dropdown">
              <ul class="uk-nav uk-navbar-dropdown-nav">
                <li>
                  <RouterLink to="/">Dashboard</RouterLink>
                </li>
                <!-- <li>
                  <RouterLink to="/profile">Profile</RouterLink>
                </li> -->
                <li>
                  <RouterLink to="/logs">Logs</RouterLink>
                </li>
              </ul>
            </div>
          </li>
        </ul>
      </div>
    </nav>
  </header>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      mod: "",
    };
  },
  methods: {
    getMod() {
      axios
        .get("/api/getMod")
        .then((response) => this.handleGetModResponse(response.data))
        .catch((error) => console.log(error));
    },
    handleGetModResponse: function (data) {
      if (data.Body.Mod === "production") {
        this.mod = "production";
      } else {
        this.mod = "sandbox";
      }
    },
  },
  mounted() {
    this.getMod();
  },
};
</script>
