
const storageKey = "HatCommentsUserData"



const SaveUser = (user) => {

    if(!user){
        localStorage.removeItem(storageKey)
        return
    }

    localStorage.setItem(storageKey, JSON.stringify(user))
}



const GetUser = async () => {
    const v = localStorage.getItem(storageKey)

    try{
        const data = JSON.parse(v)
        return data
    }catch (e) {
        console.log(e)
    }

    return null
}


export {SaveUser, GetUser}