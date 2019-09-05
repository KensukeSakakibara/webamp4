import * as React from 'react'
import * as ReactDOM from 'react-dom'
import { createStore } from 'redux';
import { Provider } from 'react-redux'
import { indexReducer } from "./reducers/IndexReducer";
import IndexComponent from './containers/IndexContainer';

const store = createStore(indexReducer)

ReactDOM.render(
  <Provider store={store}>
    <IndexComponent />
  </Provider>,
  document.getElementById('index')
)
