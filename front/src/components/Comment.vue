<template>
  <div class="comment" :id="'hatcomment-'+comment.id">
    <div @click="view()" class="comment-user" :class='{"comment-new": isNewNew, "comment-highlight": comment.hightlight}' :data-id="'hatcomment-' + comment.id">
      <span href="#" class="comment-img" >
        <img :src="comment.user?.image ? comment.user?.image : apiUrl + '/user/avatar?name=' + name" :alt="name">

        <AchievementIcon v-if="comment.user.icon !== ''" :user="comment.user" />

      </span>

      <div class="comment-message">
        <div class="comment-author-wrapper">
          <div>

            <div class="comment-author-withreply">
              <a class="comment-author" href="#">{{name}}</a>
              <span class="comment-reply-to-user" v-if="replyToName" @click="findParent()">
            <span class="icon">
              <svg  version="1.1" id="_x32_" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" viewBox="0 0 512 512"  xml:space="preserve">
                <g>
                  <path class="st0" d="M292.497,168.968c-21.134,0-40.287,0-57.542,0V65.394L0,255.995l234.955,190.61V334.395
                    c7.132,0,14.331,0,21.578,0c95.305,0,227.772-2.396,237.359,100.701C541.847,322.408,501.086,168.968,292.497,168.968z"/>
                </g>
              </svg>
            </span>
            <span class="label">{{replyToName}}</span>
          </span>
            </div>

            <span class="comment-created-at">{{humanDate(comment.created_at)}}</span>

          </div>



          <div class="more" :class='{"opened": more}' @click="toggleMoreActions" v-if="canEdit || canDelete || canView || canRestore">
            <svg width="24px" height="24px" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
              <path d="M5.70711 9.71069C5.31658 10.1012 5.31658 10.7344 5.70711 11.1249L10.5993 16.0123C11.3805 16.7927 12.6463 16.7924 13.4271 16.0117L18.3174 11.1213C18.708 10.7308 18.708 10.0976 18.3174 9.70708C17.9269 9.31655 17.2937 9.31655 16.9032 9.70708L12.7176 13.8927C12.3271 14.2833 11.6939 14.2832 11.3034 13.8927L7.12132 9.71069C6.7308 9.32016 6.09763 9.32016 5.70711 9.71069Z"/>
            </svg>

            <div class="more-wrapper" v-show="more">

              <span class="reply-action-btn text-action-btn" v-if="!comment.deleted && canEdit">
                <a href="#" @click.prevent="edit()" >Редактировать</a>
              </span>

              <span class="reply-action-btn text-action-btn" v-if="!comment.deleted && canDelete">
                <a href="#" @click.prevent="deleteComment()" >Удалить</a>
              </span>

              <span class="reply-action-btn text-action-btn" v-if="hidden && comment.deleted && canView">
                <a href="#" @click.prevent="viewDeletedComment(false)" >Просмотреть</a>
              </span>

              <span class="reply-action-btn text-action-btn" v-if="!hidden && comment.deleted && canView">
                <a href="#" @click.prevent="viewDeletedComment(true)" >Скрыть</a>
              </span>

              <span class="reply-action-btn text-action-btn" v-if="comment.deleted && canRestore">
                <a href="#" @click.prevent="restoreComment()" >Восстановить</a>
              </span>

            </div>
          </div>

        </div>

        <div v-if="isEdit" class="hatcomment-editmode">
          <TextInput
             :text="comment.orig.text"
             :attach="comment.orig.attachments"
             :editmode="true"
             :comment="comment"
             @cancelEdit="cancelEdit"
             @onSuccessEdit="onSuccessEdit"
          />
        </div>


        <template v-else>
          <div class="hatcomment-content" :class='{"deleted": hidden}'>
            <p v-if="!hidden"  v-html="comment.text"></p>

            <p v-else >сообщение удалено</p>

          </div>

          <div class="comment-attachments" v-if="!hidden && comment.attachments.length > 0">
            <AttachmentList :attachments="comment.attachments" />
          </div>



          <div class="reply-to-wrapper" >
          <span class="reply-action-btn text-action-btn" v-if="!comment.deleted">
            <a href="#" @click.prevent="replyOpen()" >Ответить</a>
          </span>

            <div class="like-wrapper">
            <span v-if="comment.like" class="like-count" :class="{'negative': comment.like < 0}">
              {{comment.like > 0 ? "+" : ""}}{{comment.like}}
            </span>

              <span class="icon" v-if="loading_like">
              <Loading fill="#a3a3a3"/>
            </span>

              <span v-else @click="like(1)" class="icon reply-action-btn">
              <svg class="like" :class='{"active": myReaction == 1}' version="1.1" viewBox="0 0 33.867 33.867" xmlns="http://www.w3.org/2000/svg"><g transform="translate(0 -244.08)"><path d="m20.457 275.96c-1.5247-0.26113-2.7181-0.57778-5.1853-1.3758-1.8822-0.60878-2.4846-0.78039-3.6956-1.0529-0.8605-0.19365-0.8797-0.35787-0.8797-7.5208 0-7.4101-0.07864-6.9469 1.2894-7.594 0.97126-0.45949 1.2861-0.80315 5.5285-6.0338 0.50421-0.62168 0.74363-1.2792 1.266-3.4771 0.34479-1.4507 0.80203-2.3677 1.4017-2.8111 0.38162-0.28217 0.49151-0.29102 1.5251-0.12275 1.4897 0.24254 2.315 0.76367 2.9278 1.8487 0.4627 0.8193 0.47371 0.88492 0.47371 2.8236 0.17821 1.5156-0.09467 2.1806-0.72442 3.4832-0.39843 0.82413-0.72443 1.5781-0.72444 1.6755-1.6e-5 0.12841 0.82739 0.17704 3.0121 0.17704 3.5985 0 4.0594 0.11866 5.143 1.3238 0.3798 0.42244 0.69054 0.86358 0.69054 0.98033s0.30502 0.71575 0.30502 1.4177c0 1.0956-0.04852 1.2913-0.53275 2.148-0.50516 0.89371-0.52388 0.97777-0.36147 1.6228 0.20918 0.83073 0.08061 1.886-0.33419 2.7428-0.2218 0.45818-0.32146 1.006-0.34818 1.9137-0.03815 1.2963-0.36164 2.2408-0.92867 2.7114-0.16963 0.14078-0.24089 0.47287-0.24326 1.1336-0.0074 2.069-1.3088 3.5833-3.4738 4.0422-1.0616 0.22502-4.6837 0.19202-6.131-0.0558zm-17.504-16.975c-0.62703 0-1.0101 6.1e-4 -1.2602 0.0176-0.25007 0.0169-0.38206 0.0533-0.4715 0.13814-0.044719 0.0424-0.080509 0.10302-0.10245 0.19083-0.021937 0.0878-0.034387 0.20722-0.0432 0.391-0.017627 0.36759-0.018515 0.99443-0.018515 2.1178v8.7227c0 1.1233 8.847e-4 1.7502 0.018515 2.1178 0.00881 0.18378 0.021262 0.30322 0.0432 0.39101 0.021938 0.0878 0.057711 0.14844 0.10245 0.19082 0.089447 0.0847 0.22143 0.1212 0.4715 0.13814 0.25006 0.0169 0.63319 0.0164 1.2602 0.0164h4.7372c0.62703 0 1.0102 5.4e-4 1.2602-0.0164 0.25005-0.0169 0.38212-0.0546 0.4715-0.13931 0.044687-0.0424 0.080536-0.10187 0.10245-0.18965 0.02191-0.0878 0.034402-0.20723 0.043203-0.39101 0.017603-0.36756 0.018515-0.99442 0.018515-2.1178v-8.7227c0-1.1234-8.846e-4 -1.7502-0.018515-2.1178-0.00885-0.18378-0.021293-0.3032-0.043203-0.391s-0.057777-0.14841-0.10245-0.19083c-0.089363-0.0849-0.22145-0.12118-0.4715-0.13814-0.25005-0.017-0.63319-0.0176-1.2602-0.0176z"></path></g></svg>
            </span>


              <span class="icon" v-if="loading_dislike">
              <Loading fill="#a3a3a3"/>
            </span>

              <span v-else @click="like(-1)" class="icon reply-action-btn">
              <svg class="dislike" :class='{"active": myReaction < 0}' version="1.1" viewBox="0 0 33.867 33.867" xmlns="http://www.w3.org/2000/svg"><path d="m20.457 1.987c-1.5247 0.26113-2.7181 0.57778-5.1853 1.3758-1.8822 0.60878-2.4846 0.78039-3.6956 1.0529-0.8605 0.19365-0.87971 0.35787-0.87971 7.5208 0 7.4101-0.0786 6.9469 1.2894 7.594 0.97126 0.45949 1.2861 0.80315 5.5285 6.0338 0.50421 0.62168 0.74363 1.2792 1.266 3.4771 0.34479 1.4507 0.80203 2.3677 1.4017 2.8111 0.38162 0.28217 0.49151 0.29102 1.5251 0.12275 1.4897-0.24254 2.315-0.76367 2.9278-1.8487 0.4627-0.8193 0.47371-0.88492 0.47371-2.8236 0.17821-1.5156-0.0947-2.1806-0.72442-3.4832-0.39843-0.82413-0.72443-1.5781-0.72445-1.6755-1e-5 -0.12841 0.8274-0.17704 3.0121-0.17704 3.5984 0 4.0594-0.11866 5.143-1.3238 0.3798-0.42244 0.69054-0.86358 0.69054-0.98033s0.30502-0.71575 0.30502-1.4177c0-1.0956-0.0485-1.2913-0.53275-2.148-0.50515-0.89371-0.52388-0.97777-0.36147-1.6228 0.20918-0.83073 0.0806-1.886-0.33419-2.7428-0.2218-0.45818-0.32146-1.006-0.34818-1.9137-0.0382-1.2963-0.36164-2.2408-0.92867-2.7114-0.16963-0.14078-0.24089-0.47287-0.24326-1.1336-7e-3 -2.069-1.3088-3.5833-3.4738-4.0422-1.0616-0.22502-4.6837-0.19202-6.131 0.0558zm-17.504 16.975c-0.62703 0-1.0101-6.1e-4 -1.2602-0.0176-0.25008-0.0169-0.38207-0.0533-0.4715-0.13814-0.0447-0.0424-0.0805-0.10302-0.10245-0.19083-0.0219-0.0878-0.0344-0.20722-0.0432-0.391-0.0176-0.36759-0.0185-0.99443-0.0185-2.1178v-8.7227c0-1.1233 8.8e-4 -1.7502 0.0185-2.1178 9e-3 -0.18378 0.0213-0.30322 0.0432-0.39101 0.0219-0.0878 0.0577-0.14844 0.10245-0.19082 0.0894-0.0847 0.22143-0.1212 0.4715-0.13814 0.25006-0.0169 0.63319-0.0164 1.2602-0.0164h4.7372c0.62703 0 1.0102-5.3e-4 1.2602 0.0164 0.25004 0.0169 0.38212 0.0546 0.4715 0.13931 0.0447 0.0424 0.0805 0.10187 0.10244 0.18965 0.0219 0.0878 0.0344 0.20723 0.0432 0.39101 0.0176 0.36756 0.0185 0.99442 0.0185 2.1178v8.7227c0 1.1234-8.7e-4 1.7502-0.0185 2.1178-9e-3 0.18378-0.0213 0.3032-0.0432 0.391s-0.0578 0.14841-0.10244 0.19083c-0.0894 0.0849-0.22145 0.12118-0.4715 0.13814-0.25005 0.017-0.63319 0.0176-1.2602 0.0176z"></path></svg>
            </span>



            </div>
          </div>

          <div class="reply-wrapper" v-if="comment.reply">
            <TextInput :parent="comment" @onSend="onSendReply($event)"/>
          </div>
        </template>
      </div>

    </div>

    <div class="comment-level" v-if="comment.children.length > 0">
      <Comment
          class="comment"
          v-for="c in comment.children"
          :key="c.id"
          :comment="c"
          :parent="comment"
          @onReplyTo="this.$emit('onReplyTo', $event)"
          @onClickReply="this.$emit('onClickReply', $event)"
          @onCallParent="this.$emit('onCallParent', $event)"
      />
    </div>
  </div>
