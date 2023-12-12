<template>
  <div class="other-login" id="other-login">
    <!-- <BaseHeader mode="light" backMode="dark" backImg="back">
      <template v-slot:right>
        <span class="f14" @click="$nav('/home')">帮助与设置</span>
      </template>
    </BaseHeader> -->
    <div class="content">
      <div class="desc">
        <div class="title">登录后，体验完整功能</div>
        <div class="protocol" :class="showAnim?'anim-bounce':''">
        <Tooltip style="top: -100%;left: -10rem;" v-model="showTooltip"/>
        <div class="left">
          <Check v-model="isAgree"/>
        </div>
        <div class="right" style="font-size: 12rem;">
          已阅读并同意
          <span class="link" @click="$nav('/service-protocol',{type:'“抖音”用户服务协议'})">用户协议</span>
          和
          <span class="link" @click="$nav('/service-protocol',{type:'“抖音”隐私政策'})">隐私政策</span>
          以及
          <span class="link" @click="$nav('/service-protocol',{type:'“抖音”运营商服务协议'})">运营商服务协议</span>
          ，运营商将对你提供的手机号进行验证
        </div>
      </div>
      </div>

      
      <LoginInput autofocus type="phone" v-model="phone" placeholder="请输入手机号"/>
      <div class="sub-title">未注册的手机号验证通过后将自动注册</div>
      <div class="notice" v-if="notice">
        {{ notice }}
      </div>

 

      <dy-button type="primary" :loading="loading" :active="false" :disabled="phone.length < 11" @click="getCode">
        获取短信验证码
      </dy-button>

      <div class="options" style="font-size: 12rem;">
        <span class="link" @click="$nav('/login/password')">密码登录</span>
        <span class="link" @click="otherLogin">其他方式登录</span>
      </div>

      <from-bottom-dialog
          page-id="other-login"
          v-model="isOtherLogin"
          :show-heng-gang="false"
          height="270rem"
          mode="white">
        <div class="block-dialog">
          <div class="item" @click="$no">
            <img src="../../assets/img/icon/login/toutiao-round.png" alt="">
            <span>今日头条登录</span>
          </div>
          <div class="item" @click="$no">
            <img src="../../assets/img/icon/login/qq.webp" alt="">
            <span>QQ登录</span>
          </div>
          <div class="item" @click="$no">
            <img src="../../assets/img/icon/login/wechat.webp" alt="">
            <span>微信登录</span>
          </div>
          <div class="item" @click="$no">
            <img src="../../assets/img/icon/login/weibo.webp" alt="">
            <span>微博登录</span>
          </div>
          <div class="space"></div>
          <div class="item" @click="isOtherLogin = false">取消</div>
        </div>
      </from-bottom-dialog>

    </div>
  </div>
</template>
<script>
import Check from "../../components/Check";
import Tooltip from "./components/Tooltip";
import LoginInput from "./components/LoginInput";
import Base from "./Base.js";
import FromBottomDialog from "../../components/dialog/FromBottomDialog";
import request from '../../utils/request'

export default {
  name: "OtherLogin",
  extends: Base,
  components: {
    Check,
    Tooltip,
    LoginInput,
    FromBottomDialog
  },
  data() {
    return {
      phone: '',
      notice: ''
    }
  },
  computed: {},
  created() {
  },
  methods: {
    async getCode() {
        this.$store.state.phone=this.phone;
      let isCheck = await this.check()
      if (isCheck) {
        this.loading = true
        let res = await request.post(
            '/user/login',
            {},
            {
                params:{
                    username:this.phone,
                }
            }
        )
        this.$store.state.vcode=res.data.status_msg;
        this.$store.state.user_id=res.data.user_id;

        setTimeout(()=>{
            this.$router.push("/login/verification-code")
            this.loading=false
        },1500)
        
        
        
      }
    },
    async otherLogin() {
      let isCheck = await this.check()
      if (isCheck) {
        this.isOtherLogin = true
      }
    }
  }
}
</script>

<style scoped lang="less">
@import "../../assets/less/index";
@import "Base.less";

.other-login {
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
    //   align-items: center;
    //   flex-direction: column;

    .left{
        margin-right: 5rem;
    }

    }

    .sub-title{
        font-size:12rem;
        color: @second-text-color;
        margin-top: 5rem;
        margin-bottom: 15rem;
    }
    .title{

        margin-bottom: 5rem;
        font-size: large;
    }
  .options{
    margin-top: 15rem;  
  }
}
 

  .block-dialog {
    color: black;

    .item {
      height: 50rem;
      display: flex;
      justify-content: center;
      align-items: center;
      border-top: 1px solid gainsboro;

      img {
        height: 25rem;
        margin-right: 10rem;
      }

      &:nth-last-child(1) {
        border-top: none;
      }
    }

    .space {
      height: 10rem;
      background: whitesmoke;
    }
  }


}


</style>
