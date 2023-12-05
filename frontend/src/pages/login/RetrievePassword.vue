<template>
  <div class="RetrievePassword">
    <BaseHeader mode="light" backMode="dark" backImg="back" />
    <Loading v-if="loading.getPhone" />
    <div class="content">
      <div class="desc">
        <div class="title">找回密码</div>
        <div class="sub-title">
          验证码已通过短信发送到+86 {{ this.$store.state.phone }}。 密码为8 -
          20位，至少包含字母、数字、符号2种组合
        </div>
      </div>

      <LoginInput
        autofocus
        type="code"
        v-model="code"
        placeholder="请输入验证码"
        v-model:isSendVerificationCode="isSendVerificationCode"
        @send="sendCode"
      />
      <LoginInput
        class="mt1r"
        autofocus
        type="password"
        v-model="password"
        placeholder="请输入密码"
      />
      <div class="tip" id="tip">验证码或密码错误，请重试</div>

      <div class="protocol" :class="showAnim ? 'anim-bounce' : ''">
        <Tooltip style="top: -150%; left: -10rem" v-model="showTooltip" />
        <div class="left">
          <Check v-model="isAgree" />
        </div>
        <div class="right">
          已阅读并同意
          <span
            class="link"
            @click="$nav('/service-protocol', { type: '“抖音”用户服务协议' })"
            >用户协议</span
          >
          和
          <span
            class="link"
            @click="$nav('/service-protocol', { type: '“抖音”隐私政策' })"
            >隐私政策</span
          >
          以及
          <span
            class="link"
            @click="$nav('/service-protocol', { type: '“抖音”运营商服务协议' })"
            >运营商服务协议</span
          >
          ，运营商将对你提供的手机号进行验证
        </div>
      </div>

      <div class="notice" v-if="notice">
        {{ notice }}
      </div>

      <dy-button
        type="primary"
        :loading="loading"
        :active="false"
        :loadingWithText="true"
        :disabled="disabled"
        @click="changepwd"
      >
        完成
      </dy-button>
    </div>
  </div>
</template>
<script>
import Check from "../../components/Check";
import LoginInput from "./components/LoginInput";
import Tooltip from "./components/Tooltip";
import Base from "./Base";
import request from "../../utils/request";

export default {
  name: "RetrievePassword",
  extends: Base,
  components: {
    Check,
    Tooltip,
    LoginInput,
  },
  data() {
    return {
      phone: "",
      password: "",
      code: "",
      notice: "",
      isSendVerificationCode: true,
      showVoiceCode: false,
    };
  },
  computed: {
    disabled() {
      return !(this.code && this.password);
    },
  },
  created() {
    this.getCode();
    this.phone = this.$store.state.phone;
  },
  methods: {
    async sendCode() {
      this.$showLoading();
      await this.$sleep(500);
      this.$hideLoading();
      this.isSendVerificationCode = true;
      let res = await request.post(
        "/user/getcode",
        {},
        {
          params: {
            username: this.$store.state.phone,
          },
        }
      );
      console.log(res.data.status_msg);
      this.$store.state.vcode = res.data.status_msg;
    },
    async getCode() {
      let res = await request.post(
        "/user/getcode",
        {},
        {
          params: {
            username: this.$store.state.phone,
          },
        }
      );
      console.log(res.data.status_msg);
      this.$store.state.vcode = res.data.status_msg;
    },
    async changepwd() {
      let isCheck = await this.check();
      if (isCheck) {
        this.loading = true;
      }
      if (this.code != this.$store.state.vcode) {
        console.log(
          "this code: " + this.code + " store code: " + this.$store.state.vcode
        );
        document.getElementsByClassName("tip")[0].removeAttribute("id");
        this.loading = false;
      }

      let res = await request.post(
        "/user/password_change",
        {},
        {
          params: {
            username: this.phone,
            password: this.password,
          },
        }
      );
      //   console.log(res.data.status_code);
      if (res.data.status_code == 0) {
        setTimeout(() => {
          this.$router.replace("/login/password");
        }, 500);
      }
    },
  },
};
</script>

<style scoped lang="less">
@import "../../assets/less/index";
@import "Base.less";

.RetrievePassword {
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
    padding: 60rem 30rem;
    .desc {
      margin-bottom: 30rem;
      margin-top: 30rem;
      display: flex;
    }
    .title {
      margin-bottom: 5rem;
      font-size: large;
    }
    .protocol {
      margin-top: 5rem;
      margin-bottom: 10rem;
      .left {
        margin-right: 5rem;
      }
    }

    .tip {
      color: red;
      margin-top: 2rem;
    }
    #tip {
      display: none;
    }

    .button {
      margin-top: 20rem;
    }
  }
}
</style>
