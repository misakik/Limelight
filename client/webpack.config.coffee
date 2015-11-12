webpack = require 'webpack'
BowerWebpackPlugin = require 'bower-webpack-plugin'

module.exports =
  entry:
    renderer: './renderer.coffee'
  output:
    path: __dirname + './dist'
    filename: '[name].js'
  target: 'atom'
  resolve:
    root: __dirname
    extensions: ['', '.js', '.coffee', 'json', '.hamlc']
    modulesDirectories: [
      'node_modules'
      'bower_components'
    ]
  module: loaders: [
    { test: /\.coffee$/, loader: 'coffee' }
  ]
  plugins: [
    new BowerWebpackPlugin
      modulesDirectories: [__dirname + 'bower_components']
      manifestFiles: __dirname + 'bower.json'
  ]
