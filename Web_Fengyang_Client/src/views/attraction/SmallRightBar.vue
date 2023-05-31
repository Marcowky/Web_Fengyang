<template>
    <el-tabs v-model="activeName" :tab-position="tabPosition" style="height: 500px" class="demo-tabs" @tab-click="handleClick">
      <el-tab-pane label="热门景点" name="first">
        <template v-if="showFirst">
            <el-row>
              
              <el-col :span="6">
                <a href="http://www.baidu.com">
                  <el-card class="card" :body-style="{ padding: '0px' }" shadow="hover">
                    <img src="../../assets/pic7.jpg" class="image"/>
                    <div class="text-wrapper">
                    <div style="font-size: 15px; color: #000;">
                      <span>阳光牧场</span>
                    </div>
                    <div style="font-size: 13px; color: #999;">
                      <span>开放时间8:00-18:00</span>
                    </div>
                  </div>
                  </el-card>
                </a>
              </el-col>
    
              <el-col :span="6" :offset="2">
                <a href="http://www.baidu.com">
                  <el-card class="card" :body-style="{ padding: '0px' }" shadow="hover">
                    <img src="../../assets/pic8.jpg" class="image"/>
                    <div class="text-wrapper">
                    <div style="padding: 0px; font-size: 15px; color: #000;">
                      <span>青山绿水</span>
                    </div>
                    <div style="padding: 0px; font-size: 13px; color: #999;">
                      <span>开放时间8:00-18:00</span>
                    </div>
                  </div>
                  </el-card>
                </a>
              </el-col>
    
              <el-col :span="6" :offset="2">
                <a href="http://www.baidu.com">
                  <el-card class="card" :body-style="{ padding: '0px' }" shadow="hover">
                    <img src="../../assets/pic9.jpg" class="image"/>
                    <div class="text-wrapper">
                      <div style="padding: 0px; font-size: 15px; color: #000;">
                        <span>万里晴空</span>
                      </div>
                      <div style="padding: 0px; font-size: 13px; color: #999;">
                        <span>开放时间8:00-18:00</span>
                      </div>
                    </div>
                    </el-card>
                  </a>
                </el-col>
              </el-row>
      
            <el-row class="card-wrapper">
              <el-col :span="6" >
                <a href="http://www.baidu.com">
                <el-card class="card" :body-style="{ padding: '0px' }" shadow="hover">
                  <img src="../../assets/pic10.jpg" class="image"/>
                  <div class="text-wrapper">
                  <div style="padding: 0px; font-size: 15px; color: #000;">
                    <span>阡陌交通</span>
                  </div>
                  <div style="padding: 0px; font-size: 13px; color: #999;">
                    <span>开放时间8:00-18:00</span>
                  </div>
                </div>
                </el-card>
                </a>
              </el-col>
    
              <el-col :span="6" :offset="2">
                <a href="http://www.baidu.com">
                <el-card class="card" :body-style="{ padding: '0px' }" shadow="hover">
                  <img src="../../assets/pic11.jpg" class="image"/>
                  <div class="text-wrapper">
                  <div style="padding: 0px; font-size: 15px; color: #000;">
                    <span>夕阳乡巷</span>
                  </div>
                  <div style="padding: 0px; font-size: 13px; color: #999;">
                    <span>开放时间8:00-18:00</span>
                  </div>
                </div>
                </el-card>
                </a>
              </el-col>
    
              <el-col :span="6" :offset="2">
                <a href="http://www.baidu.com">
                <el-card class="card" :body-style="{ padding: '0px' }" shadow="hover">
                  <img src="../../assets/pic12.jpg" class="image"/>
                  <div class="text-wrapper">
                  <div style="font-size: 15px; color: #000;">
                    <span>落日余晖</span>
                  </div>
                  <div style="font-size: 13px; color: #999;">
                    <span>开放时间8:00-18:00</span>
                  </div>
                  </div>
                </el-card>
                </a>
              </el-col>
    
            </el-row>
    
          </template>
      </el-tab-pane>

      <el-tab-pane label="游玩攻略" name="second"></el-tab-pane>

      <el-tab-pane label="路线规划" name="third"></el-tab-pane>

      <el-tab-pane label="未完待续" name="fourth"></el-tab-pane>

    </el-tabs>
</template>

  <script lang="ts" setup>
  import { ref } from 'vue'
  import type { TabsPaneContext } from 'element-plus'

  const tabPosition = ref('left')
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
    padding: 32px;
    color: #6b778c;
    font-size: 32px;
    font-weight: 600;
  }
  
  .el-tabs--right .el-tabs__content,
  .el-tabs--left .el-tabs__content {
    height: 100%;
  }

  .card-wrapper {
    margin-top: 20px;
  }
  
  .text-wrapper {
    display: flex;
    flex-direction: row;
    justify-content: center;
    align-items: flex-end;
  }
  
  .card {
    width: 280px;
    height: 220px;
    border-radius: 15px;
  }
  
  .card:hover{
    transform: scale(1.05);
  }
  
  .image {
    width: 280px;
    height: 180px;
    object-fit: cover;
  }
  
  a{
    text-decoration:none;
  }
  </style>
  