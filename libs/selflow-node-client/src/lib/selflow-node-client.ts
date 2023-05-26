import { promisify } from 'util';
import {
  DaemonClient,
  GetRunStatus_Request,
  GetRunStatus_Response,
  StartRun_Request,
  StartRun_Response,
} from '../generated/daemon';
import * as grpc from '@grpc/grpc-js';
import { mapObject } from './utils';

export function selflowNodeClient(): string {
  return 'selflow-node-client';
}

export type StepStatus = {
  name: string;
  code: number;
  isFinished: boolean;
  isCancellable: boolean;
};

export type WorkflowConfiguration = {
  workflow: {
    timeout: string;
    steps: {
      [key: string]: {
        kind: string;
        needs: string[];
        with: {
          image: string;
          commands: string;
        };
      };
    };
  };
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
    return promisify<GetRunStatus_Request, GetRunStatus_Response>(
      this.getRunStatus
    )
      .bind(this)({ runId })
      .then((response) => {
        if (!response) {
          return { state: {}, dependencies: {}, steps: {} };
        }

        const responseState = response.state ?? {};
        const responseDependencies = response.dependencies ?? {};
        return {
          state: mapObject<StepStatus, StepStatus>(responseState, (v) => ({
            ...v,
          })),
          dependencies: mapObject<{ dependencies: string[] }, string[]>(
            responseDependencies,
            (v) => v.dependencies
          ),
          steps: mapObject(response.stepDefinitions, (v) =>
            JSON.parse(Buffer.from(v).toString('utf-8'))
          ),
        };
      });
  }

  public async doStartRun(runConfig: WorkflowConfiguration) {
    const encodedWorkflow = new TextEncoder().encode(JSON.stringify(runConfig));
    return promisify<StartRun_Request, StartRun_Response>(this.startRun)
      .bind(this)({ runConfig: encodedWorkflow })
      .then((response) => response.runId);
  }
}
