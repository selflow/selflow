import { describe, expect, it } from 'vitest';
import { generateAppProject, generateLibProject } from './projectConfigBuilder';

describe('generateAppProject', () => {
  it('should generate correct app project configuration', () => {
    const project = {
      packageName: 'main',
      dir: 'src/app',
      sourceFile: 'src/app/main.go',
    };

    const [sourceFile, config] = generateAppProject(project);

    expect(sourceFile).toBe('src/app/main.go');
    expect(config.projects['src/app']).toMatchObject({
      projectType: 'application',
      tags: ['lang:go'],
    });
  });
});

describe('generateLibProject', () => {
  it('should generate correct library project configuration', () => {
    const project = {
      packageName: 'mylib',
      dir: 'src/lib',
      sourceFile: 'src/lib/lib.go',
    };

    const [sourceFile, config] = generateLibProject(project);

    expect(sourceFile).toBe('src/lib/lib.go');
    expect(config.projects['src/lib']).toMatchObject({
      projectType: 'library',
      tags: ['lang:go'],
    });
  });
});
