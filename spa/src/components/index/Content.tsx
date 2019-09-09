import * as React from 'react'
import { IndexProps } from '../../containers/IndexContainer'

class IndexContent extends React.Component<IndexProps> {
  render() {
    return (
      <article className="main-content">
        <section className="top">
          <h1></h1>
          <p className="logout"><a href="/index/logout">Logout</a></p>
          <button type="button" onClick={(e) => this.props.addTodoAction("aiueo")}>テスト</button>
        </section>
      </article>
    )
  }
}

export default IndexContent
