import { WorkflowExecutionTrace } from './trace';

export interface LogLine {
  timeStamp: Date;
  type: string;
  name: string;
  message: string;
  order: number;
}

export interface Event {
  timeStamp: Date;
  eventType: 'start' | 'stop';
  stepId: string;
  status: string;
  order: number;
}

const logRegex =
  /^(\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}) \[([A-Z]+)]\s+(\S+): ([^\n]+)/;

function parseLogLine(log: string, index: number): LogLine | null {
  const match = log.match(logRegex);
  if (!match) {
    return null;
  }
  return {
    timeStamp: new Date(match[1]),
    type: match[2],
    name: match[3],
    message: match[4],
    order: index,
  };
}

const stepIdRegex = /step-id=(\S+)/;
const statusRegex = /status=(\S+)/;

const getStepId = (log: string) => log.match(stepIdRegex)?.[1];
const getStatus = (log: string) => log.match(statusRegex)?.[1];

export function parseLogs(logs: string): WorkflowExecutionTrace {
  const parsedLogs = logs
    .split('\n')
    .map(parseLogLine)
    .filter((log) => !!log);

  const events: Event[] = parsedLogs
    .filter((log) => log.type === 'INFO')
    .map((log) => {
      if (log.message.includes('step started')) {
        return {
          timeStamp: log.timeStamp,
          eventType: 'start',
          stepId: getStepId(log.message),
          status: '',
          order: log.order,
        } as Event;
      }

      if (log.message.includes('step terminated')) {
        return {
          timeStamp: log.timeStamp,
          eventType: 'stop',
          stepId: getStepId(log.message),
          status: getStatus(log.message),
          order: log.order,
        } as Event;
      }

      return null;
    })
    .filter((log) => !!log);

  const stepLogs = parsedLogs
    .filter((log) => log.type === 'DEBUG')
    .reduce<Record<string, string[]>>((acc, log) => {
      const name = log.name.split('.').pop();
      return {
        ...acc,
        [name]: [...(acc[name] ?? []), log.message],
      };
    }, {});

  return {
    stepLogs,
    logs: parsedLogs,
    events,
  };
}
