<template >
  <div class="hatcomments">
    <!--    <div>{{room}}</div>-->
<!--    <div>{{store.room}}</div>-->

    <div class="hatcomments-info" v-if="store.newcomments.length > 0" @click="scrollToComment(store.newcomments[0])">
      <span>Новых комментариев:</span>
      <span class="hatcomments-info-count">{{store.newcomments.length}}</span>

      <div class="hatcomments-info-close" @click="clearNewComments()">
        <svg viewBox="0 0 1024 1024" xmlns="http://www.w3.org/2000/svg"><path fill="#ffffff" d="M195.2 195.2a64 64 0 0 1 90.496 0L512 421.504 738.304 195.2a64 64 0 0 1 90.496 90.496L602.496 512 828.8 738.304a64 64 0 0 1-90.496 90.496L512 602.496 285.696 828.8a64 64 0 0 1-90.496-90.496L421.504 512 195.2 285.696a64 64 0 0 1 0-90.496z"/></svg>

      </div>
    </div>


    <div class="hatcomment-header">


      <div class="hatcomment-header-part" >
        <div class="action action-clients" v-if="!store.signin">
          <div class="icon">
            <svg viewBox="0 0 519.578 519.578">
              <path fill="#cfd6da" d="M513.095,245.101c0,0-140.683-139.842-253.291-139.842c-112.608,0-253.292,139.842-253.292,139.842c-8.645,8.109-8.721,21.42,0,29.375c0,0,140.684,139.843,253.292,139.843c112.608,0,253.291-139.843,253.291-139.843C521.663,266.368,521.816,253.134,513.095,245.101z M260.875,372.397c-61.889,0-112.149-50.185-112.149-112.149s50.184-112.149,112.149-112.149c61.965,0,112.148,50.26,112.148,112.149S322.763,372.397,260.875,372.397z"></path>
              <path fill="#cfd6da" d="M259.574,187.726c-39.321,0-71.222,32.053-71.222,71.451c0,39.397,31.901,71.451,71.222,71.451c39.321,0,71.222-32.054,71.222-71.451C330.796,219.78,298.896,187.726,259.574,187.726z M286.426,259.407c-12.163,0-22.108-9.946-22.108-22.262s9.945-22.261,22.108-22.261s22.108,9.945,22.108,22.261S298.742,259.407,286.426,259.407z"></path></svg>

          </div>
          <div class="label" v-if="numClients"> {{ numClients }}</div>
          <div class="action-users">
            <img :src="v" v-for="(v,i) in avatars" :style="'left: ' + (i*15).toString() + 'px'" :key="'kk'+i"/>
          </div>
        </div>
      </div>

      <div class="hatcomment-header-part" v-if="!store.signin">
        <div class="action action-signin" @click="actionOnSignInOrOut()" :class='{"action-signout": user}'>

          <div class="icon">
            <svg viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" clip-rule="evenodd" d="M6 8a1 1 0 0 0 1-1V5.923c0-.459.022-.57.082-.684a.364.364 0 0 1 .157-.157c.113-.06.225-.082.684-.082h10.154c.459 0 .57.022.684.082.07.038.12.087.157.157.06.113.082.225.082.684v12.154c0 .459-.022.57-.082.684a.363.363 0 0 1-.157.157c-.113.06-.225.082-.684.082H7.923c-.459 0-.57-.022-.684-.082a.363.363 0 0 1-.157-.157c-.06-.113-.082-.225-.082-.684V17a1 1 0 1 0-2 0v1.077c0 .76.082 1.185.319 1.627.223.419.558.753.977.977.442.237.866.319 1.627.319h10.154c.76 0 1.185-.082 1.627-.319.419-.224.753-.558.977-.977.237-.442.319-.866.319-1.627V5.923c0-.76-.082-1.185-.319-1.627a2.363 2.363 0 0 0-.977-.977C19.262 3.082 18.838 3 18.077 3H7.923c-.76 0-1.185.082-1.627.319a2.363 2.363 0 0 0-.978.977C5.083 4.738 5 5.162 5 5.923V7a1 1 0 0 0 1 1zm9.593 2.943c.584.585.584 1.53 0 2.116L12.71 15.95c-.39.39-1.03.39-1.42 0a.996.996 0 0 1 0-1.41 9.552 9.552 0 0 1 1.689-1.345l.387-.242-.207-.206a10 10 0 0 1-2.24.254H2.998a1 1 0 1 1 0-2h7.921a10 10 0 0 1 2.24.254l.207-.206-.386-.241a9.562 9.562 0 0 1-1.69-1.348.996.996 0 0 1 0-1.41c.39-.39 1.03-.39 1.42 0l2.883 2.893z"/></svg>
          </div>

          <div class="label" v-if="!user">Войти</div>

          <div class="label" v-if="user">Выйти</div>

        </div>

        <div class="action action-signin" v-if="user">
          <div class="action-signin-avatar" >
            <img :src="user.info.image" :alt="username">
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

    <Signin v-if="store.signin" @onRegister="registerUser" :onlysocials="true"/>

    <TextInput v-show="!store.signin" @onSend="onReplyTo($event)" :attachments="[]"/>

    <div class="comments" v-show="!store.signin">
      <CommentList
          :comments="comments"
          @onReplyTo="onReplyTo($event)"
          @onCallParent="onCallParent"
      />

      <div class="comments-load-more" v-if="can_loading_more">
        <span v-if="loading_comments">
            <LoadingSvg fill="#222"/>
          </span>
        <a href="#" @click.prevent="loadMore()">
          <span v-if="this.comments.length > 0">Загрузить еще</span>
        </a>
      </div>
    </div>
  </div>

