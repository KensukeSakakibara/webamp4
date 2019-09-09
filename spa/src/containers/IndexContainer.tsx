import { Dispatch } from 'redux';
import { connect } from 'react-redux'
import { Action } from 'typescript-fsa';
import * as indexAction from '../actions/IndexAction'
import IndexComponent from '../components/IndexComponent'

// ステートの定義
interface TaskInterface {
  id: number;
  text: string;
  done: boolean;
}

export interface IndexStateInterface {
  tasks: TaskInterface[];
}

const mapStateToProps = (state: IndexStateInterface) => {
  return {
    tasks: state.tasks,
  }
}

// アクションの定義
interface IndexActionInterface {
  addTodoAction: (v: string) => Action<string>;
}

const mapDispatchToProps = (dispatch: Dispatch<Action<string>>) => {
  return {
    addTodoAction: (todo: string) => dispatch(indexAction.addTodoAction(todo)),
  }
}

// コネクト
export default connect(mapStateToProps, mapDispatchToProps)(IndexComponent)
export type IndexProps = IndexStateInterface & IndexActionInterface
