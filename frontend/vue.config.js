module.exports = {
    devServer: {
        proxy: {
            '/api': {
                target: {
                    host: "localhost",
                    protocol: 'http:',
                    port: 3000
                },
                changeOrigin: true
            }
        }
    }
}