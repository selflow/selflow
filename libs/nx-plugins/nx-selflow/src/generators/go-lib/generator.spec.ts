import { createTreeWithEmptyWorkspace } from '@nrwl/devkit/testing';
import { Tree, readProjectConfiguration } from '@nrwl/devkit';

import generator from './generator';
import { NxPluginsNxSelflowGeneratorSchema } from './schema';

describe('go-lib generator', () => {
  let appTree: Tree;
  const options: NxPluginsNxSelflowGeneratorSchema = {
    name: 'test',
    directory: 'pkg',
  };

  beforeEach(() => {
    appTree = createTreeWithEmptyWorkspace();
  });

  it('should run successfully', async () => {
    await generator(appTree, options);
    const config = readProjectConfiguration(appTree, 'test');
    expect(config).toBeDefined();
  });
});
