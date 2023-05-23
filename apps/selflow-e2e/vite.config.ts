/// <reference types="vitest" />
import { defineConfig } from 'vite';

import viteTsConfigPaths from 'vite-tsconfig-paths';
import { join } from 'path';

export default defineConfig({
  cacheDir: '../../node_modules/.vite/selflow-e2e',

  plugins: [
    viteTsConfigPaths({
      root: '../../',
    }),
  ],

  test: {
    globals: true,
    cache: {
      dir: '../../node_modules/.vitest',
    },
    environment: 'node',
    include: ['src/**/*.{test,spec}.{js,mjs,cjs,ts,mts,cts,jsx,tsx}'],
    singleThread: true,
    globalSetup: join(__dirname, 'src/setup.ts'),
    maxConcurrency: 1,
  },
});
