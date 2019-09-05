import * as React from 'react';
import * as ReactDOM from 'react-dom';
import Header from "./components/common/Header";
import Footer from "./components/common/Footer";
import Content from "./components/index/Content";

const IndexPage: React.FC = () => {
  return (
    <div>
      <Header />
      <Content content="hello world" />
      <Footer />
    </div>
  )
}

ReactDOM.render(<IndexPage />, document.getElementById('index'));
