import { createApp } from 'vue'
import App from './app/App.vue'
import {Centrifuge} from "@/api/centrifuge";
import { createPinia } from 'pinia'

window.hatComments = ({project, room, selector, actions}) => {

    const app = createApp(App)

    app.use(createPinia())
    app.config.globalProperties.$config = {
        project: project,
        room: room,
        selector: selector,
    }

    app.config.globalProperties.$centrifuge = Centrifuge

    app.config.globalProperties.$actions = actions

    app.mount(selector)
}


window.hatCommentConfig && window.hatComments(window.hatCommentConfig)




