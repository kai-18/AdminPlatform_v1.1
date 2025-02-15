module.exports = {
  devServer: {
    proxy: {
      '/employees': {
        target: 'http://localhost:8080',
        changeOrigin: true,
        pathRewrite: { '^/employees': '/employees' },
      }
    }
  }
}
