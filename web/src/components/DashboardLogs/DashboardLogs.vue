<template>
  <DashboardCommonHeader />

  <div class="uk-card uk-card-default uk-card-body">
    <DashboardLogsTable :logList = logList />
  </div>
</template>

<script setup lang="ts">
import {ref, onMounted, onUnmounted} from 'vue';

import DashboardCommonHeader from "@/components/DashboardCommon/DashboardCommonHeader.vue";
import DashboardLogsTable from "./DashboardLogsTable.vue";

import { ILog } from '@/modules/Logs/model/ILog';
import logService from '@/modules/Logs/service/LogService';
import {DATA_UPDATE_INTERVAL_SEC_DEFAULT} from '@/const/CommonConst';

const getLogs = () => logService.getLogList().then(data => {logList.value = data});
const logList = ref<ILog[]>([]);
let intervalId: any;

onMounted(() => {
  getLogs();
  intervalId = setInterval(getLogs, DATA_UPDATE_INTERVAL_SEC_DEFAULT);
});

onUnmounted(() => clearInterval(intervalId));
</script>
