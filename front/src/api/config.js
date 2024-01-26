export const API_URL = process.env.NODE_ENV === 'development' ? '' : ''
export const WS_URL = process.env.NODE_ENV === 'development' ? 'ws://localhost:8001/connection/websocket' : 'ws://localhost:8001/connection/websocket'

export const ICONS_URL = API_URL + "/images/icons/"