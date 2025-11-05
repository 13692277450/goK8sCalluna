import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";

// https://vite.dev/config/
export default defineConfig({
  build: {
    chunkSizeWarningLimit: 1500, // 将限制提高到 1500 KB
  },
  plugins: [react()],
  server: {
    open: false,
    proxy: {
      "/api": {
        target: "http://localhost:5678",
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/api/, ""),
      },
    },
  },
});
