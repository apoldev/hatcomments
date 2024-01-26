import {$api} from "@/api/api";


const PostComment = async (project_id, room_id, message, parent = undefined) => {

    const {input, attachments} = message


    // Если есть родитель - тогда его айди
    let reply_to = undefined
    if(parent) {
        reply_to = parent.id
    }

    const data = {
        room_id,
        text: input,
        reply_to
    }

    if(attachments){
        data['attachments'] = attachments.map(item => {
            return item.id
        });
    }

    console.log(data)

    let response = await $api.post("/project/" +project_id+ "/comment", data)

    return response.data

}

export {PostComment}