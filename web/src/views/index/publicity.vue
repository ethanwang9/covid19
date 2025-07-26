<template>
  <el-row justify="center">
    <el-col :span="22">
      <!--菜单-->
      <Menu></Menu>

      <!--防疫宣传内容-->
      <el-row class="content">
        <el-col :md="16" class="main">
          <el-col class="main-title">
            <h1>防疫宣传</h1>
          </el-col>
          <el-col class="main-info">
            <h2 id="title1">新型冠状病毒感染的肺炎通用预防指南</h2>
            <img alt="防疫宣传" src="@/assets/images/index/publicity/1.jpg">
          </el-col>
        </el-col>
        <el-col :md="6" class="side hidden-sm-and-down">
          <el-col class="side-info">
            <!-- COVID-19 info-->
            <el-col class="side-info-ins">
              <el-col :span="10">
                <img alt="COVID-19" src="@/assets/icon/du.png">
              </el-col>
              <el-col :span="14">
                <h3>COVID-19</h3>
                <p>2019年12月</p>
              </el-col>
              <el-col>
                <p>新型冠状病毒肺炎（新冠肺炎，COVID-19）为新发急性呼吸道传染病，目前已成为全球性重大的公共卫生事件。</p>
                <p>
                  潜伏期1～14天，多为3～7天。以发热、干咳、乏力为主要表现。部分患者以嗅觉、味觉减退或丧失等为首发症状，少数患者伴有鼻塞、流涕、咽痛、结膜炎、肌痛和腹泻等症状。</p>
              </el-col>
            </el-col>
          </el-col>
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
import Footer from '@/components/index/Footer.vue'</script>

<style lang="scss" scoped>
.content {
  display: flex;
  justify-content: center;

  & > * {
    margin: 10px;
  }

  .main, .side > * {
    color: #000;
    background-color: #fff;
    border-radius: 8px;
    box-shadow: 0 0 8px rgba(0, 0, 0, 0.3);
    padding: 15px;
  }


  .main {
    &-title {
      h1 {
        text-align: center;
        margin: 30px 0 60px;
        letter-spacing: 6px;
      }
    }

    &-info {
      img {
        width: 100%;
        vertical-align: middle;
        border-radius: 8px;
      }

      h2 {
        letter-spacing: 2px;
      }
    }
  }

  .side {
    height: 300px;

    * {
      margin-bottom: 15px;
    }

    a {
      text-decoration: unset;
      color: #000;
      cursor: pointer;
    }

    h1 {
      margin: 0;
      letter-spacing: 2px;
    }

    p {
      letter-spacing: 2px;
    }

    &-info {
      &-ins {
        display: flex;
        justify-content: center;
        align-items: center;
        flex-direction: column;

        img {
          width: 80px;
          height: 80px;
        }

        h3 {
          margin: 0;
        }

        p {
          margin: 10px 0 0;
        }
      }
    }
  }
}
</style>