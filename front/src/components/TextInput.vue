<template>
  <div class="hatcomment-form-wrapper">
    <div class="form">
      <div>
        <textarea
            ref="input"
            placeholder="Написать комментарий..."
            v-model="input"
            v-on:keydown="send(false, $event)"
            @paste="onPaste"
        ></textarea>
      </div>

      <div class="form-buttons">

        <input ref="uploader" type="file" accept="image/*" @change="uploadImage($event)" style="display:none">

        <div class="btn-actions">
          <div class="btn-action btn-action-emoji" @click="emojiBoxToggle()">
            <svg fill="#000000" width="800px" height="800px" viewBox="-8 0 512 512" xmlns="http://www.w3.org/2000/svg"><path d="M248 8C111 8 0 119 0 256s111 248 248 248 248-111 248-248S385 8 248 8zm0 448c-110.3 0-200-89.7-200-200S137.7 56 248 56s200 89.7 200 200-89.7 200-200 200zm117.8-146.4c-10.2-8.5-25.3-7.1-33.8 3.1-20.8 25-51.5 39.4-84 39.4s-63.2-14.3-84-39.4c-8.5-10.2-23.7-11.5-33.8-3.1-10.2 8.5-11.5 23.6-3.1 33.8 30 36 74.1 56.6 120.9 56.6s90.9-20.6 120.9-56.6c8.5-10.2 7.1-25.3-3.1-33.8zM168 240c17.7 0 32-14.3 32-32s-14.3-32-32-32-32 14.3-32 32 14.3 32 32 32zm160-60c-25.7 0-55.9 16.9-59.9 42.1-1.7 11.2 11.5 18.2 19.8 10.8l9.5-8.5c14.8-13.2 46.2-13.2 61 0l9.5 8.5c8.5 7.4 21.6.3 19.8-10.8-3.8-25.2-34-42.1-59.7-42.1z"/></svg>
          </div>


          <div class="btn-action btn-action-image" @click="handleSelectImage()">

            <svg viewBox="0 -2 20 20" version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink">
              <g id="Page-1" stroke="none" stroke-width="1" fill="none" fill-rule="evenodd">
                <g id="Dribbble-Light-Preview" transform="translate(-340.000000, -3881.000000)" fill="#000000">
                  <g id="icons" transform="translate(56.000000, 160.000000)">
                    <path d="M296,3725.5 C296,3724.948 296.448,3724.5 297,3724.5 C297.552,3724.5 298,3724.948 298,3725.5 C298,3726.052 297.552,3726.5 297,3726.5 C296.448,3726.5 296,3726.052 296,3725.5 L296,3725.5 Z M296.75,3728.75 L300,3733 L288,3733 L292.518,3726.812 L295.354,3730.625 L296.75,3728.75 Z M302,3734 C302,3734.552 301.552,3735 301,3735 L287,3735 C286.448,3735 286,3734.552 286,3734 L286,3724 C286,3723.448 286.448,3723 287,3723 L301,3723 C301.552,3723 302,3723.448 302,3724 L302,3734 Z M302,3721 L286,3721 C284.896,3721 284,3721.895 284,3723 L284,3735 C284,3736.104 284.896,3737 286,3737 L302,3737 C303.105,3737 304,3736.104 304,3735 L304,3723 C304,3721.895 303.105,3721 302,3721 L302,3721 Z" id="image_picture-[#973]">

                    </path>
                  </g>
                </g>
              </g>
            </svg>
          </div>

          <div class="btn-action btn-action-gif" @click="openGifBox()">
            <svg viewBox="0 0 24 24" version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink">
              <g  stroke="none" stroke-width="1" fill="none" fill-rule="evenodd">
                <g id="ic_fluent_gif_24_filled" fill="#212121" fill-rule="nonzero">
                  <path d="M18.75,3.50054297 C20.5449254,3.50054297 22,4.95561754 22,6.75054297 L22,17.2531195 C22,19.048045 20.5449254,20.5031195 18.75,20.5031195 L5.25,20.5031195 C3.45507456,20.5031195 2,19.048045 2,17.2531195 L2,6.75054297 C2,4.95561754 3.45507456,3.50054297 5.25,3.50054297 L18.75,3.50054297 Z M8.01459972,8.87193666 C6.38839145,8.87193666 5.26103525,10.2816525 5.26103525,11.9943017 C5.26103525,13.707564 6.38857781,15.1202789 8.01459972,15.1202789 C8.90237918,15.1202789 9.71768065,14.6931811 10.1262731,13.9063503 L10.2024697,13.7442077 L10.226,13.674543 L10.2440163,13.5999276 L10.2440163,13.5999276 L10.2516169,13.5169334 L10.2518215,11.9961937 L10.2450448,11.9038358 C10.2053646,11.6359388 9.99569349,11.4234501 9.72919932,11.3795378 L9.62682145,11.3711937 L8.62521827,11.3711937 L8.53286035,11.3779703 C8.26496328,11.4176506 8.05247466,11.6273217 8.00856234,11.8938159 L8.00021827,11.9961937 L8.00699487,12.0885517 C8.0466751,12.3564487 8.25634623,12.5689373 8.5228404,12.6128497 L8.62521827,12.6211937 L9.00103525,12.6209367 L9.00103525,13.3549367 L8.99484486,13.3695045 C8.80607251,13.6904125 8.44322427,13.8702789 8.01459972,13.8702789 C7.14873038,13.8702789 6.51103525,13.0713011 6.51103525,11.9943017 C6.51103525,10.9182985 7.14788947,10.1219367 8.01459972,10.1219367 C8.43601415,10.1219367 8.67582824,10.1681491 8.97565738,10.3121334 C9.28681641,10.4615586 9.6601937,10.3304474 9.80961888,10.0192884 C9.95904407,9.70812933 9.82793289,9.33475204 9.51677386,9.18532686 C9.03352891,8.95326234 8.61149825,8.87193666 8.01459972,8.87193666 Z M12.6289445,8.99393497 C12.3151463,8.99393497 12.0553614,9.22519285 12.0107211,9.52657705 L12.0039445,9.61893497 L12.0039445,14.381065 L12.0107211,14.4734229 C12.0553614,14.7748072 12.3151463,15.006065 12.6289445,15.006065 C12.9427427,15.006065 13.2025276,14.7748072 13.2471679,14.4734229 L13.2539445,14.381065 L13.2539445,9.61893497 L13.2471679,9.52657705 C13.2025276,9.22519285 12.9427427,8.99393497 12.6289445,8.99393497 Z M17.6221579,9.00083497 L15.6247564,8.99393111 C15.3109601,8.99285493 15.0503782,9.22321481 15.0046948,9.52444312 L14.9975984,9.61677709 L14.9975984,14.3649711 L15.0043751,14.4573291 C15.0440553,14.7252261 15.2537265,14.9377148 15.5202206,14.9816271 L15.6225985,14.9899711 L15.7149564,14.9831945 C15.9828535,14.9435143 16.1953421,14.7338432 16.2392544,14.467349 L16.2475985,14.3649711 L16.2470353,13.2499367 L17.37,13.2504012 L17.4623579,13.2436246 C17.730255,13.2039444 17.9427436,12.9942732 17.9866559,12.7277791 L17.995,12.6254012 L17.9882234,12.5330433 C17.9485432,12.2651462 17.738872,12.0526576 17.4723779,12.0087453 L17.37,12.0004012 L16.2470353,11.9999367 L16.2470353,10.2449367 L17.6178421,10.2508313 L17.7102229,10.2443727 C18.0117595,10.2007704 18.2439132,9.94178541 18.2450039,9.62798912 C18.24608,9.31419284 18.0157202,9.05361096 17.7144919,9.00793041 L17.6221579,9.00083497 L15.6247564,8.99393111 L17.6221579,9.00083497 Z" id="🎨-Color">

                  </path>
                </g>
              </g>
            </svg>
          </div>
        </div>

        <div class="btn-actions btn-action-right">

          <a href="#" @click.prevent="cancelEdit" class="form-button form-button-cancel" v-if="editmode">
            <span class="label">Отмена</span>
          </a>

          <a href="#" class="form-button" :class='{"disabled": loading}' @click.prevent="send(true)">
          <span class="icon">

            <Loading fill="#ddd" v-if="loading" />

            <svg v-if="!loading" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
              <path d="M7.39999 6.32003L15.89 3.49003C19.7 2.22003 21.77 4.30003 20.51 8.11003L17.68 16.6C15.78 22.31 12.66 22.31 10.76 16.6L9.91999 14.08L7.39999 13.24C1.68999 11.34 1.68999 8.23003 7.39999 6.32003Z" stroke="#fff" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/>
              <path d="M10.11 13.6501L13.69 10.0601"  stroke="#fff"  stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/>
            </svg>



          </span>
            <span class="label" v-if="!editmode">Отправить</span>
            <span class="label" v-if="editmode">Сохранить</span>
          </a>


        </div>
      </div>


      <GifBox v-if="gifBox" @onAddGif="onAddGif"/>

      <div class="emoji-box" v-if="emojiBox">
        <emoji-picker :show_arrow="false" @emojiClick="emojiClick"/>
      </div>


      <div class="form-attachments" v-if="attachments.length > 0">

        <div class="form-attachments-data">
          <span class="form-attachments-header">Вложенные файлы</span>

          <AttachmentList :attachments="attachments" :edited="true" @onDelAttach="delAttach($event)"/>
        </div>


      </div>




    </div>


  </div>
