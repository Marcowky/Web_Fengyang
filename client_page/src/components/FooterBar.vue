<template>
    <br><br><br>
  <!--footer-->
  <footer class="footer">
    <div class="footer-bottom">
      <div class="footer-container">
        <ul class="footer-nav">
          <!--景区概况-->
          <ul class="css-ul">
            <li>
            <a class="css-main_title" href="/HomePage" >景区概况</a>
            </li>
            <li>
              <a href="/HomePage">四色丰阳</a>
            </li>
            <li>
              <a href="/HomePage">民族风俗</a>
            </li>
            <li>
              <a href="/HomePage">历史人文</a>
            </li>
          </ul>

          <!--景区资讯-->
          <ul class="css-ul">
            <li>
            <a class="css-main_title" href="/InfoPage" >景区资讯</a>
            </li>
            <li>
              <a href="/InfoPage">景区活动</a>
            </li>
            <li>
              <a href="/InfoPage">景区公告</a>
            </li>
            <li>
              <a href="/InfoPage">志愿活动</a>
            </li>
          </ul>

          <!--景区景点-->
          <ul class="css-ul">
            <li>
            <a class="css-main_title" href="#" >景区景点</a>
            </li>
            <li>
              <a href="/InfoPage">热门景点</a>
            </li>
            <li>
              <a href="/InfoPage">游玩攻略</a>
            </li>
            <li>
              <a href="/InfoPage">路线规划</a>
            </li>
          </ul>

          <!--景区消费-->
          <ul class="css-ul">
            <li>
            <a class="css-main_title" href="#" >景区消费</a>
            </li>
            <li>
              <a href="/ConsumptionPage">酒店民宿</a>
            </li>
            <li>
              <a href="/ConsumptionPage">娱乐项目</a>
            </li>
            <li>
              <a href="/ConsumptionPage">美食特产</a>
            </li>
          </ul>

          <!--投诉建议-->
          <ul class="css-ul">
            <li>
            <a class="css-main_title" href="#" >投诉建议</a>
            </li>
            <li>
              <a href="/ForumPage">景区论坛</a>
            </li>
          </ul>

        </ul>
      </div>
    </div>
  </footer>
</template>

<style scoped>
/* footer的css */
.footer{
  position: relative;
}

.footer-bottom{
  background:#353432;
  width: 100%;
  padding: 0;
}

.footer-container{
  position: static;
  padding: 0px 50px;
  /* margin:10 auto; */
}

.footer-nav{
  display: flex;
  flex-wrap: wrap;
  padding: 40px 0px;
}

.css-ul{
  margin-left: 40px;
}

.css-main_title{
  display: inline-block;
  padding: 10px 0;
  font-size: 16px;
  color: #ffffff;
  border-bottom: 1px solid #A0A0A0;   /* 景区概况底下那条横线 */
}

ul, li{
  list-style: none;   /* 去掉列表前面的小圈圈 */
}

/* .css-li{
  display: block;
  font-size: 14px;
  line-height: 30px;
  color: #aaaaaa;
} */

a{    
  display: block;
  font-size: 10px;
  line-height: 30px;
  color: #aaaaaa;
}
</style>