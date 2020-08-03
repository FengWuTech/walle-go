<template>
  <div class="container-main">
    <div class="left"></div>
    <div class="right">
      <div class="from">
        <h4 class="from-haed" style="margin-bottom:60px">账号登录</h4>
        <div class="main">
          <a-form
            id="formLogin"
            ref="formLogin"
            layout="vertical"
            class="user-layout-login"
            :form="form"
            @submit="handleSubmit"
          >
            <a-form-item label="用户名">
              <a-input
                v-decorator="[
                  'email',
                  {rules: [{ required: true, message: '请输入帐户名' }], validateTrigger: 'change'}
                ]"
                size="large"
                type="text"
                placeholder="请输入要登录的账号"
              />
            </a-form-item>

            <a-form-item label="密码">
              <a-input
                v-decorator="[
                  'password',
                  {rules: [{ required: true, message: '请输入密码' }], validateTrigger: 'blur'}
                ]"
                class="input-item"
                size="large"
                type="password"
                autocomplete="new-password"
                placeholder="请输入要登录的密码"
              />
              <!-- <div class="forge-password-main" @click="forgetpswfn">
                <a class="forge-password" style="float: right;">忘记密码?</a>
              </div> -->
            </a-form-item>
            <a-form-item>
              <a-button
                block
                type="primary"
                html-type="submit"
                class="login-button"
                :loading="state.loginBtn"
                :disabled="state.loginBtn"
              >登录</a-button>
            </a-form-item>
          </a-form>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { mapActions } from 'vuex'
import { getProviderInfo } from '@/api/user'
export default {
  name: 'Login',
  components: {
  },
  data () {
    return {
      loginBtn: false,
      forgetpsw: false,
      loginType: 0,
      data: [],
      isLoginError: false,
      requiredTwoStepCaptcha: false,
      stepCaptchaVisible: false,
      form: this.$form.createForm(this),
      query: this.$route.query,
      isLoginErrorText: '账户或密码错误',
      from: 'pc',
      WeChatUri: 'https://open.work.weixin.qq.com/wwopen/sso/3rd_qrConnect',
      state: {
        time: 60,
        loginBtn: false,
        loginQyBtn: false,
        // login type: 0 email, 1 username, 2 telephone
        loginType: 0,
        wxBtn: false
      }
    }
  },
  mounted () {
    document.body.classList.add('userLayout')
  },
  beforeDestroy () {
    document.body.classList.remove('userLayout')
  },
  created () {
    this.query = this.$route.query
  },
  methods: {
    ...mapActions(['Login', 'Logout']),
    getWxAuthCode () {
      this.state.wxBtn = true
      getProviderInfo()
        .then(res => {
          const redirectUri = encodeURIComponent(window.location.href)
          location.href = `${this.WeChatUri}?appid=${res.data.provider_corp_id}&redirect_uri=${redirectUri}&usertype=member`
        })
        .catch(() => {
          this.state.wxBtn = false
        })
    },
    handleSubmit (e) {
      e.preventDefault()
      const {
        form: { validateFields },
        state
      } = this
      state.Btn = true
      validateFields((err, values) => {
        if (!err) {
          try {
            values.from = this.from
            values.redirect_url = this.query.from
            if (this.data.length === 1) {
              values.company_id = this.data[0].id
            }
            this.$store
              .dispatch('login', values)
              .then(res => {
                this.$router.replace({
                  path: '/index'
                })
              })
              .catch(err => {
                if (err.code === 401) {
                  this.$notification.destroy()
                  this.$notification['error']({
                    message: '温馨提示',
                    description: '用户不存在',
                    duration: 4
                  })
                } else {
                  this.isLoginError = true
                  this.isLoginErrorText = err.msg
                  this.$notification.destroy()
                  this.$notification['error']({
                    message: '温馨提示',
                    description: err.msg,
                    duration: 4
                  })
                }
              })
          } catch (error) {
            // console.log(error)
          }
        } else {
          state.loginBtn = false
        }
      })
    }
  }
}
</script>

<style lang="less">

/*媒体查询*/
/*当页面大于1200px 时，大屏幕，主要是PC 端*/
@media (min-width: 1200px) {
    .container-main{
      .left{
        width: 45% !important;
      }
      .right .from{
        margin-top: 150px !important;
      }
       .right .user-layout-login .ant-form-item {
          margin-bottom: 30px !important;
      }
    }
}
/*在992 和1199 像素之间的屏幕里，中等屏幕，分辨率低的PC*/
@media (min-width: 992px) and (max-width: 1199px) {
    .container-main{
      .left{
        width: 500px !important;
      }
    }
}

/*在小于767 像素的屏幕，微小屏幕，更低分辨率的手机*/
@media (max-width: 992px) {
    .container-main{
      .left{
        display: none;
      }
      .right .from{
        margin-left: 0 !important;
        .from-haed{
          margin-left: 12px !important;
        }
      }
      .right .user-layout-login   {
          margin: 0 auto !important;
      }
    }
}

