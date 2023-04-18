<template>
  <button class="uk-button uk-button-default" type="button" uk-toggle="target: #modal-add-new-instrument">+</button>

  <div id="modal-add-new-instrument" uk-modal>
    <div class="uk-modal-dialog uk-modal-body">
        <h2 class="uk-modal-title">Add a new instrument to the work</h2>
        <select class="uk-select" v-model="instrumentSectorForm">
            <option disabled selected>Select the instrument sector</option>
            <option v-for="instrumentSector in instrumentSectorList" :key="instrumentSector.Id" :value="instrumentSector.Id">{{ instrumentSector.Name }}</option>
        </select>
        <p class="uk-text-right">
            <button class="uk-button uk-button-default uk-modal-close" type="button">Cancel</button>
            <button class="uk-button uk-button-primary" type="button">Save</button>
        </p>
    </div>
</div>
</template>

<script setup lang="ts">
import { IInstrumentSector } from '@/modules/Instruments/model/IInstrumentSector';
import instrumentService from '@/modules/Instruments/service/InstrumentService';
import { ref } from 'vue';

const instrumentSectorForm = ref('');
const instrumentForm = ref('');

const instrumentSectorList = ref<IInstrumentSector[]>([]);
instrumentService.getInstrumentSectorList().then(data => {instrumentSectorList.value = data});
</script>
