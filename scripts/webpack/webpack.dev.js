'use strict';

const merge = require('webpack-merge');
const common = require('./webpack.common.js');
const path = require('path');
const webpack = require('webpack');
const HtmlWebpackPlugin = require("html-webpack-plugin");
const ExtractTextPlugin = require("extract-text-webpack-plugin");
const WebpackCleanupPlugin = require('webpack-cleanup-plugin');
const BundleAnalyzerPlugin = require('webpack-bundle-analyzer').BundleAnalyzerPlugin;
const pkg = require('../../package.json');
const _ = require('lodash');

let dependencies = Object.keys(pkg.dependencies);
// remove jquery
dependencies = _.filter(dependencies, function(key) {
  return key !== 'jquery';
});
// add it first
dependencies.unshift('jquery');

module.exports = merge(common, {
  devtool: "source-map",

  entry: {
    dark: './public/sass/grafana.dark.scss',
    light: './public/sass/grafana.light.scss',
    vendor: dependencies,
  },

  module: {
    rules: [
      require('./sass.rule.js')({
        sourceMap: true, minimize: false
      })
    ]
  },

  plugins: [
    new ExtractTextPlugin({ // define where to save the file
      filename: 'grafana.[name].css',
    }),
    new HtmlWebpackPlugin({
      filename: path.resolve(__dirname, '../../public/views/index.html'),
      template: path.resolve(__dirname, '../../public/views/index.template.html'),
      inject: 'body',
      chunks: ['manifest', 'vendor', 'app'],
    }),
    new webpack.DefinePlugin({
      'process.env': {
        'NODE_ENV': JSON.stringify('development')
      }
    }),
    new webpack.optimize.CommonsChunkPlugin({
      names: ['vendor', 'manifest'],
    }),
    new WebpackCleanupPlugin(),
    // new BundleAnalyzerPlugin({
    //   analyzerPort: 8889
    // })
  ]
});
