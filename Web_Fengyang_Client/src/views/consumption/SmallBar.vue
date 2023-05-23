<template>
    <el-tabs v-model="activeName" class="demo-tabs" @tab-click="handleClick">
      <el-tab-pane label="特产" name="first"></el-tab-pane>
      <el-tab-pane label="美食" name="second"></el-tab-pane>
      <el-tab-pane label="住宿" name="third"></el-tab-pane>
      <el-tab-pane label="娱乐" name="fourth"></el-tab-pane>
    </el-tabs>

    <template v-if="showFirst">
      <el-row>
        <el-col
          v-for="(o, index) in 6"
          :key="o"
          :span="6"
          :offset="index % 3 ? 2 : 0"
        >
          <el-card class="card" :body-style="{ padding: '0px' }">
            <img :src="`../src/assets/pic${index + 7}.jpg`" class="image"/>
            <div style="padding: 5px; font-size: 13px; color: #999;">
              <span>开放时间8:00-18:00</span>
            </div>
          </el-card>
        </el-col>
      </el-row>
    </template>
  
    <template v-if="showSecond">
      <div>
      <el-row>
        <el-col
          v-for="(o, index) in 6"
          :key="o"
          :span="6"
          :offset="index % 3 ? 2 : 0"
        >
          <el-card class="card" :body-style="{ padding: '0px' }">
            <img src="../assets/pic7.jpg" class="image"/>
            <div style="padding: 5px; font-size: 13px; color: #999;">
              <span>开放时间8:00-18:00</span>
            </div>
          </el-card>
        </el-col>
      </el-row>
      </div>
    </template>

</template>

  <script lang="ts" setup>
  import { ref } from 'vue'
  import type { TabsPaneContext } from 'element-plus'
  
  const activeName = ref('first')
  const showFirst  = ref(true)
  const showSecond = ref(false)
  
  const handleClick = (tab: TabsPaneContext, event: Event) => {
    console.log(tab, event)
    
    if (tab.paneName === 'first') {
      showFirst.value = true
    } else {
      showFirst.value = false
    }
  
    if (tab.paneName === 'second') {
      showSecond.value = true
    } else {
      showSecond.value = false
    }
  
  }
  </script>
  
  <style>
  .demo-tabs > .el-tabs__content {
    padding: 12px;
    color: #6b778c;
  }
  
  .card {
    width: 280px;
    height: 200px;
  }
  
  .image {
    width: 280px;
    height: 160px;
    object-fit: cover;
  }
  </style>
  