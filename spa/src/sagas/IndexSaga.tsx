import { takeEvery } from "redux-saga/effects"
import * as indexActions from '../actions/IndexAction'

function* fetchUser(action: ReturnType<typeof indexActions.fetchUserAction>) {
  let text = action.payload
  yield console.log(text)
}

function* indexSaga() {
  yield takeEvery(indexActions.SAGA_FETCH_USER, fetchUser)
}

export default indexSaga