.container-main{
  // display: flex;
  height: 100%;
  .left{
    // width: 600px;
    // flex:0 0 600px;
    background: #eff3f6 url(~@/assets/log-bg1.png);
    background-repeat: no-repeat;
    background-size: auto 100%;
    position: relative;
    float: left;
    height: 100%;
    min-height: 700px;
    width: 600px;
    // margin-right: 20px;    //形成20px的间隔
    .content{
      top: 35%;
      position: absolute;
      left: 50%;
      text-align: center;
      transform: translateX(-50%);
      .com-name{
        margin-top: 50px;
        margin-bottom: 6px;
        height:30px;
        font-size:24px;
        font-family:PingFangSC-Medium,PingFang SC;
        font-weight:500;
        color:rgba(243,189,140,1);
        line-height:30px;
        letter-spacing:1px;
      }
      .dec{
        height:18px;
        font-size:14px;
        font-family:PingFangSC-Light,PingFang SC;
        font-weight:300;
        color:rgba(243,189,140,1);
        line-height:18px;
        letter-spacing:1px;
      }
    }
    .footer{
      position: absolute;
      left: 32px;
      bottom: 20px;
      height:20px;
      font-size:14px;
      font-family:PingFangSC-Regular,PingFang SC;
      font-weight:400;
      color:rgba(255,255,255,1);
      line-height:20px;
    }
  }
  .right{
    // flex: 1;
    overflow: hidden;
    height: 100%;
    min-height: 700px;
    background:rgba(239,243,246,1);
    .from {
      margin-left: 100px;
      margin-top: 80px;
    }
    .user-layout-login{
      width: 370px;
      .ant-form-item{
        padding-bottom: 0;
        margin-bottom: 18px;
      }
      .ant-form-item-required::before{
        display: none;
      }
      .ant-form-item-label{
        label{
          font-size:14px;
          font-family:PingFangSC-Regular,PingFang SC;
          font-weight:400;
          color:rgba(106,119,159,1);
          line-height:19px;
        }
      }
      .ant-input:hover {
          border-color: rgba(106,119,159,0.4);
          border-right-width: 1px !important;
      }
      .ant-input:focus {
          border-color: rgba(106,119,159,0.4);
      }
      .ant-btn-primary{
        background-color: #5D6EA4;
        border-color: #5D6EA4;
      }
      .ant-btn-primary:disabled{
          color: #fff;
          opacity: 0.7;
      }
      .login-bak-button,
      .wx-login {
          border-color: #5D6EA4 !important;
          color: #5D6EA4;
      }
    }
  }
  .m-t-sm {
    height:19px;
    font-size:12px;
    margin-top: 50px;
    font-family:PingFangSC-Regular,PingFang SC;
    font-weight:400;
    color:rgba(93,110,164,1);
    line-height:19px;
    a{
      color:rgba(93,110,164,1);
    }
    // text-align: center;
  }
}

.top-nav {
  display: none;
  position: fixed;
  left: 0;
  top: 0;
  height: 66px;
  width: 100%;
  z-index: 1;
}
.head-index {
  float: right;
  margin-top: 6px;
}
.home {
  background-image: url(~@/assets/home.png);
  display: inline-block;
  width: 32px;
  height: 32px;
  background-size: cover;
}
.info {
  height: 22px;
  font-size: 16px;
  font-family: PingFangSC-Regular, PingFang SC;
  font-weight: 400;
  color: rgba(50, 144, 255, 1);
  line-height: 22px;
  margin: 30px 0 100px;
  text-align: center;
}
.img-main {
  text-align: center;
  padding-top: 20px;
}

.from .has-error .ant-form-explain,
.from .has-error .ant-form-split {
  color: #f5222d;
  margin-bottom: 6px;
}

.img-main {
  max-width: 725px;
  margin: 0 auto;
}
.img-main img {
  width: 100%;
}
.from-haed {
  font-size:32px;
  font-family:PingFangSC-Medium,PingFang SC;
  font-weight:500;
  color:rgba(38,38,38,1);
  line-height:45px;
}
.input-item:-webkit-autofill,
select:-webkit-autofill {
  -webkit-box-shadow: 0 0 0px 1000px white inset !important;
}

.input-main {
  padding: 60px 70px;
}

.from .input-item:focus {
  outline: 0;
  -webkit-box-shadow: 0 0 0 0 transparent;
  box-shadow: 0 0 0 0 transparent;
}
.has-error .input-item:focus {
  border-color: #ff4d4f;
  border-right-width: 1px !important;
  outline: 0;
  -webkit-box-shadow: 0 0 0 0 transparent;
  box-shadow: 0 0 0 0 transparent;
}
.input-item:focus,
.input-item:hover {
  border-color: #40a9ff;
  border-right-width: 1px !important;
}
.has-error .input-item:focus,
.has-error .input-item:hover {
  border-color: #f5222d;
  border-right-width: 1px !important;
}
.has-error .ant-form-explain {
  margin-top: 5px !important;
}
.wx-login {
  border-color: rgba(7, 193, 96, 1) !important;
  color: rgba(7, 193, 96, 1);
}
#userLayout.user-layout-wrapper .container .ant-row-flex-center {
  position: absolute;
  top: 40%;
  left: 50%;
  /* background-color: #000; */
  width: 100%;
  height: 50%;
  -webkit-transform: translateX(-50%) translateY(-50%);
}

