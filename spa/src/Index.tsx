import * as React from 'react'
import * as ReactDOM from 'react-dom'
import { createStore, applyMiddleware } from 'redux'
import { Provider } from 'react-redux'
import createSagaMiddleware from "redux-saga"
import { indexReducer } from "./reducers/IndexReducer"
import indexSaga from "./sagas/IndexSaga"
import IndexComponent from './containers/IndexContainer'

// create the saga middleware
const sagaMiddleware = createSagaMiddleware()

// ストアの作成
const store = createStore(
  indexReducer,
  applyMiddleware(sagaMiddleware)
)

// Sagaを走らせる
sagaMiddleware.run(indexSaga)

ReactDOM.render(
  <Provider store={store}>
    <IndexComponent />
  </Provider>,
  document.getElementById('index')
)
