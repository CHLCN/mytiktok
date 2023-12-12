<template>
  <div class="Publish">
    <BaseHeader mode="light" backMode="dark" backImg="back"> </BaseHeader>
    <div class="content">
      <div class="desc">
        <div class="title">发布</div>
        <div class="main">
          <textarea
            class="left"
            rows="5"
            v-model="des"
            placeholder="添加文字描述..."
          ></textarea>
          <div
            class="right"
            :style="{
              backgroundImage: 'url(' + bgimg + ')',
              backgroundSize: '50px 100px',
            }"
            ref="right"
          >
            <input
              type="file"
              ref="uploadInput"
              id="uploadBtn"
              class="file"
              accept="video/MP4"
              @change="handleFileUpload"
            />

            <input type="button" class="btn" @click="clickFile" />
          </div>
        </div>
      </div>
      <dy-button
        class="pub-btn"
        type="primary"
        :loading="loading"
        :active="false"
        :disabled="disabled"
        @click="publish"
        >发布
      </dy-button>
    </div>
  </div>
</template>
<script>
import request from "../../utils/request";
export default {
  name: "Publish",
  data() {
    return {
      video: null,
      des: "",
      thumbnailUrl: "",
      bgimg: "../../src/assets/img/upload/video.svg",
      loading: false,
    };
  },
  computed: {
    disabled() {
      return !(this.video && this.des);
    },
  },
  mounted() {},
  methods: {
    clickFile() {
      const input = document.querySelector(".file");
      input.click();
    },
    handleFileUpload() {
      //   const file = event.target.files[0];
      let base64URL = "";
      this.video = this.$refs.uploadInput.files[0];
      //   console.log(this.video);
      const video = document.createElement("video");
      video.setAttribute("crossOrigin", "anonymous"); //处理跨域
      video.setAttribute("src", URL.createObjectURL(this.video));
      video.currentTime = 1;

      video.addEventListener("loadeddata", function () {
        let canvas = document.createElement("canvas");
        //使用视频的宽高作为canvas、预览图的宽高
        let width = video.videoWidth;
        let height = video.videoHeight;
        canvas.width = width;
        canvas.height = height;
        canvas.getContext("2d").drawImage(video, 0, 0, width, height); //绘制canvas
        base64URL = canvas.toDataURL("image/jpeg"); //转换为base64，图片格式默认为png，这里修改为jpeg
      });
      setTimeout(() => {
        this.bgimg = base64URL;
        // console.log(this.bgimg);
      }, 500);
    },
    async publish() {
      let user_id = this.$store.state.user_id;
      //   const link = document.createElement("a");
      //   link.href = this.bgimg;
      //   link.download = ".image.png"; // 设置下载的文件名
      //   link.click(); // 触发下载
      this.loading = true;
      let res = await request.post(
        "/publish/action",
        {},
        {
          params: {
            author_id: user_id,
            title: this.des,
            play_url: "/src/assets/video/" + this.video.name,
            cover_url:
              "/src/assets/img/thumb/" + this.video.name.slice(0, -4) + ".png",
            // cover_url: this.bgimg,
          },
        }
      );
      //   console.log(res.data);
      if (res.data.status_code == 0) {
        setTimeout(() => {
          this.loading = false;
        }, 1000);
        this.$router.push("/home");
      }
    },
  },
};
</script>

<style scoped lang="less">
@import "../../assets/less/index";

.Publish {
  position: fixed;
  left: 0;
  right: 0;
  bottom: 0;
  top: 0;
  overflow: auto;
  color: black;
  font-size: 14rem;
  background: white;
  .content {
    text-align: center;
    padding-left: 10rem;
    padding-right: 10rem;
    display: flex;
    flex-direction: column;
    height: 100vh;
    .desc {
      flex: 1;
      .title {
        font-size: large;
        font-weight: bold;
        margin-top: 20rem;
      }
      .main {
        margin-top: 20rem;
        margin-bottom: 20rem;
        display: flex;
        .left {
          flex: 2;
          outline: none;
          border: none;
          background-color: whitesmoke;
          font-size: 16px;
          font-family: Arial;
          color: black;
          border-radius: 4px;
          padding: 10px;
          margin-right: 10px;
        }
        .right {
          flex: 1;
          background: center center no-repeat;
          background-size: 40px 40px;
          opacity: 0.8;
          z-index: 999;
          //   background: #fff;
          border: 1px dashed #d7dde4;
          border-radius: 4px;
          text-align: center;
          overflow: hidden;
          transition: border-color 0.2s ease;
          margin-left: 3px;

          .file {
            // 将type="file"隐藏
            display: none;
          }
          .btn {
            height: 100%;
            width: 100%;
            flex: 1;
            background: rgba(0, 0, 0, 0);
            z-index: 1;
            border: none;
            cursor: pointer;
          }
        }
      }
    }

    .pub-btn {
      position: relative;
      bottom: 30rem;
    }
  }
}
</style>
