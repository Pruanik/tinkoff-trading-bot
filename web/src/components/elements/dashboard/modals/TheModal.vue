<template>
  <div>MOODAAAL<button class="uk-modal-close-default" type="button" uk-close @click="close"></button></div>
</template>

<script>
export default {
  name: "TheModal",
  methods: {
    close() {
      this.$emit("close");
    },
  },
};
</script>
