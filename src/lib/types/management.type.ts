import { type User } from "./user.type";

export const PERMISSION_UI_MANAGEMENT = "ui.management";
export const PERMISSION_USERS_MANAGEMENT = "management.users";
export const PERMISSION_ROLES_MANAGEMENT = "management.roles";

export type ManagementUsers = {
	users: User[]
	total: number
	roles: string[]
	currenncies: string[]
}
