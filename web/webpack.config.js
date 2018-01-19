const path = require('path');
const webpack = require('webpack');

const APP_DIR = path.resolve(__dirname, 'src');

var config = {
      entry: {
          app: APP_DIR + '/index.js',
      },
      context: APP_DIR,
      output: {
          path: path.resolve(__dirname, 'dist'),
          filename: '[name].bundle.js',
      },
      module : {
          loaders : [
              {
                  test : /\.jsx?/,
                  include : APP_DIR,
                  loader : 'babel-loader'
              }
          ]
      }
};

module.exports = config;
