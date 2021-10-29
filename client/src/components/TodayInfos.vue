<template>
  <!-- TODO : Ici c'est pour toi loïs -->
  <div class="rectangle">
    <div class="today">
      <p>Aujourd'hui</p>
    </div>
    <div class="content">
      <div class="carte">
        <p class="texteCarte temperature">{{ todayData?.pressure }}</p>
        <p class="indicateur">TEMPERATURE</p>
      </div>
      <div class="carte">
        <p class="texteCarte pression">{{ todayData?.temp }}</p>
        <p class="indicateur">PRESSION</p>
      </div>
      <div class="carte">
        <p class="texteCarte vent">{{ todayData?.wind }}</p>
        <p class="indicateur">VENT</p>
      </div>
    </div>
  </div>
</template>
<script lang="ts">
import { PropType, defineComponent } from 'vue'
export default defineComponent({
  props: {
    todayData: {
      type: Object as PropType<Measure>,
    },
  },
})
</script>
<style lang="scss" scoped>
$darkBlue: #3c91ad;

.rectangle {
  background-color: $darkBlue;
  margin-left: 55px;
  margin-right: 55px;
  border-radius: 10px;
  padding-top: 1px;
  padding-left: 10px;
  padding-right: 10px;
}

.today {
  text-align: center;
  p {
    font-size: 22px;
    color: #f9fafa;
    margin-top: 11px;
    margin-bottom: 14px;
  }
}

.content {
  display: flex;
  justify-content: space-around;

  .carte {
    background-color: #24aad3;
    border-radius: 20px;
    width: 110px;
    height: 110px;
    margin-bottom: 20px;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
  }

  .indicateur {
    font-size: 12px;
    color: #ffffff;
    opacity: 0.7;
    line-height: 90%;
    padding-bottom: 15px;
  }

  .temperature::after {
    content: '°C';
    font-weight: normal;
    font-size: 18px;
    position: relative;
    bottom: 27px;
  }

  .pression::after {
    content: 'Pa';
    font-weight: normal;
    font-size: 18px;
    position: relative;
    bottom: 27px;
  }

  .vent::after {
    content: 'km/h';
    font-weight: normal;
    font-size: 9px;
    position: relative;
    bottom: 34px;
  }
}

.texteCarte {
  font-weight: 500;
  font-size: 45px;
  color: #f9fafa;
  margin-left: 17px;
}
</style>
