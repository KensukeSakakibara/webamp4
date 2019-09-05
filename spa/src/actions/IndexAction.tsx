export const addTodoAction = (todo: string) => {
  return { 
    type: 'ADD_TODO',
    payload: { todo: todo }
  }
}
