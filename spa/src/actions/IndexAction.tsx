import { actionCreatorFactory } from 'typescript-fsa'
const actionCreator = actionCreatorFactory()

// 定数の定義
export const ADD_TODO = "ADD_TODO"
export const SAGA_FETCH_USER = "SAGA_FETCH_USER"

// アクションの発行
export const addTodoAction = actionCreator<string>(ADD_TODO)
export const fetchUserAction = actionCreator<number>(SAGA_FETCH_USER)
