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
          app: APP_DIR + '/index.tsx',
      },
      context: APP_DIR,
      output: {
          path: path.resolve(__dirname, 'dist'),
          filename: '[name].bundle.js',
      },

      // Add '.ts' and '.tsx' as resolvable extensions.
      resolve: {
          extensions: [".ts", ".tsx", ".js", ".json"]
      },
      module: {
          rules: [
              // Load TypeScript files
              { test: /\.tsx?$/, loader: "awesome-typescript-loader" },
              { enforce: "pre", test: /\.js$/, loader: "source-map-loader" }
          ],
          loaders : [
              {
                  test : /\.jsx?/,
                  include : APP_DIR,
                  loader : 'babel-loader'
              }, {
                  test : /\.tsx?/,
                  include : APP_DIR,
                  loader : 'awesome-typescript-loader'
              }
          ]
      },
      plugins: [new HtmlWebpackPlugin(htmlIndexConfig)],
      devtool: "source-map",
};

module.exports = config;
