
interface initialState {
  todoList: string[]
}

export const indexReducer = (state: initialState, action: any) => {
  switch (action.type) {
    case 'ADD_TODO':
      const todo: string = action.payload.todo;
      const newState = Object.assign({}, state);
      newState.todoList.push(todo);
      return newState;

    default:
      return state;
  }
}
