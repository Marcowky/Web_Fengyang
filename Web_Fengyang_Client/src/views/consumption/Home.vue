<template>
    
    
    <div class="box">
        <div style=" position: relative; margin-top:100px; left: 0;  width: 100%; ">
            <Chart>
            </Chart>
        </div>

        <div style=" position: relative; margin: auto; margin-left: 100px;  margin-right: 100px; width: 80%; ">
          <SmallBar>
          </SmallBar>
      </div>       
    </div>


</template>

<script setup>


import Chart from './Chart.vue'
import SmallBar from './SmallBar.vue'


</script>

<style lang="scss" scoped>
.box {
    position: relative;
    min-height: 100%;
}
</style>