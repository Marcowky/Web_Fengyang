<template>

  <!--footer-->
  <div class="space"></div>
  <footer class="footer">
    <div class="footer-bottom">
      <div class="footer-container">
        <ul class="footer-nav">
          <template v-for="item in menuItems">
              <ul class="css-ul" v-if="item.mainMenu && !item.index.includes('-')">
                  <li>
                    <text class="css-main_title" >{{ item.label }}</text>
                  </li>
                
                  <template v-for="subItem in menuItems">
                    <li v-if="!subItem.mainMenu && subItem.index.startsWith(item.index + '?')">
                      <el-link :href="link+subItem.index">{{ subItem.label }}</el-link>
                      <!-- <el-link type="primary">primary</el-link> -->
                    </li>
                  </template> 
              </ul>
          </template>
        </ul>
      </div>
    </div>
  </footer>
</template>

<script lang="ts" setup>
import { onMounted, ref } from "vue";
// 导入菜单选项配置文件
import config from '../config/config.json';

interface MenuItem {
    index: string;
    label: string;
    mainMenu: boolean;
    hasSub: boolean;  
};

const link = "http://localhost:5173/#"
const menuItems = ref(config.menuItems as MenuItem[]);

</script>

<style scoped>
/* footer的css */
.footer {
  position: absolute;
  width: 100%;
  bottom: 0px;

}

.space {
  height: 300px;
}

.footer-bottom {
  background: #353432;
  width: 100%;
  margin-top: 50px;
}

.footer-container {
  position: static;
  /* padding: 10px 90px; */
  /* margin:10 auto; */

}

.footer-nav {
  display: flex;
  flex-wrap: wrap;
  padding: 40px 0px;
}

.css-ul {
  margin-left: 0px;
}

.css-main_title {
  display: inline-block;
  padding: 10px 0;
  font-size: 16px;
  color: #ffffff;
  border-bottom: 1px solid #A0A0A0;
  /* 景区概况底下那条横线 */
}

ul,
li {

  list-style: none;
  /* 去掉列表前面的小圈圈 */
  justify-content: space-evenly;
  /* 设置flex中间距一样 */

}

/* .css-li{
  display: block;
  font-size: 14px;
  line-height: 30px;
  color: #aaaaaa;
} */

a {
  text-decoration: none;
  display: block;
  font-size: 13px;
  line-height: 30px;
  color: #aaaaaa;
}

/* 选择鼠标移到a标签链接上的样式： */
@media (hover: hover) {
  a:hover {
    background-color: hsla(160, 100%, 37%, 0.2);
  }
}</style>