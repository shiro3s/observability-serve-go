import {defineConfig} from "vite"

export default defineConfig({
  publicDir: "./schema",
  define: {
    global: "window",
  }
})