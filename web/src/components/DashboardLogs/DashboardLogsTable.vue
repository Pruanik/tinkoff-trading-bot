<template>
  <div class="uk-overflow-auto">
    <table class="uk-table uk-table-divider">
      <thead>
          <tr>
          <th>ID</th>
          <th>Category</th>
          <th>Level</th>
          <th>Message</th>
          <th>Context</th>
          <th>Created at</th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="log in logList"
          :key="log.Id"
          :class="'log-' + log.Level + '-color'"
        >
          <td>{{ log.Id }}</td>
          <td>{{ log.Category }}</td>
          <td>{{ log.Level }}</td>
          <td>{{ log.Message }}</td>
          <td>{{ log.Context }}</td>
          <td>{{ log.CreatedAt }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script setup lang="ts">
import { ILog } from '@/modules/Logs/model/ILog';

interface Props {
  logList: ILog[]
}

const props = defineProps<Props>()
</script>

<style lang="scss">
.log-info-color {
  background-color: #ccefeb;
}
.log-warning-color {
  background-color: #efebcc;
}
.log-debug-color {
  background-color: #efefcc;
}
.log-error-color {
  background-color: #efd8cc;
}
.log-fatal-color {
  background-color: #ceccef;
}
.log-panic-color {
  background-color: #a8a7a1;
}
</style>
