<template>
  <div>
    <div class="uk-flex">
      <v-select
        :options="options"
        label="country"
        index="code"
        style="width: 120px"
      />
      <button class="uk-button uk-button-default">+</button>
    </div>

    <ul class="uk-tab-left" uk-tab="animation: uk-animation-fade">
      <li><a href="#">Apple</a></li>
      <li><a href="#">IBM</a></li>
      <li><a href="#">Yandex</a></li>
    </ul>
  </div>
</template>

<script>
import vSelect from "vue-select";

export default {
  name: "InstrumentsSidebar",
  components: {
    "v-select": vSelect,
  },
  data() {
    return {
      options: ["Apple", "IBM", "Vk", "Yandex"],
    };
  },
  methods: {
    showModal() {
      this.isModalVisible = true;
    },
    closeModal() {
      this.isModalVisible = false;
    },
  },
};
</script>
