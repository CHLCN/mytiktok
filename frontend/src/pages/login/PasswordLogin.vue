<template>
  <div class="PasswordLogin">
    <BaseHeader mode="light" backMode="dark" backImg="back">
      <template v-slot:right>
        <span class="f14" @click="$nav('/login/help')">帮助与设置</span>
      </template>
    </BaseHeader>
    <div class="content">
      <div class="desc">
        <div class="title">手机号密码登录</div>
      </div>

      <LoginInput
        autofocus
        type="phone"
        v-model="phone"
        placeholder="请输入手机号"
      />
      <LoginInput
        autofocus
        class="mt1r"
        type="password"
        v-model="password"
        placeholder="请输入密码"
      />

      <div class="tip" id="tip">账号或密码错误，请重试</div>
      <div class="tip_phone" id="tip">请输入正确的手机号</div>

      <div class="protocol" :class="showAnim ? 'anim-bounce' : ''">
        <Tooltip style="top: -150%; left: -10rem" v-model="showTooltip" />
        <div class="left">
          <Check v-model="isAgree" />
        </div>
        <div class="right" style="font-size: 12rem;">
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
        </div>
      </div>

      <div class="notice" v-if="notice">
        {{ notice }}
      </div>

      <dy-button
        type="primary"
        :loading="loading"
        :active="false"
        :disabled="disabled"
        @click="login"
      >
        {{ loading ? "登录中" : "登录" }}
      </dy-button>

      <div class="options" style="font-size: 12rem;">
        <span>
          忘记了？<span class="link" @click="isJump">找回密码</span>
        </span>
      </div>
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
  name: "PasswordLogin",
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
      isRight: true,
    };
  },
  computed: {
    disabled() {
      return !(this.phone && this.password);
    },
  },
  created() {},
  methods: {
    async login() {
      let isCheck = await this.check();
      if (isCheck) {
        this.loading = true;
        let res = await request.post(
          "/user/password_login",
          {},
          {
            params: {
              username: this.phone,
              password: this.password,
            },
          }
        );
        if (res.data.status_code == 0) {
          this.$store.state.phone = this.phone;
          this.$store.state.user_id=res.data.user_id;
          this.$router.push("/home");
        } else {
          document.getElementsByClassName("tip")[0].removeAttribute("id");
          this.loading = false;
        }
      }
    },
    async isJump() {
      // console.log("this is store phone: "+this.$store.state.phone)
      if (this.phone != "") {
        let res = await request.get("/user/number_existence", {
          params: {
            username: this.phone,
          },
        });
        if (res.data.status_code != 0)
          document.getElementsByClassName("tip_phone")[0].removeAttribute("id");
        else {
          this.$store.state.phone = this.phone;
          this.$store.state.getPhone = true;
          document.getElementsByClassName("tip")[0].setAttribute("id", "tip");
          document
            .getElementsByClassName("tip_phone")[0]
            .setAttribute("id", "tip");
          await this.$router.push("/login/retrieve-password");
        }
      } else {
        document.getElementsByClassName("tip_phone")[0].removeAttribute("id");
      }
    },
  },
};
</script>

<style scoped lang="less">
@import "../../assets/less/index";
@import "Base.less";

.PasswordLogin {
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
    .tip {
      color: red;
      margin-top: 2rem;
    }
    .tip_phone {
      color: red;
      margin-top: 2rem;
    }
    #tip {
      display: none;
    }

    .protocol {
      margin-top: 5rem;
      margin-bottom: 10rem;
      .left {
        margin-right: 5rem;
      }
    }

    .title {
      font-size: large;
    }
    .options {
      margin-top: 15rem;
      color: @second-text-color;
    }
  }
}
</style>
