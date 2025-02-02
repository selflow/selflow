import { describe, expect, it } from 'vitest';
import { groupGoFileDirectories } from './goFileUtils';

describe('groupGoFileDirectories', () => {
  it('should group files by directory', () => {
    const files = ['src/main.go', 'src/lib/util.go', 'src/lib/helper.go'];

    const result = groupGoFileDirectories(files);

    expect(result.get('src')).toEqual(['src/main.go']);
    expect(result.get('src/lib')).toEqual([
      'src/lib/util.go',
      'src/lib/helper.go',
    ]);
  });

  it('should handle empty input', () => {
    const result = groupGoFileDirectories([]);
    expect(result.size).toBe(0);
  });
});
