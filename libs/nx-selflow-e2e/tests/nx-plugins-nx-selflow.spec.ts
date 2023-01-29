import {
  checkFilesExist,
  ensureNxProject,
  readJson,
  runNxCommandAsync,
  uniq,
} from '@nrwl/nx-plugin/testing';
import {snakeCase} from "change-case";

describe('go-lib e2e', () => {
  // Setting up individual workspaces per
  // test can cause e2e runs to take a long time.
  // For this reason, we recommend each suite only
  // consumes 1 workspace. The tests should each operate
  // on a unique project in the workspace, such that they
  // are not dependant on one another.
  beforeAll(() => {
    ensureNxProject('@selflow/nx-selflow', 'dist/libs/nx-plugins/nx-selflow');
  });

  afterAll(() => {
    // `nx reset` kills the daemon, and performs
    // some work which can help clean up e2e leftovers
    runNxCommandAsync('reset');
  });

  it('should create go-lib in pkg directory', async () => {
    const project = uniq('go-lib');
    await runNxCommandAsync(
      `generate @selflow/nx-selflow:go-lib ${project} --directory pkg`
    );
    expect(() =>
      checkFilesExist(`pkg/${project}/${snakeCase(project)}.go`)
    ).not.toThrow();
  }, 120000);

  it('should create go-lib in internal directory', async () => {
    const project = uniq('go-lib');
    await runNxCommandAsync(
      `generate @selflow/nx-selflow:go-lib ${project} --directory internal`
    );
    expect(() =>
      checkFilesExist(`internal/${project}/${snakeCase(project)}.go`)
    ).not.toThrow();
  }, 120000);

  describe('--directory', () => {
    it('should create src in the specified directory', async () => {
      const project = uniq('go-lib');
      await runNxCommandAsync(
        `generate @selflow/nx-selflow:go-lib ${project} --directory subdir`
      );
      expect(() =>
        checkFilesExist(`libs/subdir/${project}/src/index.ts`)
      ).not.toThrow();
    }, 120000);
  });

  describe('--tags', () => {
    it('should add tags to the project', async () => {
      const projectName = uniq('go-lib');
      ensureNxProject('@selflow/nx-selflow', 'dist/libs/nx-plugins/nx-selflow');
      await runNxCommandAsync(
        `generate @selflow/nx-selflow:go-lib ${projectName} --tags e2etag,e2ePackage`
      );
      const project = readJson(`libs/${projectName}/project.json`);
      expect(project.tags).toEqual(['e2etag', 'e2ePackage']);
    }, 120000);
  });
});
