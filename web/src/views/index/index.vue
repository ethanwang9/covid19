<template>
  <el-row justify="center">
    <el-col :span="22">
      <!--菜单-->
      <Menu></Menu>

      <!--内容1-->
      <el-row class="content1">
        <el-col class="content1-text">
          <!--文字-->
          <h1>COVID-19</h1>
          <h1 style="margin-top: -50px;">大数据<span class="content1-text-c1">可视化</span>系统</h1>
          <p style="margin-top: 50px">可视化网络爬虫测绘网络空间疫情数据</p>
          <p style="margin-top: -20px">依托网络空间大数据分析、处理、图表化展示，让你比病毒更了解病毒</p>
        </el-col>
        <!--按钮-->
        <el-col class="content1-btn">
          <el-button class="content1-btn-go" color="#ffffff" size="large" @click="gotoPage('/login')">
            <span>立刻体验</span>
            <i-ep-right></i-ep-right>
          </el-button>
        </el-col>
        <!--图标-->
        <el-col class="content1-down">
          <i-ep-ArrowDownBold class=""></i-ep-ArrowDownBold>
        </el-col>
      </el-row>

      <!--内容2-->
      <el-row class="content2">
        <el-col class="content2-text">
          <!--标题-->
          <h1>快捷准确的数据呈现方式</h1>
          <!--介绍-->
          <p>
            数据来源采用阿里健康、腾讯健康、丁香园、国家卫健委等权威机构。<br>
            让数据动起来，将信息转化为圆圆圈圈的图表，人人都说So Easy ！
          </p>
        </el-col>
        <!--图表-->
        <el-col class="content2-chars">
          <el-col :md="10" :sm="18" :xs="24" class="content2-chars-item1">
            <img alt="echars1" src="@/assets/images/index/index/e-chars-1.png">
          </el-col>
          <el-col :md="10" :sm="18" :xs="24" class="content2-chars-item2">
            <img alt="echars2" src="@/assets/images/index/index/e-chars-2.png">
          </el-col>
        </el-col>
      </el-row>

      <!--内容3-->
      <el-row class="content3">
        <el-col class="content3-text">
          <!--标题-->
          <h1>主要功能</h1>
        </el-col>
        <!--图表-->
        <el-row class="content3-box">
          <el-col :md="14" :sm="24" class="content3-box-item1">
            <h2>确诊信息</h2>
            <p>根据用户归属地展示当前区域确诊等相关信息，使用图标、趋势线等表现手法让用户了解当前区域疫情情况</p>
          </el-col>
          <el-col :md="9" :sm="24" class="content3-box-item2">
            <h2>热点动态</h2>
            <p>根据用户归属地展示用户所在地区疫情动态信息</p>
          </el-col>
          <el-col :md="9" :sm="24" class="content3-box-item3">
            <h2>出行政策</h2>
            <p>根据用户出发地和目的地查询疫情防控出行条例，展示相关工作人员联系信息</p>
          </el-col>
          <el-col :md="9" :sm="24" class="content3-box-item4">
            <h2>风险地区</h2>
            <p>根据国家卫健委全国风险地区数据，展示用户归属地风险地区相关信息</p>
          </el-col>
          <el-col :md="5" :sm="24" class="content3-box-item5">
            <h2>未来</h2>
            <p>疫情已结束，人人摘下了口罩，出行交通已不再限制...</p>
          </el-col>
        </el-row>
      </el-row>

      <!--内容4-->
      <el-row class="content4">
        <el-col class="content4-text">
          <h1>开启数据可视化</h1>
          <p>让图表量化数据，让信息传播更高效</p>
          <el-button class="content4-btn" color="#ea0000" size="large" @click="gotoPage('/login')">
            <span>立刻体验</span>
            <i-ep-right></i-ep-right>
          </el-button>
        </el-col>
      </el-row>

      <!--版权-->
      <Footer></Footer>

      <!--End-->
    </el-col>
  </el-row>

  <!--  返回顶部-->
  <el-backtop :bottom="100">
    <div class="TOTOP">UP</div>
  </el-backtop>
