<template>
  <div class="infos">
    <p class="temperatureText">{{ temperature + '°' }}</p>
    <div>
      <p class="airportText">{{ airport }}</p>
      <p class="timeText">{{ displayHour + ' - ' + displayDate }}</p>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref, computed } from 'vue'

const temperature = '18'
const airport = 'Nantes'

const date = ref(new Date())
setInterval(() => refreshDate(), 1000)

const displayDate = computed(
  () =>
    days[date.value.getDay()] +
    ' ' +
    date.value.getDate() +
    ' ' +
    months[date.value.getMonth()] +
    ' ' +
    date.value.getFullYear()
)

const displayHour = computed(
  () => date.value.getHours() + ':' + date.value.getMinutes()
)

function refreshDate() {
  date.value = new Date()
}

const months = [
  'Décembre',
  'Janvier',
  'Février',
  'Mars',
  'Avril',
  'Mai',
  'Juin',
  'Juillet',
  'Aout',
  'Septembre',
  'Octobre',
  'Novembre',
]
const days = [
  'Dimanche',
  'Lundi',
  'Mardi',
  'Mercredi',
  'Jeudi',
  'Vendredi',
  'Samedi',
]
</script>

<style scoped>
* {
  @apply font-oxygen;
}
.infos {
  display: flex;
  align-items: center;
  padding-left: 5%;
  padding-bottom: 2%;
}
.temperatureText {
  font-size: 100px;
}
.airportText {
  font-size: 54px;
}
.timeText {
  font-size: 13px;
}
</style>
