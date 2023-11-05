import { selflowNodeClient } from './selflow-node-client';

describe('selflowNodeClient', () => {
  it('should work', () => {
    expect(selflowNodeClient()).toEqual('selflow-node-client');
  });
});
