import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  // lintOnSave: false,
  server: {
    port: '5173',
    proxy: {
        '/api': {
            target: "http://8.130.98.106:9999",
            changeOrigin: true,
            rewrite: (path) => path.replace(/^\/api/, ""), 
        },
    },
  },
})