</template>

<script lang="ts" setup>
import Menu from '@/components/index/Menu.vue'
import Footer from '@/components/index/Footer.vue'
import {useRouter} from "vue-router";

// router
const router = useRouter()

// 跳转页面
function gotoPage(path: string) {
  router.push(path)
}
</script>

<style lang="scss" scoped>
//内容1
.content1 {
  margin-top: 20px;
  text-align: center;

  .content1-text {

    h1 {
      font-size: 70px;
      font-weight: bold;
      letter-spacing: 20px;
      line-height: 100px;
    }

    p {
      font-size: 30px;
      line-height: 50px;
    }

    &-c1 {
      color: #25c6da;
      background-image: linear-gradient(to right, #21A3AF, #21B568);
      -webkit-background-clip: text;
      -webkit-text-fill-color: transparent;
    }
  }

  .content1-btn {
    margin-top: 30px;

    &-go {
      font-size: 30px;
      width: 280px;
      height: 70px;
      color: #0256FF;
      border-radius: 8px;
      letter-spacing: 8px;

      span {
        padding-right: 10px;
      }
    }
  }

  &-down {
    font-size: 80px;
    margin-top: 120px;
    animation: bounce 2s 2s infinite;
  }
}

//内容2
.content2 {
  margin-top: 150px;

  &-text {
    h1 {
      font-size: 50px;
      font-weight: bold;
      letter-spacing: 6px;
    }

    p {
      font-size: 30px;
      font-weight: 350;
      letter-spacing: 6px;
      line-height: 45px;
      font-weight: lighter;
    }
  }

  &-chars {
    display: flex;
    justify-content: space-around;
    align-items: center;
    flex-wrap: wrap;
    margin-top: 40px;

    &-item1, &-item2 {
      width: 100%;
      height: auto;
      border-radius: 15px;
      margin: 15px 0;
      display: flex;
      justify-content: center;
      align-items: center;
    }

    &-item1 {
      background-color: #4199FF;
    }

    &-item2 {
      background-color: #5C47FF;
    }

    img {
      width: 90%;
      height: auto;
    }
  }
}

//内容3
.content3 {
  margin-top: 50px;

  &-text {
    h1 {
      font-size: 50px;
      font-weight: bold;
      letter-spacing: 6px;
    }
  }

  &-box {
    display: flex;
    justify-content: space-between;
    flex-wrap: wrap;
    overflow: hidden;

    div {
      height: 320px;
      border-radius: 15px;
      box-shadow: 0px 0px 6px 0px rgba(0, 0, 0, 0.3);
      margin: 13px 0;
      padding: 20px;
      box-sizing: border-box;
    }

    h2 {
      font-size: 50px;
      font-weight: bold;
      letter-spacing: 0.1em;
      margin: 0;
    }

    p {
      font-size: 26px;
      line-height: 40px;
      font-weight: lighter;
      letter-spacing: 0.1em;
    }

    &-item1 {
      color: #000;
      //width: 58.8%;
      background-color: #FDC900;
    }

    &-item2 {
      //width: 38.8%;
      background: #3662EC;
    }

    &-item3 {
      //width: 36.5%;
      background: #2DA35D;
    }

    &-item4 {
      //width: 36.5%;
      background: #B35EEB;
    }

    &-item5 {
      //width: 22%;
      background: #374F72;
    }
  }
}

//内容4
.content4 {
  text-align: center;
  margin-top: 150px;

  &-text {
    h1 {
      font-size: 50px;
      font-weight: bold;
      letter-spacing: 0.1em;
    }

    p {
      font-size: 30px;
      font-weight: 500;
      letter-spacing: 0.3em;
      margin-top: 30px;
      line-height: 50px;
    }

    .content4-btn {
      margin-top: 50px;
      font-size: 28px;
      letter-spacing: 0.3em;
      border-radius: 8px;
      padding: 30px 50px;
      border: 0;
    }
  }
}
</style>