<template>
  <div class="VerificationCode">
    <BaseHeader mode="light" backMode="dark" backImg="back">
      <template v-slot:right>
        <span class="f14" @click="$nav('/login/help')">帮助与设置</span>
      </template>
    </BaseHeader>
    <div class="content">
      <div class="desc">
        <div class="title">请输入验证码</div>
        <div class="sub-title">验证码已通过短信发送到+86 {{this.$store.state.phone}}</div>
      </div>

      <LoginInput autofocus type="code"
                  v-model="code"
                  placeholder="请输入验证码"
                  v-model:isSendVerificationCode="isSendVerificationCode"
                  @send="sendCode"
      />
      <div class="tip"  id="tip">验证码错误，请重试</div>

      <dy-button  type="primary"  :loading="loading" :active="false" :disabled="code.length < 4" @click="login">
        {{ loading ? '登录中' : '登录' }}
      </dy-button>

    </div>
  </div>
</template>
<script>
import Check from "../../components/Check";
import LoginInput from "./components/LoginInput";
import request from '../../utils/request'

export default {
  name: "VerificationCode",
  components: {
    Check,
    LoginInput,
  },
  data() {
    return {
      showAnim: false,
      showTooltip: false,
      loading: false,
      phone: '',
      password: '',
      code: '',
      isSendVerificationCode: true,
      showVoiceCode: false
    }
  },
  created() {
    setTimeout(() => {
      this.showVoiceCode = true
    }, 3000)
  },
  methods: {

    async sendCode() {
      this.$showLoading()
      await this.$sleep(500)
      this.$hideLoading()
      this.isSendVerificationCode = true
      let res = await request.post(
            '/user/login',
            {},
            {
                params:{
                    username:this.$store.state.phone,
                }
            }
        )
    },
    login() {
      this.loading = true
      if(this.code==this.$store.state.vcode){
        setTimeout(() => {
        this.isSendVerificationCode = true
        this.$router.push("/home")
        this.loading = false
      }, 1000)
      }else{
        
        setTimeout(()=>{
            document.getElementsByClassName("tip")[0].removeAttribute("id");
            this.loading=false;
        },500)
        
      }
      
    }
  }
}
</script>

<style scoped lang="less">
@import "../../assets/less/index";
@import "Base.less";

.VerificationCode {
  position: fixed;
  left: 0;
  right: 0;
  bottom: 0;
  top: 0;
  overflow: auto;
  color: black;
  font-size: 14rem;
  background: white;

  .content{
    padding: 60rem 30rem;
    .desc {
      margin-bottom: 30rem;
      margin-top: 30rem;
      display: flex;
    }
    .title{
    margin-bottom: 5rem;
    font-size: large;
    }

    .tip{
        color: red;
        margin-top: 2rem;
    }
    #tip{
        display: none;
    }

  }

  .button{
    margin-top: 20rem;
  }

}
</style>
