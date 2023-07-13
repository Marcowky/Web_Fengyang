import { axios } from '../main.js'
import { userStore } from '../main.js'
import config from '../config/config.json'

const renameData = (data) => {
    // 替换每个项的名称
    const replacedData = data.map(item => {
        return {
            index: item.Order.toString(),
            label: item.Label,
            mainMenu: '/' + item.PageName
        };
    });
    return replacedData
}

export function menuList() {
    return userStore.menuItem
}

export async function categoryStore() {
    try {
        let res = await axios.get(`category/categoryList`)
        userStore.menuItem = renameData(res.data.data.categorys).concat(config.menuItems)
    } catch (error) {
    }
}

export function categoryGetByArticleType(type) {
    return userStore.menuItem.filter(item => item.mainMenu == type)
}

export function categoryGetLabel(type, category) {
    return userStore.menuItem.find(item => item.mainMenu == type && item.index == category).label
}
