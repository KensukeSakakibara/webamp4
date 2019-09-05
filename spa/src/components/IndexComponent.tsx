import * as React from 'react'
import Header from "./common/Header"
import Footer from "./common/Footer"
import Content from "./index/Content"

const IndexComponent: React.FC = () => {
  return (
    <div>
      <Header />
      <Content content="hello world" />
      <Footer />
    </div>
  )
}

export default IndexComponent
