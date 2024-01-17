import { mergeConfig } from 'vite';
import eslint from 'vite-plugin-eslint';
import baseConfig from './vite.config.base';

export default mergeConfig(
    {
        mode: 'development',
        server: {
            open: true,
            fs: {
                strict: true,
            },
            proxy: { // 跨域代理
                '/api': {
                    target: 'http://123.57.245.226:5000',
                    rewrite: (path) => path.replace(/^\/api/, ""),
                    changeOrigin: true,
                },
            },
            host: '0.0.0.0', // 允许其他电脑访问
        },
        plugins: [
            eslint({
                cache: false,
                include: ['src/**/*.ts', 'src/**/*.tsx', 'src/**/*.vue'],
                exclude: ['node_modules'],
            }),
        ],
    },
    baseConfig
);