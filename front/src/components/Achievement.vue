<template>

  <div class="achievement achievement-tooltip" :class="iconClass">
      <span class="emojiicon" v-if="isEmoji">{{user.icon}}</span>
      <span class="tooltiptext">{{ user.status }}</span>
  </div>
  
</template>

<script>

import {ICONS_URL} from "./../api/config"

export default {
  name: 'AchievementIcon',

  props: ["user"],
  computed: {
    isEmoji(){
      const s = this.user.icon.trim()
      return /\p{Emoji}/u.test(s)
    },
    iconClass(){

      const s = this.user.icon.trim()

      // если смайлик эмоджи
      if(this.isEmoji){
        return 'achievement-emoji'
      }

      return 'achievement-' + s
    }
  },
  data(){
    return {
      ICONS_URL: ICONS_URL
    }
  }
}
</script>

<style scoped>




.comment-img .achievement{
  position: absolute;
  right: -5px;
  top: -5px;
  width: 22px;
  height: 22px;
  background-repeat: no-repeat;
  background-size: auto;
}


.comment-img .achievement.achievement-slug {
  width: 36px;
  height: 36px;
  transform: rotate(25deg);

  background-image: url("./../../images/icons/slug.svg");
}
.comment-img .achievement.achievement-slug  {

  top: -20px;
  right: -15px;
  width: 36px;
  height: 36px;

}

.comment-img .achievement.achievement-brain  {
  top: -14px;
  right: -8px;
  width: 30px;
  height: 30px;

  background-image: url("./../../images/icons/brain.svg");
}

.comment-img .achievement.achievement-emoji {
  right: -8px;
  top: -5px
}

.comment-img .achievement.achievement-emoji span.emojiicon {
 line-height: 22px;
  font-size: 22px;
}

.achievement-tooltip {
  position: relative;
  display: inline-block;
}

.achievement-tooltip {
  position: relative;
  display: inline-block;
}

.achievement-tooltip .tooltiptext {
  visibility: hidden;
  background-color: black;
  color: #fff;
  text-align: center;
  border-radius: 6px;
  padding: 5px 10px;

  /* Position the tooltip */
  position: absolute;
  z-index: 1;
  top: -13px;
  left: 0;
  text-wrap: nowrap;
  margin-left: 35px;
}

.achievement-tooltip:hover .tooltiptext {
  visibility: visible;
}

.comment-img .achievement.achievement-emoji.achievement-tooltip .tooltiptext {
  top: -20px
}


@media (max-width: 480px) {

  .comment-img .achievement.achievement-brain {
    width: 20px;
    height: 20px;
    right: -5px;
    top: -10px;
    background-size: auto;
  }
  .comment-img .achievement.achievement-slug {
    width: 24px;
    height: 24px;
    right: -10px;
    top: -13px;
    background-size: auto;
  }

  .comment-img .achievement.achievement-emoji {
    width: 16px;
    height: 16px;
    right: -5px;
    top: -5px;
  }
  .comment-img .achievement.achievement-emoji span.emojiicon {
    line-height: 16px;
    font-size: 16px;
  }
}

</style>