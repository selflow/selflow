import { z } from 'zod';
import { procedure, router } from '../trpc';
import { DaemonService } from '@selflow/selflow-node-client';

const daemonService = new DaemonService(
  process.env.DAEMON_URL ?? '127.0.0.1:10011'
);

export const appRouter = router({
  status: procedure.input(z.string()).query(async (opts) => {
    return await daemonService.doGetRunStatus(opts.input);
  }),
  run: procedure
    .input(
      z.object({
        workflow: z.object({
          timeout: z.string(),
          steps: z.record(
            z.object({
              kind: z.string(),
              needs: z.array(z.string()),
              with: z.object({
                image: z.string(),
                commands: z.string(),
              }),
            })
          ),
        }),
      })
    )
    .mutation((opts) => daemonService.doStartRun(opts.input)),
});
// export type definition of API
export type AppRouter = typeof appRouter;