</template>

<script>
import TextInput from '../components/TextInput.vue'
import CommentList from '../components/CommentList.vue'
import Signin from '../components/Signin.vue'
import {$api} from "@/api/api";
import {TreeComments, findId} from '../helpers/tree-comments'
import {Linkify} from '../helpers/anchors'
import {SaveUser, GetUser} from '../store/storage'
import {useStore} from '../store/store'
import LoadingSvg from "@/components/Loading";
import {API_URL} from "@/api/config";


export default {
  name: 'App',
  components: {
    LoadingSvg,
    TextInput,
    CommentList,
    Signin
  },
  props: ["d"],
  setup(){

    const store = useStore()

    return {
      store
    }

  },
  data(){

    return {
      baseUrl: API_URL,
      comments: [],
      room: {
        id: 0,
        slug: "",
        name: "",
        project_id: 0,
      },
      presence: {},
      signin: false,
      user: null,
      loading_comments: false,
      offset: 0,
      limit: 20,
      can_loading_more: true,

      sub: null,
      sub_offset: 0,
    }
  },

  async mounted() {

    this.room.name = this.$config.room
    this.room.project_id = this.$config.project;


    this.store.room.name = this.$config.room
    this.store.room.project_id = this.$config.project;
    this.store.selector = this.$config.selector

    await this.loadCommentProject()


    let user = await GetUser()
    user = await this.auth(user)

    this.user = user
    this.store.setUser(user)

    this.$centrifuge.init(user?.token, this.callCentrifugeConnected)
    this.sub = this.centrifugeGetSub()
    await this.centrifugeConnect()

  },

  created() {

  },

  computed: {
    username(){
      return (this.user?.info?.first_name + " " + this.user?.info?.last_name).trim()
    },
    uniqPresenceByUser(){

      const data = {};

      if(Object.keys(this.presence).length > 0){
        for(let i in this.presence){
          data[this.presence[i].user] = this.presence[i]
        }

      }

      return data;

    },
    numClients(){
      return Object.keys(this.uniqPresenceByUser).length
    },
    avatars(){

      const avatars = {}

      if(this.numClients > 0){
        for(let i in this.uniqPresenceByUser){

          let p = this.uniqPresenceByUser[i]

          const info = p.conn_info || p.connInfo

          if(info && info?.image){
            avatars[p.user] = info.image
          }
        }
      }

      return Object.values(avatars).slice(0, 5);
    }
  },
  methods: {

    clearNewComments(){
      this.store.newcomments = []
    },

    scrollToComment(comment){
      const y = document.getElementById("hatcomment-" + comment.id).getBoundingClientRect().top + window.scrollY;
      window.scroll({
        top: y - 25,
        behavior: 'smooth'
      });


      setTimeout(() => {
        this.store.removeCommentFromInfo(comment)
      }, 500)

    },

    onCallParent(comment){


      comment["hightlight"] = true
      setTimeout(() => {
        comment["hightlight"] = false
      }, 1000)


      this.scrollToComment(this.comment)

    },


    async auth(user){

      if(!user){
        return null
      }

      const resp = await $api.post("/user/auth", {token: user.token})
      return resp.data

    },


    actionOnSignInOrOut(){

      if(this.store.user){
        this.signOut()
        return
      }

      this.store.setViewSignin(true)

    },
    async loadMore(){


      if(this.comments.length > 0 && this.offset === 0){
        return
      }

      await this.loadCommentProject()

    },

    async signOut(){

      this.user = null;

      this.store.setUser(null)

      await SaveUser(null)

      this.centrifugeReconnect(null)

    },

    async registerUser(userData){

      let user = await this.auth(userData)
      this.user = user

      this.store.setUser(user)
      this.store.setViewSignin(false)

      SaveUser(user)

      // Переподключимся к центрифуге с новым токеном
      this.centrifugeReconnect(user.token)

    },

    async callCentrifugeConnected(ctx){


      if(this.sub_offset > 0) {
        const h = await this.sub.history({limit: 100, since: {offset: this.sub_offset}});
        // console.log("[h]", h)
        this.sub_offset = h.offset
        for(let i in h.publications){
          this.subFunc(h.publications[i])
        }
      }

    },

    textTransformer(text){


      text = Linkify(text)

      if(this.$actions.textTransformer){
        text = this.$actions.textTransformer(text)
      }



      return text

    },

    subFunc(ctx){

      if(!ctx.data){
        return
      }

      // console.log("[ctx]", ctx)
      this.sub_offset = ctx.offset

      const data = ctx.data.data
      const method = ctx.data.method

      if(method === "comment") {

        data.children = [];
        data.votes = [];

        data.orig = {
          text: data.text,
          attachments: [...data.attachments]
        }
        data.text = this.textTransformer(data.text)

        data.new = true

        if(data?.level === 0){
          this.comments.unshift(data)

        }else{

          let findComment = findId(data.reply_to, this.comments)

          if(findComment){
            findComment.children.push(data)

          }
        }

      }

      if(method === "edit") {

        let findComment = findId(data.comment_id, this.comments)
        if(findComment){

          findComment.orig = {
            text: data.text,
            attachments: [...data.attachments]
          }
          findComment.text = this.textTransformer(data.text)
          findComment.attachments = data.attachments

        }
      }

      if(method === "vote") {

        let findComment = findId(data.comment_id, this.comments)
        if(findComment){
          findComment.like = data.likes
          findComment.votes = data.votes
        }
      }


      if(method === "delete") {
        let findComment = findId(data.id, this.comments)
        if(findComment){
          findComment.deleted = true
          findComment.deleted_by_id = data.deleted_by_id
        }
      }


      if(method === "restore") {
        let findComment = findId(data.id, this.comments)
        if(findComment){
          findComment.deleted = false
          findComment.deleted_by_id = null
        }
      }


    },

    centrifugeGetSub(){

      const channel = "comments:room-" + this.room.slug;

      const sub = this.$centrifuge.subscribe(channel, this.subFunc,
          (join) => {
            // console.log("join", join)
            if(join.info && join.info?.client){
              this.presence[join.info.client] = join.info
            }
          },
          (leave) => {
            if(leave.info && leave.info?.client){
              delete this.presence[leave.info.client]
            }
          }, (error) => {
            console.log(error)
          })

      return sub
    },
    centrifugeReconnect(token){

      this.$centrifuge.disconnect()
      this.$centrifuge.init(token, this.callCentrifugeConnected)
      this.sub = this.centrifugeGetSub()
      this.centrifugeConnect()

    },

    async centrifugeConnect(){

      this.$centrifuge.connect()

      const pp = await this.sub.presence()
      if(pp?.clients){
        this.presence = pp.clients
      }


    },

    async loadCommentProject(){

      this.loading_comments = true

      try {

        let response = await $api.get("/project/"+this.room.project_id+ "?limit=" +this.limit+ "&offset=" +this.offset+ "&room=" + encodeURIComponent(this.room.name))

        const comments = response.data.comments || []

        comments.forEach((item =>{

          item.orig = {
            text: item.text,
            attachments: [...item.attachments]
          }

          item.text = this.textTransformer(item.text)


        }))

        this.comments = this.comments.concat(TreeComments(comments))


        if(comments.length > 0){
          this.offset = comments[comments.length - 1].id
        }

        if(this.offset > 0 && comments.length === 0){
          this.can_loading_more = false
        }

        // todo
        this.room = response.data.room
        // pinia
        this.store.room = response.data.room

      }catch (e) {
        console.log(e)
      }finally {

        this.loading_comments = false
      }

    },

    onReplyTo: async function(data) {
      if(data?.comment){
        data.comment['reply'] = false;
      }
    }
  }
}
</script>

<style lang="scss">
@import "../../css/app";
</style>
