// name: 登录接口
// author: Ethan.Wang
// desc:

import request from "./index";
import {PromiseRes} from './api'

// 获取登录链接 - 返回结构体
export interface UserInfoRes {
    uid: string
    nickname: string
    avatar: string
}

// 获取用户基本信息
export const UserBaseInfo = (): PromiseRes<UserInfoRes> => request({
    method: "post",
    data: {},
    url: "/user/info/base",
})

// ===================================

// 获取用户列表 - 请求
export interface UserListReq {
    limit: number
    page: number
}

// 获取用户列表 - 返回
export interface UserListRes {
    list: Array<UserListData>
    total: number
}

export interface UserListData {
    uid: string
    avatar: string
    nickname: string
    level: string
    location: string
    create_at: string
    updated_at: string
}

// 获取用户列表
export const UserList = (data: UserListReq): PromiseRes<UserListRes> => request({
    method: "get",
    params: data,
    url: "/user/list",
})

// ===================================

// 复合条件查询用户信息 - 请求
export interface UserQueryReq {
    uid: string
    nickname: string
    level: string
    limit: number
    page: number
}

// 复合条件查询用户信息
export const UserQuery = (data: UserQueryReq): PromiseRes<UserListRes> => request({
    method: "get",
    params: data,
    url: "/user/query",
})

// ===================================

// 更新用户权限 - 请求
export interface UpdateUserLevelReq {
    uid: string
    level: string
}

// 更新用户权限
export const UpdateUserLevel = (data: UpdateUserLevelReq): PromiseRes<null> => request({
    method: "post",
    data,
    url: "/user/update/level",
})