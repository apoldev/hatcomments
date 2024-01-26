<template>

  <div class="search-gif-box-wrapper">

    <div class="search-gif-input-wrapper">
      <input ref="gifsearch" type="text" v-model="search" @keyup="onSearch()" placeholder="Начните писать для поиска gif..."/>
    </div>

    <div class="gif-box-wrapper" v-if="loading || gifs.length > 0">

      <div class="loading" v-if="loading">
        <LoadingSvg  fill="#222"/>
      </div>


      <div class="search-gif-box" v-else>
        <div class="search-gif-box-el" v-for="v in gifs" >
          <img :src="v.nanogif" :alt="v.title" :data-gif="v.gif" @click="addAttachment(v)"/>
        </div>
      </div>

    </div>



  </div>

</template>


<script>

import axios from 'axios'
import LoadingSvg from "@/components/Loading";

export default {
  name: 'GifBox',
  components: {LoadingSvg},
  data(){
    return {
      search: '',
      gifs: [],
      loading: false
    }
  },
  mounted() {

    setTimeout(() => {
      this.$refs.gifsearch.focus()
    }, 100)

  },
  methods: {

    addAttachment(gif){
      this.$emit('onAddGif', gif)
    },


    async onSearch(){


      this.gifs = [];

      if(!this.search){
        return
      }

      this.loading = true

      try {
        const url = "https://api.tenor.com/v1/search?key=43TZW7V8IHV3&q=" + encodeURIComponent(this.search)+ "&limit=50&media_filter=basic&_=1687114618031"

        const resp = await axios.get(url)

        if(resp.data?.results){
          this.gifs = resp.data.results.map(item => {
            return {
              id: item.id,
              title: item.title || item.content_description,
              nanogif: item.media[0].nanogif.url,
              gif: item.media[0].gif.url
            }
          })
        }

      }catch (e) {

      }finally {
        this.loading = false
      }




    }

  }
}
</script>

<style scoped>

.search-gif-box-wrapper{
  padding-top: 10px;
  margin-top: 10px;
  border-top: 1px solid rgba(0, 0, 0, 0.1);
}

.search-gif-box > * {

}
.search-gif-box .search-gif-box-el{
  display: inline-block;
}
.search-gif-box .search-gif-box-el > * {
  height: 80px;
  margin: 0 15px 15px 0;
}

.search-gif-box-wrapper .loading > svg {
  width: 30px;
  height: 30px;
  margin: 15px 0;
}


.search-gif-input-wrapper {

}

.search-gif-input-wrapper {
  display: flex;
  align-items: center;

  box-shadow: 0 0 5px 1px rgba(0, 0, 0, .0975);
  padding: 8px 10px;
  border-radius: 6px;
  transition: all 0.3s;
  border: 1px solid rgba(0, 0, 0, 0.1)

}
.search-gif-input-wrapper input {
  border: 0;
  outline: none;
  font-size: 15px;
  width: 100%;
}

.search-gif-box {
  max-height: 11rem;
  overflow: auto;

}

.gif-box-wrapper{
  margin: 10px 0;
  display: flex;
  align-content: center;
  justify-content: center;
}

.search-gif-box {
  display: flex;
  /*align-content: center;*/
  /*justify-content: center;*/
  flex-wrap: wrap;
}

.search-gif-box-el img{
  cursor: pointer;
}
</style>