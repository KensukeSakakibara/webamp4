import { actionCreatorFactory } from 'typescript-fsa';
const actionCreator = actionCreatorFactory();

export const addTodoAction = actionCreator<string>('ADD_TODO');

