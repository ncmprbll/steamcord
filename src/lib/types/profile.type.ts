import type { Product } from "./product.type"

export type DisplayComment = {
    id: number
    commentator: string
	avatar: string
	display_name: string
	text: string
	created_at: string
}

export type ProfileComments = {
	comments: DisplayComment[]
	total: number
}

export type FriendStatus = {
    isFriend: boolean
    hasIncomingInvite: boolean
    hasOutgoingInvite: boolean
}

export type Games = {
    games: Product[]
    total: number
}