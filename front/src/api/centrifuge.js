import { Centrifuge } from 'centrifuge';
import {WS_URL} from "@/api/config";

let centrifuge = new Centrifuge(WS_URL);

class CentrifugeClass {

    centrifuge = null;
    connected = false

    connect(){
        centrifuge.connect()
    }

    disconnect(){
        centrifuge.disconnect()
    }

    init(withNewToken, connected){

        const o = {}
        if(withNewToken) {
            o.token = withNewToken
        }
        centrifuge = new Centrifuge(WS_URL, o);


        if(connected){
            centrifuge.on('connected', connected);
        }


        // centrifuge.on('disconnected', function(ctx) {
        //     console.log("[disconnected] ", ctx)
        // });
        //
        // centrifuge.on('connecting', function(ctx) {
        //     console.log("[connecting] ", ctx)
        // });
    }


    subscribe(channelName, f, fJoin, fLeave, fError){

        const sub = centrifuge.newSubscription(channelName, {
            joinLeave: true,
        });

        f && sub.on('publication', f);
        fJoin && sub.on("join", fJoin)
        fLeave && sub.on("leave", fLeave)
        fError && sub.on("error", fError)

        sub.subscribe();

        return sub;
    }

    publish(channelName, data) {
        return centrifuge.publish(channelName, data)
    }

}

const CentrifugeCl = new CentrifugeClass()

// Centrifuge.centrifuge = centrifuge

export {CentrifugeCl as Centrifuge}

export const deleteComment = async (comment) => {
    await centrifuge.publish("send:delete", {
        comment_id: comment.id,
    })
}

export const restoreComment = async (comment) => {
    await centrifuge.publish("send:restore", {
        comment_id: comment.id,
    })
}

export const sendComment = async ({comment, message, room}) => {

    if(comment){
        message.parent = comment.id
    }

    message.room_id = room.id
    message.project_id = room.project_id

    message.attachments = message.attachments.map(item => {
        return item.id
    })

    await centrifuge.publish("send:comment", message)
}


export const editComment = async ({message, comment}) => {

    message.attachments = message.attachments.map(item => {
        return item.id
    })

    message.comment_id = comment.id

    await centrifuge.publish("send:edit", message)
}

export const sendVote = async (comment, room, vote) => {


    await centrifuge.publish("send:vote", {
        comment_id: comment.id,
        vote: vote,
        room_id: room.id,
        project_id: room.project_id,
    })
}