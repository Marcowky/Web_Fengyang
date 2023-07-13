<template>
  <!--footer-->
  <div class="space"></div>
  <footer class="footer">
    <div class="footer-bottom">
      <div class="footer-container">
        <div class="weather_wrap">
          <el-row>
            <el-col v-for="(item, index) in weatherData" :key="index" :span="6">
              <div class="day_weather_wrap">
                <div class="weather_component">
                  <div class="weather_icon">
                    <el-icon :size="60" color="white">
                      <component :is="getWeatherIcon(item.dayWeather)" />
                    </el-icon>
                  </div>
                  <div class="weather_detail">
                    <li>{{ formatMonthDay(item.date, item.week) }}</li>
                    <li>{{ item.nightTemp }}°C~{{ item.dayTemp }}°C</li>
                    <li>{{ item.dayWeather }}</li>
                  </div>
                </div>

              </div>
            </el-col>
          </el-row>
        </div>
        <el-divider class="divider_style" />
        <ul class="footer-nav">
          <template v-for="item in menuItems">
            <div v-if="!item.mainMenu.includes('/')">
              <li>
                <text class="css-main_title">{{ item.label }}</text>
              </li>

              <template v-for="subItem in menuItems">
                <li v-if="subItem.mainMenu == item.index && subItem.mainMenu != '/hotel'">
                  <el-link :href="link + item.index + '?category=' + subItem.index">{{ subItem.label }}</el-link>
                  <!-- <el-link type="primary">primary</el-link> -->
                </li>
                <li v-if="subItem.mainMenu == item.index && subItem.mainMenu == '/hotel'">
                  <el-link :href="subItem.index">{{ subItem.label }}</el-link>
                  <!-- <el-link type="primary">primary</el-link> -->
                </li>
              </template>
            </div>
          </template>
        </ul>
      </div>
    </div>
  </footer>
</template>

<script setup>
// 引入高德地图api,准备实现天气显示
import AMapLoader from '@amap/amap-jsapi-loader'
import { onMounted, ref } from "vue"
import { menuList } from '../api/category';
// element-ui 的图标引入图标
import { Sunny, Cloudy, Pouring, Lightning, Drizzling, PartlyCloudy, MostlyCloudy } from '@element-plus/icons-vue'
import '@element-plus/icons-vue'
// 安全秘钥
window._AMapSecurityConfig = {
  securityJsCode: '3626fd4fb9d0809a04788c6df4d45953'
}

const link = window.location.origin
const menuItems = ref(menuList())
const weatherData = ref(null)

const getWeatherIcon = (dayWeather) => {
  const sunnyWeather = ['热', '晴', '平静'];
  const partlyCloudyWeather = ['少云', '晴间多云'];
  const cloudyWeather = ['阴', '霾', '中度霾', '重度霾', '严重霾', '雾', '浓雾', '强浓雾', '清雾', '大雾', '特强浓雾'];
  const mostlyCloudyWeather = ['多云'];
  const drizzlingWeather = ['阵雨', '小雨', '中雨', '小雨-中雨'];
  const pouringWeather = ['大雨', '暴雨', '大暴雨', '特大暴雨', '极端降雨', '中雨-大雨', '大雨-暴雨', '暴雨-特大暴雨', '雨雪天气', '雨夹雪', '阵雨夹雪', '冻雨'];
  const lightningWeather = ['雷阵雨', '雷阵雨并伴有冰雹'];

  if (sunnyWeather.includes(dayWeather)) {
    return Sunny
  } else if (partlyCloudyWeather.includes(dayWeather)) {
    return PartlyCloudy
  } else if (cloudyWeather.includes(dayWeather)) {
    return Cloudy
  } else if (mostlyCloudyWeather.includes(dayWeather)) {
    return MostlyCloudy
  } else if (drizzlingWeather.includes(dayWeather)) {
    return Drizzling
  } else if (pouringWeather.includes(dayWeather)) {
    return Pouring
  } else if (lightningWeather.includes(dayWeather)) {
    return Lightning
  } else {
    return MostlyCloudy
  }
}

const formatMonthDay = (dateString, week) => {
  const [year, month, day] = dateString.split('-')
  const formattedMonth = parseInt(month, 10).toString() // 去掉月份前导零
  const formattedDay = parseInt(day, 10).toString() // 去掉日期前导零
  let week_day
  switch (week) {
    case '1':
      week_day = '周一'
      break
    case '2':
      week_day = '周二'
      break
    case '3':
      week_day = '周三'
      break
    case '4':
      week_day = '周四'
      break
    case '5':
      week_day = '周五'
      break
    case '6':
      week_day = '周六'
      break
    case '7':
      week_day = '周日'
      break
    default:
      week_day = ''
      break
  }
  return `${formattedMonth}-${formattedDay}(${week_day})`
}
const initMap = () => {
  AMapLoader.load({
    key: 'e0eb946f632c6473f40386d512e056d6', // 申请好的Web端开发者Key，首次调用 load 时必填
    version: '2.0', // 指定要加载的 JSAPI 的版本，缺省时默认为 1.4.15
    plugins: ['AMap.Weather'] // 需要使用的插件列表，
  }).then((AMap) => {
    var weather = new AMap.Weather()
    weather.getForecast('连州市', function (err, data) {
      weatherData.value = data.forecasts
      // console.log(weatherData.value)
      // console.log(weatherData.value[0])
    })
  })
}

onMounted(() => {
  initMap()
})

</script>

<style scoped>
/* footer的css */
.footer {
  position: absolute;
  width: 100%;
  bottom: 0;
}

.weather_wrap {
  top: 20px;
  height: 80px;
  display: flex;
  justify-content: center;
  /* padding-left: 20px; */

  .day_weather_wrap {
    display: flex;
    align-items: center;

    .weather_component {
      display: flex;
      align-items: center;
      padding-left: 90px;
      padding-right: 90px;

      .weather_icon {
        top: 5px;
        align-items: center;
        padding-right: 10px;
      }

      .weather_detail {
        color: white;
        left: 4px;
      }
    }
  }

}

.space {
  height: 500px;
}

.footer-bottom {
  background: #353432;
  width: 100%;
  margin-top: 100px;
}

.footer-container {
  position: static;
  padding: 10px 90px;
}

.divider_style {
  left: 11%;
  width: 79%;
  top: 20px;
}

.footer-nav {
  display: flex;
  flex-wrap: wrap;
  padding: 40px 0px;
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
}
</style>