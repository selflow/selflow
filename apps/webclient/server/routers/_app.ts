import {z} from 'zod';
import {procedure, router} from '../trpc';
import {DaemonService} from "@selflow/selflow-node-client";

const daemonService = new DaemonService(process.env.DAEMON_URL ?? "127.0.0.1:1001")


export const appRouter = router({
  status: procedure
    .input(
      z.string()
    )
    .query(async opts => {
      return await daemonService.doGetRunStatus(opts.input);
    })

});
// export type definition of API
export type AppRouter = typeof appRouter;
