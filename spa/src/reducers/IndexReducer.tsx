import { reducerWithInitialState } from 'typescript-fsa-reducers'
import * as indexActions from '../actions/IndexAction'
import * as indexContainer from '../containers/IndexContainer'

// Stateの初期値
const initialTasks: indexContainer.IndexStateInterface = {
  tasks: [{
    id: 1,
    text: 'initial task',
    done: false,
  }],
};

let idCounter: number = 1;

// Reducerの処理
export const indexReducer = reducerWithInitialState(initialTasks)
  .case(indexActions.addTodoAction, (state: indexContainer.IndexStateInterface, payload: string) => {
    const newState = Object.assign({}, state)
    newState.tasks.push({
      id: ++idCounter,
      text: payload,
      done: false
    })
    return newState
  })
