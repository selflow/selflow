// @ts-check
// Note: type annotations allow type checking and IDEs autocompletion

const lightCodeTheme = require('prism-react-renderer/themes/github');
const darkCodeTheme = require('prism-react-renderer/themes/dracula');

const repositoryUrl = 'https://github.com/selflow/selflow';
const repositoryDocumentationPath = 'apps/selflow-documentation';
const documentationEditUrl = `${repositoryUrl}/edit/main/${repositoryDocumentationPath}/`;

/** @type {import('@docusaurus/types').Config} */
const config = {
  title: 'Selflow - Documentation',
  tagline: 'A Workflow Orchestration Framework designed to be self-hosted',
  url: 'https://selflow.github.io',
  baseUrl: '/selflow/',
  onBrokenLinks: 'throw',
  onBrokenMarkdownLinks: 'warn',
  favicon: 'img/favicon.ico',
  organizationName: 'selflow',
  projectName: 'selflow',

  presets: [
    [
      'classic',
      /** @type {import('@docusaurus/preset-classic').Options} */
      ({
        docs: {
          sidebarPath: require.resolve('./sidebars.js'),
          editUrl: documentationEditUrl,
        },
        blog: {
          showReadingTime: true,
          editUrl: documentationEditUrl,
        },
        theme: {
          customCss: require.resolve('./src/css/custom.css'),
        },
      }),
    ],
  ],

  themeConfig:
    /** @type {import('@docusaurus/preset-classic').ThemeConfig} */
    ({
      navbar: {
        title: 'Selflow',
        logo: {
          alt: 'Selflow Logo',
          src: 'img/selflow-logo.png',
        },
        items: [
          {
            type: 'doc',
            docId: 'intro',
            position: 'left',
            label: 'Documentations',
          },
          { to: '/blog', label: 'Blog', position: 'left' },
          {
            href: repositoryUrl,
            label: 'GitHub',
            position: 'right',
          },
        ],
      },
      footer: {
        style: 'dark',
        links: [
          {
            title: 'Docs',
            items: [
              {
                label: 'Tutorial',
                to: '/docs/intro',
              },
            ],
          },
          {
            title: 'Community',
            items: [
              {
                label: 'Stack Overflow',
                href: 'https://stackoverflow.com/questions/tagged/selflow',
              },
              {
                label: 'Twitter',
                href: 'https://twitter.com/AnthonyJhoiro',
              },
            ],
          },
          {
            title: 'More',
            items: [
              {
                label: 'Blog',
                to: '/blog',
              },
              {
                label: 'GitHub',
                href: repositoryUrl,
              },
            ],
          },
        ],
        copyright: `Copyright © ${new Date().getFullYear()} Selflow community. Built with Docusaurus ❤️`,
      },
      prism: {
        theme: lightCodeTheme,
        darkTheme: darkCodeTheme,
      },
    }),
};

module.exports = config;
