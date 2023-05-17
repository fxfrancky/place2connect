import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vitejs.dev/config/
export default defineConfig({
  server: {
    proxy: {
        '/api': {
          target: 'place2connect-api:8080',
          changeOrigin: true,
          secure: false,      
          ws: true,
          rewrite: path => path.replace('/api', ''),
          configure: (proxy, _options) => {
            proxy.on('error', (err, _req, _res) => {
              console.log('proxy error', err);
            });
            proxy.on('proxyReq', (proxyReq, req, _res) => {
              console.log('Sending Request to the Target:', req.method, req.url);
            });
            proxy.on('proxyRes', (proxyRes, req, _res) => {
              console.log('Received Response from the Target:', proxyRes.statusCode, req.url);
            });
          },
        },
    },
    host: true,
    port: 3000,
    
  },
  plugins: [react()],
})

// place2connect-ui:3000

// place2connect-api:8080

// proxy: {
//         '/api': {
//           target: 'place2connect-api:8080',
//           changeOrigin: true,
//           secure: false,      
//           ws: true,
//           rewrite: path => path.replace('/api', ''),
//           configure: (proxy, _options) => {
//             proxy.on('error', (err, _req, _res) => {
//               console.log('proxy error', err);
//             });
//             proxy.on('proxyReq', (proxyReq, req, _res) => {
//               console.log('Sending Request to the Target:', req.method, req.url);
//             });
//             proxy.on('proxyRes', (proxyRes, req, _res) => {
//               console.log('Received Response from the Target:', proxyRes.statusCode, req.url);
//             });
//           },
//         }






// proxy: {
//       '/api': {
//         target: 'https://api.place2connect.com',
//         secure: true,
//         ws: true,
//         rewrite: path => path.replace('api', ''),
//       },
//     },