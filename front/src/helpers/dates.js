

export const getIntervalUpdateComment = (created_at) => {

    let timeout = 1
    let time = +new Date(created_at);
    let s = (+new Date() - time) / 1000;

    if(s < 60) {
        timeout = 1
    }else if (s < 3600) {
        timeout = 60
    } else {
        timeout = 60*60
    }

    return timeout

}