const webpack = require("webpack");
module.exports = {
  devServer:{
    proxy: 'http://localhost:9100/',
  },
  pages: {
    // 'about': {
    //   entry: 'src/pages/About/main.js',
    //   template: 'public/index.html',
    //   title: 'about page',
    //   chunks: ['chunk-vendors', 'chunk-common', 'about']
    // },
    'user':{
      entry: 'src/pages/User/main.js',
      template: 'public/index.html',
      title: 'user main page',
      chunks: ['chunk-vendors', 'chunk-common', 'user']
    },
    'posts': {
      entry: 'src/pages/Posts/main.js',
      template: 'public/index.html',
      title: 'posts page',
      chunks: ['chunk-vendors', 'chunk-common', 'posts']
    },
    'login-register': {
      entry: 'src/pages/LoginRegister/main.js',
      template: 'public/index.html',
      title: 'login-register page',
      chunks: ['chunk-vendors', 'chunk-common', 'login-register']
    }
  },
  configureWebpack: {
    plugins: [
      new webpack.ProvidePlugin({
        $: 'jquery',
        jQuery: 'jquery',
        'window.jQuery': 'jquery',
        Popper: ['popper.js', 'default']
      })
    ]
  }
}
