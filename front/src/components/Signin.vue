<template>
  <div class="signin-wrapper">
  <div class="signin">

    <h3>Войти в аккаунт</h3>


    <div class="signin-button" @click="openGoogle()">
      <div class="signin-icon">

        <svg xmlns="http://www.w3.org/2000/svg"  viewBox="0 0 48 48" width="48px" height="48px"><path fill="#fbc02d" d="M43.611,20.083H42V20H24v8h11.303c-1.649,4.657-6.08,8-11.303,8c-6.627,0-12-5.373-12-12	s5.373-12,12-12c3.059,0,5.842,1.154,7.961,3.039l5.657-5.657C34.046,6.053,29.268,4,24,4C12.955,4,4,12.955,4,24s8.955,20,20,20	s20-8.955,20-20C44,22.659,43.862,21.35,43.611,20.083z"/><path fill="#e53935" d="M6.306,14.691l6.571,4.819C14.655,15.108,18.961,12,24,12c3.059,0,5.842,1.154,7.961,3.039	l5.657-5.657C34.046,6.053,29.268,4,24,4C16.318,4,9.656,8.337,6.306,14.691z"/><path fill="#4caf50" d="M24,44c5.166,0,9.86-1.977,13.409-5.192l-6.19-5.238C29.211,35.091,26.715,36,24,36	c-5.202,0-9.619-3.317-11.283-7.946l-6.522,5.025C9.505,39.556,16.227,44,24,44z"/><path fill="#1565c0" d="M43.611,20.083L43.595,20L42,20H24v8h11.303c-0.792,2.237-2.231,4.166-4.087,5.571	c0.001-0.001,0.002-0.001,0.003-0.002l6.19,5.238C36.971,39.205,44,34,44,24C44,22.659,43.862,21.35,43.611,20.083z"/></svg>

      </div>
      <span class="signin-label">
        <span>Войти через Google</span>
        <LoadingSvg fill="#444" v-if="google"/>
      </span>

    </div>


    <div class="signin-button" @click="openVK()">
      <div class="signin-icon">


        <svg xmlns="http://www.w3.org/2000/svg"  viewBox="0 0 48 48" width="48px" height="48px"><path fill="#1976d2" d="M24 4A20 20 0 1 0 24 44A20 20 0 1 0 24 4Z"/><path fill="#fff" d="M35.937,18.041c0.046-0.151,0.068-0.291,0.062-0.416C35.984,17.263,35.735,17,35.149,17h-2.618 c-0.661,0-0.966,0.4-1.144,0.801c0,0-1.632,3.359-3.513,5.574c-0.61,0.641-0.92,0.625-1.25,0.625C26.447,24,26,23.786,26,23.199 v-5.185C26,17.32,25.827,17,25.268,17h-4.649C20.212,17,20,17.32,20,17.641c0,0.667,0.898,0.827,1,2.696v3.623 C21,24.84,20.847,25,20.517,25c-0.89,0-2.642-3-3.815-6.932C16.448,17.294,16.194,17,15.533,17h-2.643 C12.127,17,12,17.374,12,17.774c0,0.721,0.6,4.619,3.875,9.101C18.25,30.125,21.379,32,24.149,32c1.678,0,1.85-0.427,1.85-1.094 v-2.972C26,27.133,26.183,27,26.717,27c0.381,0,1.158,0.25,2.658,2c1.73,2.018,2.044,3,3.036,3h2.618 c0.608,0,0.957-0.255,0.971-0.75c0.003-0.126-0.015-0.267-0.056-0.424c-0.194-0.576-1.084-1.984-2.194-3.326 c-0.615-0.743-1.222-1.479-1.501-1.879C32.062,25.36,31.991,25.176,32,25c0.009-0.185,0.105-0.361,0.249-0.607 C32.223,24.393,35.607,19.642,35.937,18.041z"/></svg>
      </div>
      <span class="signin-label">
        <span>Войти через VK</span>
        <LoadingSvg fill="#444" v-if="vk"/>
      </span>

    </div>

    <template v-if="!onlysocials">
      <div class="signin-or">
        <span class="h3">или</span>

        <span>Введите имя для гостевого доступа</span>
      </div>

      <div class="signin-input" :class='{"error": error}'>

        <input v-model="name" type="text" placeholder="Ваше имя..." name="name"/>
      </div>

      <div class="signin-button signin-guest" @click="guestAuth()" :class='{"disabled": name.length < 1}'>
        <div class="signin-icon">

          <svg viewBox="0 0 16 16" xmlns="http://www.w3.org/2000/svg">
            <path d="m 8 1 c -1.65625 0 -3 1.34375 -3 3 s 1.34375 3 3 3 s 3 -1.34375 3 -3 s -1.34375 -3 -3 -3 z m -1.5 7 c -2.492188 0 -4.5 2.007812 -4.5 4.5 v 0.5 c 0 1.109375 0.890625 2 2 2 h 8 c 1.109375 0 2 -0.890625 2 -2 v -0.5 c 0 -2.492188 -2.007812 -4.5 -4.5 -4.5 z m 0 0" fill="#fff"/>
          </svg>

        </div>
        <span class="signin-label">
        <span>Гостевой доступ</span>
        <LoadingSvg fill="#fff" v-if="guest"/>
      </span>
      </div>
    </template>

  </div>
  </div>
</template>

<script>

import {$api} from "@/api/api"
import {GoogleAuth, VKAuth} from "@/api/oauth"
import LoadingSvg from "@/components/Loading";

export default {
  name: "SignIn",
  components: {LoadingSvg},
  props: ["onlysocials"],
  data(){
    return {
      name: "",
      error: false,
      google: false,
      vk: false,
      guest: false,
    }
  },
  methods: {

    openVK(){
      this.vk = true
      VKAuth((e) => {
        if(e){
          this.onToken({vk_token: e.vk_token, email: e.email})
          return
        }
        this.vk = false
      })
    },
    openGoogle(){
      this.google = true

      // запустим гугл окно
      GoogleAuth((e) => {

        if(e){
          this.onToken({google_token: e})
          return
        }


        this.google = false

      })

    },


    async onToken(e){

      this.google = false
      this.vk = false


      try{

        const resp = await $api.post("/user/register", e)
        this.$emit("onRegister", resp.data)

      }catch (e) {
        console.log(e)
      }



    },
    async guestAuth(){

      this.guest = true

      try{
        const resp = await $api.post("/user/register", {name: this.name})
        this.$emit("onRegister", resp.data)

      }catch(e){


        console.log(e)

        this.error = true

        setTimeout(() => {
          this.error = false
        }, 1000)

      }finally {
        this.guest = false
      }

    }
  }
}
</script>