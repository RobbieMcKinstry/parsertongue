const path = require('path');
const webpack = require('webpack');
const HtmlWebpackPlugin = require('html-webpack-plugin');

const APP_DIR = path.resolve(__dirname, 'src');
const htmlIndexConfig = {
    'title': 'Parsertongue',
    'filename': 'index.html',
    'inject': 'body',
    'template': 'index.ejs'
};

const config = {
      entry: {
          app: APP_DIR + '/index.js',
      },
      context: APP_DIR,
      output: {
          path: path.resolve(__dirname, 'dist'),
          filename: '[name].bundle.js',
      },
      module: {
          loaders : [
              {
                  test : /\.jsx?/,
                  include : APP_DIR,
                  loader : 'babel-loader'
              }
          ]
      },
      plugins: [new HtmlWebpackPlugin(htmlIndexConfig)]
};

module.exports = config;