</template>

<script>

import TextInput from './TextInput.vue'
import AttachmentList from './AttachmentList.vue'
import {sendVote, deleteComment, restoreComment} from './../api/centrifuge'
import {useStore} from "@/store/store";
import Loading from "@/components/Loading";
import {API_URL} from '@/api/config';
import {getIntervalUpdateComment} from './../helpers/dates'
import {IsVisible} from './../helpers/visible_element'
import AchievementIcon from "@/components/Achievement";

export default {
  name: 'MyComment',
  components: {
    AchievementIcon,
    Loading,
    TextInput,
    AttachmentList

  },
  props: {
    comment: Object,
    parent: Object,
  },
  setup(){

    return {
      store: useStore(),
    }
  },

  mounted() {


    if(this.comment.new){

      setTimeout(() => {
        this.isNew = true
      }, 100)

      const el = document.getElementById('hatcomment-' + this.comment.id)
      if(el){

        const removeNewMarker = () => {
          setTimeout(() => {
            this.isNew = false
          }, 30 * 1000)
        }

        const handleVisible = () => {
          const isv = IsVisible(el)

          if(isv){
            window.removeEventListener("scroll", handleVisible)
            removeNewMarker()

          }
        }

        if(!IsVisible(el)){
          window.addEventListener("scroll", handleVisible)

          this.store.addNewCommentToInfo(this.comment)

        }else{
          removeNewMarker()
        }
      }


    }


    let timeout = getIntervalUpdateComment(this.comment.created_at)
    let interval;

    const f = () => {
      return () => {
        this.$forceUpdate()
        let newTimeout = getIntervalUpdateComment(this.comment.created_at)
        if(newTimeout !== timeout) {
          clearInterval(interval)
          interval = setInterval(f(), 1000 * newTimeout)
        }
      }
    }

    interval = setInterval(f(), 5000 * timeout)


  },

  unmounted(){

  },
  computed:{
    isNewNew(){

      if(this.store.user){
        return this.isNew && this.store.user?.info?.id !== this.comment.user?.id
      }

      return this.isNew

    },

    myID(){
      return this.store.user?.info?.id
    },

    moderator(){
      if(this.store.user){
        return this.store.user?.info?.role === "1"
      }
      return false
    },
    admin(){
      if(this.store.user){
        return this.store.user?.info?.role === "99"
      }
      return false
    },

    owner(){
      if(this.store.user){
        return this.store.user?.info?.id === this.comment?.user?.id
      }
      return false
    },

    canView(){
      return ((this.owner || this.moderator) && this.comment.deleted_by_id === this.myID) || this.admin

      // return this.owner || this.moderator || this.admin
    },
    canDelete(){
      return this.owner || this.moderator || this.admin
    },
    canEdit(){
      return this.owner
    },

    canRestore(){
      return ((this.owner || this.moderator) && this.comment.deleted_by_id === this.myID) || this.admin
    },
    replyToName(){

      if(this.comment?.reply_to_user){
        return (this.comment?.reply_to_user?.first_name + " " + this.comment?.reply_to_user?.last_name).trim()
      }

      return null
    },
    myReaction(){

      if(this.store.user){

        for(let i in this.comment.votes){
          if(this.comment?.votes[i]?.user_id === this.store.user?.info?.id){
            return this.comment?.votes[i].vote;
          }
        }
      }

      return 0
    },


    humanDate(){
      return function(time){

        const formats = [
          [60, 'секунд', 1],
          [120, '1 минута'],
          [3600, 'минут', 60],
          [7200, '1 час'],
          [86400, 'часов', 3600],
          [172800, '1 день'],
          [604800, 'дней', 86400],
          [1209600, '1 неделя'],
          [2419200, 'недель', 604800],
          [4838400, '1 мес'],
          [29030400, 'мес.', 2419200],
          [58060800, '1 год'],
          [2903040000, 'лет', 29030400]
        ];

        time = +new Date(time);

        let s = (+new Date() - time) / 1000;
        if(s === 0)
          return 'Только что';

        if(s < 0)
          s = Math.abs(s);

        for(const[limit, text, one] of formats)
          if(s < limit)
            return (one ? Math.floor(s / one) + ' ' + text : text) + " назад";

        return time;
      }
    },

    name(){
      return this.comment.user.first_name + " " + this.comment.user.last_name
    }
  },
  watch: {
    "comment.deleted": function(newVal, oldVal) {
      this.hidden = newVal
    },
  },
  data(){
    return {
      more: false,
      loading_like: false,
      loading_dislike: false,
      apiUrl: API_URL,

      hidden: this.comment.deleted,
      isNew: false,
      isEdit: false
    }
  },
  methods:{

    toggleMoreActions(){
      this.more = !this.more
    },
    cancelEdit(){
      this.isEdit = false
    },

    onSuccessEdit(data){
      this.isEdit = false
    },

    edit(){
      this.isEdit = true
    },

    findParent(){
      this.$emit("onCallParent", this.parent)
    },
    view(){
      this.isNew = false
    },
    async viewDeletedComment(b){

      this.isNew = false

      if(this.hidden === b) {
        return
      }

      this.hidden = b

    },
    async restoreComment(){

      try{
        await restoreComment(this.comment)
      }catch (e) {
        console.log(e)
      }

    },
    async deleteComment(){

      try{
        await deleteComment(this.comment)
      }catch (e) {
        console.log(e)
      }

    },

    async (){

      console.log("edit")
    },
    replyOpen: function(){
      this.$emit('onClickReply', this.comment)
    },

    like: async function(vote){

      if(!this.store.user){
        this.store.setViewSignin(true)
        return
      }

      if(vote > 0){
        this.loading_like = true
      }else{
        this.loading_dislike = true
      }

      try{
        await sendVote(this.comment, this.store.room, this.myReaction === vote ? 0 : (this.myReaction !== 0 ? 0 : vote))

      }catch (e) {
        console.log(e)
      }

      this.loading_like = false
      this.loading_dislike = false

    },

    onSendReply: async function(){

      // сообщим родителям о выполнении метода создания
      this.$emit('onReplyTo', {comment: this.comment})
    },

  }

}
</script>


<style scoped>
.comment-user {
  padding: 1.5rem 15px 15px 10px;
  border-left: 5px solid transparent;
  transition: all 0.5s;
}

.comment-user.comment-highlight {
  border-left: 5px solid #1367cd;
  background: #d9ebff;
}

.comment-user.comment-new {
  border-left: 5px solid #66cd13;
  background: #d9ffdb;
}


.comment-level .comment {
  /*padding-bottom: 0;*/
}
.comment-level {
  /*margin-top: 10px*/
}

.comment-reply-to-user{
  cursor: pointer;
}

.hatcomment-editmode {
  padding: 15px 0 0 0;
}
.hatcomment-editmode .hatcomment-form-wrapper {
  padding: 0;
}

.comment-img {
  position: relative;
}

</style>