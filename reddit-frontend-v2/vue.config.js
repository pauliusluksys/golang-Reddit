const webpack = require("webpack");
module.exports = {
  pages: {
    // 'about': {
    //   entry: 'src/pages/About/main.js',
    //   template: 'public/index.html',
    //   title: 'about page',
    //   chunks: ['chunk-vendors', 'chunk-common', 'about']
    // },
    // 'chat': {
    //   entry: 'src/pages/Chat/main.js',
    //   template: 'public/index.html',
    //   title: 'chat page',
    //   chunks: ['chunk-vendors', 'chunk-common', 'chat']
    // },
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
