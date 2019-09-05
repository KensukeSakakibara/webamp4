import * as React from 'react'
import Header from "./common/Header"
import Footer from "./common/Footer"
import Content from "./index/Content"

export const IndexPage: React.FC = () => {
  return (
    <div>
      <Header />
      <Content content="hello world" />
      <Footer />
    </div>
  )
}
