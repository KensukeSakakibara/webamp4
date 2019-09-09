import * as React from 'react'
import { IndexProps } from '../containers/IndexContainer'
import CommonHeader from "./common/Header"
import CommonFooter from "./common/Footer"
import IndexContent from "./index/Content"

class IndexComponent extends React.Component<IndexProps> {
  constructor(props: IndexProps) {
    super(props);
  }

  render() {
    return (
      <div>
        <CommonHeader />
        <IndexContent {...this.props} />
        <CommonFooter />
        <button type="button" onClick={(e) => this.props.fetchUserAction(1)}>Saga Test</button>
      </div>
    )
  }
}

export default IndexComponent
