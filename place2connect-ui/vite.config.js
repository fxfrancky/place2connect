import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vitejs.dev/config/
export default defineConfig({
  server: {
    proxy: {
      '/api': {
        forward: 'https://api.place2connect.com',
        secure: true,
        ws: true,
        rewrite: path => path.replace('api', ''),
      },
    },
    host: true,
    port: 3000,
    
  },
  plugins: [react()],
})
