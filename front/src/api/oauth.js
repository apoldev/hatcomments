
const openWindow = (url, callback, listener) => {

    let myWindow = undefined


    function lll(e){

        var payload = listener(e)

        if(payload){
            callback(payload)
            myWindow && myWindow.close()

            console.log("removeEventListener")
            window.removeEventListener("message", lll)
        }
    }

    window.addEventListener("message", lll);

    const w = 600
    const h = 400
    const left = (screen.width/2)-(w/2);
    const top = (screen.height/2)-(h/2);

    myWindow = window.open(url, "", "width=600,height=400,left=" + left + ",top=" + top)

    const timer = setInterval(checkChild, 200);
    function checkChild() {
        if (myWindow.closed) {
            clearInterval(timer);
            callback(null)
            console.log("removeEventListener")
            window.removeEventListener("message", lll)
        }
    }

}

export const VKAuth = (callback) => {


    const redirect = "https://comment.1trackapp.com/callback"
    const client_id = "51700753"

    let link = "https://oauth.vk.com/authorize?client_id="+client_id+"&display=popup&scope=offline,email&response_type=token&v=5.131&redirect_uri=" + encodeURIComponent(redirect)


    openWindow(link, callback, (e) => {

        if(e.data["access_token"]){
            return {
                vk_token: e.data["access_token"],
                email: e.data["email"]
            }
        }

        return null
    })

}


export const GoogleAuth = (callback) => {

    const clientID = "945570684383-1g1vvufaq8go4h5usmhbgtm8ht4721cq.apps.googleusercontent.com"
    const redirect = "https://comment.1trackapp.com/callback"

    const url = "https://accounts.google.com/o/oauth2/v2/auth" +
        "?client_id=" +encodeURIComponent(clientID)+ "&redirect_uri=" +encodeURIComponent(redirect)+ "&scope=openid%20profile%20email&response_type=token"


    openWindow(url, callback, (e) => {
        return e.data["access_token"]
    })
}