// @ts-check
import { defineStore } from 'pinia'


export const useStore = defineStore({
    id: 'store',
    state: () => ({
        signin: false,
        room: {
            id: 0,
            slug: "",
            name: "",
            project_id: 0,
        },
        selector: null,
        user: null,
        newcomments: []
    }),

    actions: {

        removeCommentFromInfo(c){

            this.$patch((state) => {

                const x = state.newcomments.indexOf(c)
                if(x !== -1){
                    state.newcomments.splice(x, 1)
                }

            })
        },


        addNewCommentToInfo(c){

            this.$patch((state) => {
                state.newcomments.push(c)
            })
        },

        setViewSignin(b){
            this.signin = b

            // document.getElementById(this.selector.replace("#", "")).scrollIntoView();

            if(this.selector){
                const y = document.getElementById(this.selector.replace("#", "")).getBoundingClientRect().top + window.scrollY;

                window.scroll({
                    top: y,
                    behavior: 'smooth'
                });
            }


        },

        setUser(u){
            this.user = u
        },

    },
})
