import {promisify} from 'util';
import {DaemonClient} from '../generated/daemon';
import * as grpc from '@grpc/grpc-js';
import {mapObject} from './utils';

export function selflowNodeClient(): string {
  return 'selflow-node-client';
}

export type StepStatus = {
  name: string;
};

export type DaemonState = {
  state: Record<string, StepStatus>;
  dependencies: Record<string, string[]>;
};

export class DaemonService extends DaemonClient {
  constructor(target: string) {
    super(target, grpc.credentials.createInsecure());
  }

  public async doGetRunStatus(runId: string) {
    return promisify(this.getRunStatus)
      .bind(this)({ runId })
      .then<DaemonState>((response: any) => {
        if (!response) {
          return { state: {}, dependencies: {} };
        }

        const responseState = response.state ?? {};
        const responseDependencies = response.dependencies ?? {};
        return {
          state: mapObject<{ name: string }, StepStatus>(
            responseState,
            (v) => ({ ...v })
          ),
          dependencies: mapObject<{ dependencies: string[] }, string[]>(
            responseDependencies,
            (v) => v.dependencies
          ),
        };
      });
  }
}