@media only screen and (min-width: 1920px) {
  #userLayout.user-layout-wrapper .container .ant-row-flex-center {
    position: absolute;
    top: 45%;
    left: 50%;
    /* background-color: #000; */
    width: 100%;
    height: 50%;
    -webkit-transform: translateX(-50%) translateY(-50%);
  }
}

@media (min-width: 768px) {
  .top-nav {
    display: block;
    background: #fff;
    border-bottom: 1px solid #f2f2f2;
    -webkit-box-shadow: 0 2px 4px rgba(0, 0, 0, 0.08);
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.08);
  }

  .top-nav .top-nav-content {
    // width: 1180px;
    margin: 0 auto;
    padding: 10px 50px;
  }

  .top-nav .top-nav-content .menu {
    float: right;
    height: 60px;
    margin: 0;
    list-style: none;
    display: -ms-flexbox;
    display: flex;
    -ms-flex-align: center;
    align-items: center;
  }

  .top-nav .top-nav-content .menu li a {
    display: inline-block;
    padding: 10px 18px;
    font-size: 15px;
    color: #686969;
  }
}

#userLayout.user-layout-wrapper {
  height: 100%;

  &.mobile {
    .container {
      .main {
        max-width: 368px;
        width: 98%;
      }
    }
  }

  .container {
    width: 100%;
    min-height: 100%;
    // background: #f0f2f5 url(~@/assets/background.svg) no-repeat 50%;
    background: linear-gradient(
      140deg,
      rgba(15, 130, 249, 1) 0%,
      rgba(19, 165, 255, 1) 100%
    );
    opacity: 0.8;
    background-size: 100%;
    position: relative;

    a {
      text-decoration: none;
    }

    .top {
      text-align: center;

      .header {
        height: 44px;
        line-height: 44px;

        .badge {
          position: absolute;
          display: inline-block;
          line-height: 1;
          vertical-align: middle;
          margin-left: -12px;
          margin-top: -10px;
          opacity: 0.8;
        }

        .logo {
          height: 44px;
          vertical-align: top;
          margin-right: 16px;
          border-style: none;
        }

        .title {
          font-size: 33px;
          color: #000;
          font-family: Avenir, "Helvetica Neue", Arial, Helvetica, sans-serif;
          font-weight: 600;
          position: relative;
          top: 2px;
        }
      }
      .desc {
        font-size: 14px;
        color: rgba(0, 0, 0, 0.45);
        margin-top: 12px;
        margin-bottom: 40px;
      }
    }

    // .main {
    //   min-width: 260px;
    //   width: 368px;
    //   margin: 0 auto;
    // }

    .footer {
      position: fixed;
      width: 230px;
      bottom: 0;
      /* margin-left: 115px; */
      /* padding: 0 16px; */
      margin: 0 0 24px;
      text-align: center;
      left: 50%;
      transform: translate(-50%, -50%);

      .links {
        margin-bottom: 8px;
        font-size: 14px;
        a {
          color: rgba(0, 0, 0, 0.45);
          transition: all 0.3s;
          &:not(:last-child) {
            margin-right: 40px;
          }
        }
      }
      .copyright {
        // color: rgba(0, 0, 0, 0.45);
        color: #fff;
        font-size: 14px;
      }
    }
  }
}

.user-layout-login {
  label {
    font-size: 14px;
  }
  .qiyewx {
    display: inline-block;
    width: 29px;
    height: 24px;
    margin-right: 10px;
    background-image: url("~@/assets/qiyeweixin.png");
    background-size: cover;
    background-repeat: no-repeat;
    vertical-align: middle;
  }
  .getCaptcha {
    display: block;
    width: 100%;
    height: 40px;
  }
  .forge-password-main {
    position: absolute;
    top: 24px;
    right: 0;
    .forge-password{
        margin-top: 12px;
        font-size:14px;
        font-family:PingFangSC-Regular,PingFang SC;
        font-weight:400;
        color:rgba(93,110,164,1);
        line-height:19px;
    }
  }

  .ant-form-explain {
    position: absolute;
    top: 4px;
    right: -313px;
    width: 300px;
  }
  .forge-password {
    font-size: 14px;
  }

  button.login-button {
    padding: 0 15px;
    font-size: 16px;
    height: 40px;
    width: 100%;
  }

  .user-login-other {
    text-align: left;
    margin-top: 24px;
    line-height: 22px;

    .item-icon {
      font-size: 24px;
      color: rgba(0, 0, 0, 0.2);
      margin-left: 16px;
      vertical-align: middle;
      cursor: pointer;
      transition: color 0.3s;

      &:hover {
        color: #1890ff;
      }
    }

    .register {
      float: right;
    }
  }
}
</style>
