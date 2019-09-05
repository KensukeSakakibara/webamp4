const path = require('path');

module.exports = {
  mode: 'development',
  entry: './spa/src/Index.tsx',
  output: {
    path: path.join(__dirname, './static/js'),
    filename: 'index.js',
  },
  devtool: 'source-map',
  resolve: {
    extensions: ['.ts', '.tsx', '.js']
  },
  module: {
    rules: [
      {
        test: /\.tsx$/,
        use: 'ts-loader'
      }
    ]
  }
};
