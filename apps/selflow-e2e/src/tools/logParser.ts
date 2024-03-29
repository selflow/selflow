import { WorkflowExecutionTrace } from './trace';

export interface LogLine {
  timeStamp: Date;
  type: string;
  name: string;
  message: string;
  metadata: any;
  order: number;
}

export interface Event {
  timeStamp: Date;
  eventType: 'start' | 'stop';
  stepId: string;
  status: string;
  order: number;
}

function parseLogLine(log: string, index: number): LogLine | null {
  try {
    const logDetails = JSON.parse(log);
    return {
      timeStamp: new Date(logDetails['time']),
      type: logDetails['level'],
      name: logDetails['stepId'],
      message: logDetails['msg'],
      metadata: logDetails['metadata'] ?? { ...logDetails },
      order: index,
    };
  } catch (e) {
    return null;
  }
}

export function parseLogs(logs: string): WorkflowExecutionTrace {
  console.log(logs);
  const parsedLogs = logs
    .split('\n')
    .map(parseLogLine)
    .filter((log) => !!log);

  const events: Event[] = parsedLogs
    .filter((log) => log.type === 'INFO')
    .map((log) => {
      if (log.message.includes('Step started')) {
        return {
          timeStamp: log.timeStamp,
          eventType: 'start',
          stepId: log.name,
          status: '',
          order: log.order,
        } as Event;
      }

      if (log.message.includes('Step terminated')) {
        return {
          timeStamp: log.timeStamp,
          eventType: 'stop',
          stepId: log.name,
          status: log.metadata['stepStatus'],
          order: log.order,
        } as Event;
      }

      return null;
    })
    .filter((log) => !!log);

  const stepLogs = parsedLogs.reduce<Record<string, string[]>>((acc, log) => {
    if (!log.name || log.name.length === 0) return acc;
    return {
      ...acc,
      [log.name]: [...(acc[log.name] ?? []), log.message],
    };
  }, {});

  console.log(
    JSON.stringify(
      {
        stepLogs,
        logs: parsedLogs,
        events,
      },
      null,
      4
    )
  );

  return {
    stepLogs,
    logs: parsedLogs,
    events,
  };
}