</template>

<script>
import {$api} from "@/api/api";
import GifBox from "@/components/GifBox";
import AttachmentList from "@/components/AttachmentList";
import EmojiPicker from "@/components/emoji/EmojiPicker";
import Loading from "@/components/Loading";
import {sendComment, editComment} from "@/api/centrifuge";
import {useStore} from "@/store/store";


export default {
  name: 'TextInput',
  components: {AttachmentList, EmojiPicker, Loading, GifBox},
  props: ["parent", "text", "attach", "editmode", "comment"],
  setup(){

    return {
      store: useStore()
    }
  },
  mounted() {

    // if(!this.editmode){

    if(this.parent){
      this.$refs["input"].focus()
    }

    // }


  },
  data: function (){


    let a = []

    if(this.attach && this.attach?.length){
      a = [...this.attach]
    }

    return {
      input: this.text || '',
      attachments: a,
      emojiBox: false,
      loading: false,
      gifBox: false,
    }
  },
  watch: {
    "input": function(){
      const textarea = this.$refs.input
      textarea.style.height = (textarea.scrollHeight) + "px";
    }
  },
  computed: {
    message: function(){

      return {
        input: this.input,
        attachments: this.attachments,

      }
    }
  },
  methods: {

    cancelEdit(){
      this.$emit("cancelEdit")
    },

    initLoadingAttachment(){
      this.attachments.push({type: 'loading'})
    },
    endLoadingAttachment(){
      this.attachments = this.attachments.filter(item => {
        if(item.type === 'loading'){
          return false;
        }

        return true;
      })
    },
    async onAddGif(gif){

      this.gifBox = false
      this.initLoadingAttachment()

      try {
        const resp = await $api.post("/attachments/tenor", gif);


        this.endLoadingAttachment()
        if(resp.data){
          this.attachments.push(resp.data)
        }

      }catch (e) {
        console.log(e)
      }finally {
        this.endLoadingAttachment()
      }

    },

    openGifBox(){
      this.gifBox = !this.gifBox;
      this.emojiBox = false
    },

    onPaste(e){

      const items = e.clipboardData.items

      for (let i = 0; i < items.length; i++) {

        if (items[i].type.indexOf("image") == -1) continue;
        var blob = items[i].getAsFile();

        this.uploadImage(blob)
      }

    },
    emojiBoxToggle(){
      this.emojiBox = !this.emojiBox
      this.gifBox = false
    },
    emojiClick(emoji){

    this.input += emoji + ""
      this.$refs.input.focus()

    },
    handleSelectImage(){

      this.emojiBox = false
      this.gifBox = false

      this.$refs.uploader.click();

    },
    delAttach(i){

      this.attachments.splice(i, 1)

    },
    send(force, e){

      if(this.loading){
        return
      }

      if(force || ((e.metaKey || e.ctrlKey || e.altKey) && e.keyCode === 13)){
        this.onSend()
      }

    },

    async onSend(){

      if(!this.store.user){
        this.store.setViewSignin(true)
        return
      }

      if(this.input === '' && this.attachments.length === 0){
        alert("Введите сообщение")
        return
      }

      this.loading = true

      try{

        if(this.editmode){

          await editComment({message: this.message, comment: this.comment})
          this.$emit("onSuccessEdit", this.message)

        }else{

          await sendComment({room: this.store.room, comment: this.parent, message: this.message})
          this.$emit('onSend')

          this.input = ''
          this.attachments = []

        }


      }catch (e){
        console.log(e)
      }finally {
        this.loading = false
      }

    },


    imageForForm(image){

      let data = new FormData();
      data.append('name', image.name);
      data.append('file', image);

      return data;
    },


    async uploadImage(event) {

      this.initLoadingAttachment()

      const response = await $api.get("/upload/url");
      const URL = response.data.url;

      let image = event

      if(event?.target?.files){
        image = event.target.files[0]
      }

      let data = this.imageForForm(image)

      let config = {
        header : {
          'Content-Type' : 'image/png'
        }
      }

      const resp = await $api.put(
          URL,
          data,
          config
      );

      if(resp.data){
        this.attachments.push(resp.data)
      }
      this.endLoadingAttachment()

      event.target.value = ''
    }
  }
}
</script>

<style scoped>
.btn-actions .btn-action{
  align-items: center;
  justify-content: center;
  display: flex;
}

.btn-actions .btn-action.btn-action-emoji svg{
  width: 23px;
  height: 23px;
}
.btn-actions .btn-action.btn-action-image svg{
  width: 25px;
  height: 25px;
}


.btn-actions .btn-action.btn-action-gif svg{
  height: 30px;
  width: 30px;
}

.hatcomments .hatcomment-form-wrapper {
  padding: 0 15px;
  flex: 1;
}

</style>