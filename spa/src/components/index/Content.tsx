import * as React from 'react';

export interface Props {
  content: string;
}

export default class MyComponent extends React.Component<Props, {}> {
  render() {
    return (
      <div>
        <article className="main-content">
          <section className="top">
            <h1>{this.props.content}</h1>
            <p className="logout"><a href="/index/logout">Logout</a></p>
          </section>
        </article>
      </div>
    )
  }
}
