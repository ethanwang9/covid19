import path from 'path'
import {loadEnv} from 'vite'
import vue from '@vitejs/plugin-vue'
import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import {ElementPlusResolver} from 'unplugin-vue-components/resolvers'
import ElementPlus from "unplugin-element-plus/vite"
import IconsResolver from 'unplugin-icons/resolver'
import Icons from 'unplugin-icons/vite'

export default (({command, mode}) => {
    const env = loadEnv(mode, process.cwd(), '')
    return {
        resolve: {
            alias: {
                '@': path.resolve(__dirname, 'src')
            },
        },
        css: {
            preprocessorOptions: {
                scss: {
                    additionalData: `@use "@/styles/element/index.scss" as *;`,
                },
            },
        },
        plugins: [
            vue(),
            AutoImport({
                imports: ['vue'],
                resolvers: [
                    ElementPlusResolver(),
                    IconsResolver({
                        prefix: 'Icon',
                    }),
                ],
            }),
            Components({
                resolvers: [
                    IconsResolver({
                        enabledCollections: ['ep'],
                    }),
                    ElementPlusResolver(),
                ],
            }),
            Icons({
                autoInstall: true,
            }),
            ElementPlus({
                useSource: true,
            }),
        ],
        server: {
            host: true,
            port: env.VITE_WEB_PORT,
            proxy: {
                [env.VITE_API_PATH]: {
                    target: `${env.VITE_SERVER_ADDR}:${env.VITE_SERVER_PORT}/v1/`,
                    changeOrigin: true,
                    rewrite: path => path.replace(new RegExp('^' + env.VITE_API_PATH), ''),
                },
            },
        }
    }
})