<template>
  <button class="uk-button uk-button-default" type="button" uk-toggle="target: #modal-add-new-instrument">+</button>

  <div id="modal-add-new-instrument" uk-modal>
    <div class="uk-modal-dialog uk-modal-body">
        <h2 class="uk-modal-title">Add a new instrument to the work</h2>

        <label class="uk-label">Sector</label>
        <select class="uk-select uk-margin-small" v-model="addInstrumentForm.SectorId">
            <option disabled selected>Select the instrument sector</option>
            <option v-for="instrumentSector in instrumentSectorList" :key="instrumentSector.Id" :value="instrumentSector.Id">{{ instrumentSector.Name }}</option>
        </select>

        <label class="uk-label">Type</label>
        <select class="uk-select uk-margin-small" v-model="addInstrumentForm.Type">
            <option disabled selected>Select the instrument type</option>
            <option v-for="instrumentType in instrumentTypeList" :key="instrumentType.Code" :value="instrumentType.Code">{{ instrumentType.Name }}</option>
        </select>

        <label class="uk-label">Instrument</label>
        <select class="uk-select uk-margin-small" v-model="addInstrumentForm.Figi">
            <option disabled selected>Select the instrument</option>
            <option v-for="instrument in instrumentList" :key="instrument.Figi" :value="instrument.Figi">{{ instrument.Name }}</option>
        </select>

        <p class="uk-text-right">
            <button class="uk-button uk-button-default uk-modal-close" type="button" @click="clickCancelButton()">Cancel</button>
            <button class="uk-button uk-button-primary uk-margin-left" type="button" @click="clickSaveButton()">Save</button>
        </p>
    </div>
</div>
</template>

<script setup lang="ts">
import { AddInstrumentFormObject } from '@/modules/Instruments/model/FormObject/AddInstrumentFormObject';
import { IInstrument } from '@/modules/Instruments/model/IInstrument'
import { IInstrumentSector } from '@/modules/Instruments/model/IInstrumentSector.js';
import { IInstrumentType } from '@/modules/Instruments/model/IInstrumentType.js';
import instrumentService from '@/modules/Instruments/service/InstrumentService';
import { reactive, ref, watch } from 'vue';

const instrumentSectorList = ref<IInstrumentSector[]>([]);
const instrumentTypeList = ref<IInstrumentType[]>([]);
const instrumentList = ref<IInstrument[]>([]);

const addInstrumentForm = reactive(new AddInstrumentFormObject());

watch(
  () => addInstrumentForm.SectorId,
  () => {
    if (addInstrumentForm.SectorId) {
      instrumentService.getInstrumentTypeList(addInstrumentForm.SectorId)
        .then(data => {instrumentTypeList.value = data});
    }

    if (addInstrumentForm.SectorId && addInstrumentForm.Type) {
        instrumentService.getInstrumentList(
          addInstrumentForm.SectorId,
          addInstrumentForm.Type
        ).then(data => {instrumentList.value = data});
    }
  }
)

watch(
  () => addInstrumentForm.Type,
  () => {
    if (addInstrumentForm.SectorId && addInstrumentForm.Type) {
        instrumentService.getInstrumentList(
          addInstrumentForm.SectorId,
          addInstrumentForm.Type
        ).then(data => {instrumentList.value = data});
    }
  }
)

const clickCancelButton = () => {
  instrumentTypeList.value = [];
  instrumentList.value = [];

  const tmp = new AddInstrumentFormObject();
  addInstrumentForm.SectorId = tmp.SectorId;
  addInstrumentForm.Type = tmp.Type;
  addInstrumentForm.Figi = tmp.Figi;
};

const clickSaveButton = () => {
  instrumentService.setInstrumentObservable(addInstrumentForm.Figi);
};

instrumentService.getInstrumentSectorList().then(data => {instrumentSectorList.value = data});
</script>
