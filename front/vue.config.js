const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  outputDir: './../dist',
  assetsDir: 'static',

  pages: {
    index: {
      entry: 'src/main.js',
      template: 'public/index.html',
      filename: 'index.html',
      title: 'Demo Page',
      chunks: ['chunk-vendors', 'chunk-common', 'index']
    },

    admin: {
      entry: 'src/admin.js',
      template: 'public/admin.html',
      filename: 'admin.html',
      title: 'Admin Page',
      chunks: ['chunk-vendors', 'chunk-common', 'admin']
    },

  },

  lintOnSave: true,
  filenameHashing: false,
  productionSourceMap: false,
  transpileDependencies: true
})
