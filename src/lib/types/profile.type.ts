export type DisplayComment = {
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
