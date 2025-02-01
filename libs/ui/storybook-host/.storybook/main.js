const config = {
  stories: ['../../**/*.@(mdx|stories.@(js|jsx|ts|tsx))', {
    directory: '../../components-kit',
    titlePrefix: 'ui-kit',
    files: '**/*.@(mdx|stories.@(js|jsx|ts|tsx))'
  }, {
    directory: '../../workflow-editor',
    titlePrefix: 'workflow-editor',
    files: '**/*.@(mdx|stories.@(js|jsx|ts|tsx))'
  }],

  addons: ['@storybook/addon-essentials', '@chromatic-com/storybook'],

  framework: {
    name: '@storybook/nextjs',
    options: {},
  },

  docs: {},

  typescript: {
    reactDocgen: 'react-docgen-typescript'
  }
};

export default config;

// To customize your webpack configuration you can use the webpackFinal field.
// Check https://storybook.js.org/docs/react/builders/webpack#extending-storybooks-webpack-config
// and https://nx.dev/packages/storybook/documents/custom-builder-configs
