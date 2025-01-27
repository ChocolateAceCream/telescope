import { resolve, dirname } from "path";
import { fileURLToPath } from "url";
import { defineConfig, loadEnv } from "vite";
import react from "@vitejs/plugin-react";
import svgr from "vite-plugin-svgr";
import { visualizer } from "rollup-plugin-visualizer";
import Banner from "vite-plugin-banner";
import { NodeGlobalsPolyfillPlugin } from "@esbuild-plugins/node-globals-polyfill";
import { NodeModulesPolyfillPlugin } from "@esbuild-plugins/node-modules-polyfill";
function pathResolve(...args: string[]) {
  const __filename = fileURLToPath(import.meta.url);
  const __dirname = dirname(__filename);
  return resolve(__dirname, "./", ...args);
  // return resolve(__dirname, '.', ...args)
}
// https://vitejs.dev/config/
export default defineConfig((params) => {
  const { command, mode } = params;
  const ENV = loadEnv(mode, process.cwd());
  const timestamp = Date.now();
  console.info(
    `--- running mode: ${mode}, command: ${command}, ENV: ${JSON.stringify(
      ENV
    )} ---`
  );
  console.log(svgr);
  return {
    base: "./",
    root: "./", // js导入的资源路径，src
    resolve: {
      extensions: [".json", ".js", ".ts", ".vue", ".jsx", ".tsx"],
      alias: {
        "@": pathResolve("src"),
        "/img": pathResolve("src/assets/images"),
        "vue-i18n": "vue-i18n/dist/vue-i18n.cjs.js",
      },
    },
    server: {
      port: parseInt(ENV.VITE_APP_PORT, 10),
      host: ENV.VITE_APP_HOST,
      proxy: {
        "/backend": {
          target: ENV.VITE_APP_DEV_PROXY,
          changeOrigin: true,
          rewrite: (path) => path.replace(/^\/backend/, ""),
        },
        "/websocket": {
          target: ENV.VITE_WEBSOCKET_LOCAL,
          changeOrigin: true,
          rewrite: (path) => path.replace(/^\/websocket/, ""),
          ws: true,
          // configure: (proxy, _options) => {
          //   proxy.on('error', (err, _req, _res) => {
          //     console.log('proxy error', err)
          //   })
          //   proxy.on('proxyReq', (proxyReq, req, _res) => {
          //     console.log('Sending Request to the Target:', req.method, req.url)
          //   })
          //   proxy.on('proxyRes', (proxyRes, req, _res) => {
          //     console.log('Received Response from the Target:', proxyRes.statusCode, req.url)
          //   })
          // },
        },
      },
    },
    build: {
      minify: "terser", // 必须启用：terserOptions配置才会有效
      terserOptions: {
        compress: {
          // 生产环境时移除console.log调试代码
          drop_console: true,
          drop_debugger: true,
        },
      },
      target: "es2015",
      manifest: false,
      sourcemap: false,
      outDir: "dist",
      build: {
        rollupOptions: {
          output: {
            manualChunks: {
              // moment: ['moment'],
              "lodash-es": ["lodash-es"],
              "md-editor-v3": ["md-editor-v3"],
              dayjs: "dayjs",
            },
          },
        },
      },
    },
    css: {
      preprocessorOptions: {
        scss: {
          // additionalData: `$injectedColor: orange;`
          additionalData: `
            @import "@/assets/styles/globalInjectedData.scss";
          `,
        },
      },
    },
    optimizeDeps: {
      esbuildOptions: {
        define: {
          global: "globalThis",
        },
        plugins: [
          NodeGlobalsPolyfillPlugin({
            buffer: true,
          }),
          NodeModulesPolyfillPlugin(),
        ],
      },
    },
    plugins: [
      // analyze pkg size
      visualizer({
        open: true,
        gzipSize: true,
        brotliSize: true,
      }),
      react(),
      svgr({
        svgrOptions: {
          // SVGR options

          // Set this to true if you want the SVG to be scaled to fit a square viewBox
          icon: true,

          // When set to false, this option removes the width and height attributes from the SVG, making it more flexible to style with CSS.
          // dimensions: false,
        },
      }),
      [
        Banner(`
  #####                                                           #                   #####
#     # #    #  ####   ####   ####  #        ##   ##### ######   # #    ####  ###### #     # #####  ######   ##   #    #
#       #    # #    # #    # #    # #       #  #    #   #       #   #  #    # #      #       #    # #       #  #  #    #
#       ###### #    # #      #    # #      #    #   #   #####  #     # #      #####  #       #    # #####  #    # ##  ##
#       #    # #    # #      #    # #      ######   #   #      ####### #      #      #       #####  #      ###### # ## #
#     # #    # #    # #    # #    # #      #    #   #   #      #     # #    # #      #     # #   #  #      #    # #    #
 #####  #    #  ####   ####   ####  ###### #    #   #   ###### #     #  ####  ######  #####  #    # ###### #    # #    #
        \n Build on Time : ${timestamp}`),
      ],
    ],
  };
});
