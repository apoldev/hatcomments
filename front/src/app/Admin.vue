<template >
  <div class="hatcomments">

    <div class="hatcomment-header">
      <div class="hatcomment-header-part">admin panel</div>
      <div class="hatcomment-header-part" v-if="!store.signin">
        <div class="action action-signin" @click="actionOnSignInOrOut()" :class='{"action-signout": store.user}'>

          <div class="icon">
            <svg viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" clip-rule="evenodd" d="M6 8a1 1 0 0 0 1-1V5.923c0-.459.022-.57.082-.684a.364.364 0 0 1 .157-.157c.113-.06.225-.082.684-.082h10.154c.459 0 .57.022.684.082.07.038.12.087.157.157.06.113.082.225.082.684v12.154c0 .459-.022.57-.082.684a.363.363 0 0 1-.157.157c-.113.06-.225.082-.684.082H7.923c-.459 0-.57-.022-.684-.082a.363.363 0 0 1-.157-.157c-.06-.113-.082-.225-.082-.684V17a1 1 0 1 0-2 0v1.077c0 .76.082 1.185.319 1.627.223.419.558.753.977.977.442.237.866.319 1.627.319h10.154c.76 0 1.185-.082 1.627-.319.419-.224.753-.558.977-.977.237-.442.319-.866.319-1.627V5.923c0-.76-.082-1.185-.319-1.627a2.363 2.363 0 0 0-.977-.977C19.262 3.082 18.838 3 18.077 3H7.923c-.76 0-1.185.082-1.627.319a2.363 2.363 0 0 0-.978.977C5.083 4.738 5 5.162 5 5.923V7a1 1 0 0 0 1 1zm9.593 2.943c.584.585.584 1.53 0 2.116L12.71 15.95c-.39.39-1.03.39-1.42 0a.996.996 0 0 1 0-1.41 9.552 9.552 0 0 1 1.689-1.345l.387-.242-.207-.206a10 10 0 0 1-2.24.254H2.998a1 1 0 1 1 0-2h7.921a10 10 0 0 1 2.24.254l.207-.206-.386-.241a9.562 9.562 0 0 1-1.69-1.348.996.996 0 0 1 0-1.41c.39-.39 1.03-.39 1.42 0l2.883 2.893z"/></svg>
          </div>

          <div class="label" v-if="!store.user">Войти</div>

          <div class="label" v-if="store.user">Выйти</div>

        </div>

        <div class="action action-signin" v-if="store.user">
          <div class="action-signin-avatar" >
            <img :src="store.user.info.image" :alt="username">
            <!--            <div class="label" v-if="user">{{ username }}</div>-->
          </div>
        </div>

      </div>

      <div class="hatcomment-header-part" v-if="store.signin">
        <div class="action action-signin" @click="store.setViewSignin(false)">
          <div class="icon">
            <svg viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
              <path d="M4 10L3.29289 10.7071L2.58579 10L3.29289 9.29289L4 10ZM21 18C21 18.5523 20.5523 19 20 19C19.4477 19 19 18.5523 19 18L21 18ZM8.29289 15.7071L3.29289 10.7071L4.70711 9.29289L9.70711 14.2929L8.29289 15.7071ZM3.29289 9.29289L8.29289 4.29289L9.70711 5.70711L4.70711 10.7071L3.29289 9.29289ZM4 9L14 9L14 11L4 11L4 9ZM21 16L21 18L19 18L19 16L21 16ZM14 9C17.866 9 21 12.134 21 16L19 16C19 13.2386 16.7614 11 14 11L14 9Z" fill="#33363F"/>
            </svg>
          </div>
          <div class="label" >Назад</div>
        </div>
      </div>


    </div>

    <sign-in v-if="store.signin" :onlysocials="true" @onRegister="registerUser"/>


    <div v-if="store.user">

      <UsersComponent />

    </div>



  </div>
</template>

<script>
import {API_URL} from "@/api/config";

import {GetUser, SaveUser} from "@/store/storage";
import {useStore} from "@/store/store";
import {$adminApi, $api} from "@/api/api";
import {GoogleAuth} from "@/api/oauth";
import SignIn from "@/components/Signin";
import UsersComponent from "@/components/admin/Users";

export default {
  name: 'Admin',
  components: {
    UsersComponent,
    SignIn

  },

  setup(){

    const store = useStore()

    return {
      store
    }

  },
  data(){

    return {
      baseUrl: API_URL,
      google: false,
    }
  },

  async mounted() {

    let user = await GetUser()
    user = await this.auth(user)

    $adminApi.setToken(user.token)

    this.store.setUser(user)

  },
  computed:{
    username(){
      return (this.store.user?.info?.first_name + " " + this.store.user?.info?.last_name).trim()
    },
  },
  methods: {
    async signOut(){
      this.store.setUser(null)
      await SaveUser(null)
    },

    actionOnSignInOrOut(){

      if(this.store.user){
        this.signOut()
        return
      }

      this.store.setViewSignin(true)

    },

    async auth(user){

      if(!user){
        return null
      }

      const resp = await $api.post("/user/auth", {token: user.token})
      return resp.data

    },

    async registerUser(userData){

      let user = await this.auth(userData)

      this.store.setUser(user)
      this.store.setViewSignin(false)

      SaveUser(user)

    },

    openGoogle(){
      this.google = true

      // запустим гугл окно
      GoogleAuth((e) => {

        if(e){
          this.onGoogleToken(e)
          return
        }


        this.google = false

      })

    },
  }

}
</script>

<style lang="scss">
@import "../../css/app";
</style>
