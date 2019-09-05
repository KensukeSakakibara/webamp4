import { connect } from 'react-redux';
import * as actions from '../actions/IndexAction';
import IndexComponent from '../components/IndexComponent';

const mapStateToProps = (state: any) => {
  return {
    todo: state.todo,
  }
}

const mapDispatchToProps = (dispatch: any) => {
  return {
    addTodo: (todo: string) => dispatch(actions.addTodoAction(todo)),
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(IndexComponent)
