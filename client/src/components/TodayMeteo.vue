<template>
  <div class="infos">
    <p class="temperatureText">
      {{ temperature.slice(0, -1)
      }}<span class="coloredNumber"
        >{{ temperature[temperature.length - 1] }} </span
      >°
    </p>
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
  () =>
    date.value.getHours() +
    ':' +
    (date.value.getMinutes() > 10 ? '' : '0') +
    '' +
    date.value.getMinutes()
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
/* * {
  @apply font-oxygen;
} */
.coloredNumber {
  color: #24aad3;
}
.infos {
  background-color: #ffffff19;
  -webkit-backdrop-filter: blur(15px);
  backdrop-filter: blur(15px);
  border-radius: 10px;
  display: flex;
  align-items: center;
  margin-left: 60px;
  margin-bottom: 54px;
  padding: 22px 40px 30px 40px;
}
.temperatureText {
  font-weight: 200;
  font-size: 100px;
  line-height: 80%;
}
.airportText {
  font-size: 45px;
  font-weight: 300;
}
.timeText {
  font-size: 13px;
}
</style>
