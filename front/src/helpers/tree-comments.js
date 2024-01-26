
const findChildren = (parent, comments) => {

    let data = comments.filter(item => {
        return item.reply_to == parent.id
    });


    // order by date ASC
    data.sort(function(a,b){
        return new Date(a.created_at) - new Date(b.created_at);
    });

    for(let i = 0; i < data.length; i++ ){
        data[i].children = findChildren(data[i], comments)
    }

    return data;

}


// TreeComments создает дерево из массива комментариев, полученных с сервера
const TreeComments = (comments) => {

    const zeroLevel = comments.filter(item => {
        return item.level < 1 || !item.level;
    })


    // Для каждого коммента с нулевым уровнем запустим сборку детей
    for(let i = 0; i < zeroLevel.length; i++ ){
        zeroLevel[i].children = findChildren(zeroLevel[i], comments)
    }

    return zeroLevel;
}


function findId(id, list) {
    for (const _item of list) {
        if (_item.id === id) {
            // Return some top-level id
            return _item;

            // If the item is not the one, check if it has subs
        } else if (_item?.children) {

            // Use _item.subs as a list, calling the function recursively
            const subId = findId(id, _item.children);

            // Return some nested subId if found
            if (subId !== undefined) return subId;
        }
    }
    // Return undefined if not found
    return undefined;
}

export {TreeComments, findId}