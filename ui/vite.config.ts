import { defineConfig } from 'vite'
import path from 'path'

export default defineConfig(({ mode }) => {
  const isDev = mode === 'development'

  return {
    // Set root to ui directory
    root: path.resolve(__dirname),

    // Set public directory for static assets (relative to project root)
    publicDir: path.resolve(__dirname, '../assets'),

    // Build configuration
    build: {
      // Output directory for production build (relative to project root)
      outDir: path.resolve(__dirname, '../dist'),

      // Generate manifest.json for production asset mapping
      manifest: true,

      // Output assets to dist/assets
      assetsDir: 'assets',

      // Clear output directory before build
      emptyOutDir: true,

      rollupOptions: {
        // Entry point for the application
        input: {
          main: path.resolve(__dirname, 'main.ts')
        },
        output: {
          // Keep consistent naming for easier debugging
          entryFileNames: 'assets/[name]-[hash].js',
          chunkFileNames: 'assets/[name]-[hash].js',
          assetFileNames: 'assets/[name]-[hash].[ext]'
        }
      }
    },

    // Server configuration for development
    server: {
      // Origin for Vite dev server
      origin: 'http://localhost:5173',

      // Port for Vite dev server
      port: 5173,

      // Strict port - fail if port is already in use
      strictPort: true,

      // Enable CORS for Go backend
      cors: true
    },

    resolve: {
      alias: {
        '@': path.resolve(__dirname),
        '@components': path.resolve(__dirname, 'components')
      }
    }
  }
})
