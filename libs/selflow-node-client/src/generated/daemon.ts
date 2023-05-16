/* eslint-disable */
import {
  CallOptions,
  ChannelCredentials,
  Client,
  ClientOptions,
  ClientReadableStream,
  ClientUnaryCall,
  handleServerStreamingCall,
  handleUnaryCall,
  makeGenericClientConstructor,
  Metadata,
  ServiceError,
  UntypedServiceImplementation,
} from "@grpc/grpc-js";
import * as _m0 from "protobufjs/minimal";

export const protobufPackage = "";

export enum DiagnosticType {
  ERROR = 0,
  WARNING = 1,
  INFO = 2,
  DEBUG = 3,
  UNRECOGNIZED = -1,
}

export function diagnosticTypeFromJSON(object: any): DiagnosticType {
  switch (object) {
    case 0:
    case "ERROR":
      return DiagnosticType.ERROR;
    case 1:
    case "WARNING":
      return DiagnosticType.WARNING;
    case 2:
    case "INFO":
      return DiagnosticType.INFO;
    case 3:
    case "DEBUG":
      return DiagnosticType.DEBUG;
    case -1:
    case "UNRECOGNIZED":
    default:
      return DiagnosticType.UNRECOGNIZED;
  }
}

export function diagnosticTypeToJSON(object: DiagnosticType): string {
  switch (object) {
    case DiagnosticType.ERROR:
      return "ERROR";
    case DiagnosticType.WARNING:
      return "WARNING";
    case DiagnosticType.INFO:
      return "INFO";
    case DiagnosticType.DEBUG:
      return "DEBUG";
    case DiagnosticType.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface Diagnostic {
  type: DiagnosticType;
  message: string;
}

export interface StartRun {
}

export interface StartRun_Request {
  runConfig: Uint8Array;
}

export interface StartRun_Response {
  diagnostics: Diagnostic[];
  runId: string;
}

export interface GetLogStream {
}

export interface GetLogStream_Request {
  runId: string;
}

export interface Log {
  level: string;
  dateTime: string;
  name: string;
  message: string;
}

export interface GetRunStatus {
}

export interface GetRunStatus_Request {
  runId: string;
}

export interface GetRunStatus_Status {
  name: string;
}

export interface GetRunStatus_Dependence {
  dependencies: string[];
}

export interface GetRunStatus_Response {
  state: { [key: string]: GetRunStatus_Status };
  dependencies: { [key: string]: GetRunStatus_Dependence };
}

export interface GetRunStatus_Response_StateEntry {
  key: string;
  value: GetRunStatus_Status | undefined;
}

export interface GetRunStatus_Response_DependenciesEntry {
  key: string;
  value: GetRunStatus_Dependence | undefined;
}

function createBaseDiagnostic(): Diagnostic {
  return { type: 0, message: "" };
}

export const Diagnostic = {
  encode(message: Diagnostic, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.type !== 0) {
      writer.uint32(8).int32(message.type);
    }
    if (message.message !== "") {
      writer.uint32(18).string(message.message);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Diagnostic {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDiagnostic();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.type = reader.int32() as any;
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.message = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Diagnostic {
    return {
      type: isSet(object.type) ? diagnosticTypeFromJSON(object.type) : 0,
      message: isSet(object.message) ? String(object.message) : "",
    };
  },

  toJSON(message: Diagnostic): unknown {
    const obj: any = {};
    message.type !== undefined && (obj.type = diagnosticTypeToJSON(message.type));
    message.message !== undefined && (obj.message = message.message);
    return obj;
  },

  create<I extends Exact<DeepPartial<Diagnostic>, I>>(base?: I): Diagnostic {
    return Diagnostic.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<Diagnostic>, I>>(object: I): Diagnostic {
    const message = createBaseDiagnostic();
    message.type = object.type ?? 0;
    message.message = object.message ?? "";
    return message;
  },
};

function createBaseStartRun(): StartRun {
  return {};
}

export const StartRun = {
  encode(_: StartRun, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): StartRun {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseStartRun();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(_: any): StartRun {
    return {};
  },

  toJSON(_: StartRun): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<StartRun>, I>>(base?: I): StartRun {
    return StartRun.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<StartRun>, I>>(_: I): StartRun {
    const message = createBaseStartRun();
    return message;
  },
};

function createBaseStartRun_Request(): StartRun_Request {
  return { runConfig: new Uint8Array() };
}

export const StartRun_Request = {
  encode(message: StartRun_Request, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.runConfig.length !== 0) {
      writer.uint32(10).bytes(message.runConfig);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): StartRun_Request {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseStartRun_Request();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.runConfig = reader.bytes();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): StartRun_Request {
    return { runConfig: isSet(object.runConfig) ? bytesFromBase64(object.runConfig) : new Uint8Array() };
  },

  toJSON(message: StartRun_Request): unknown {
    const obj: any = {};
    message.runConfig !== undefined &&
      (obj.runConfig = base64FromBytes(message.runConfig !== undefined ? message.runConfig : new Uint8Array()));
    return obj;
  },

  create<I extends Exact<DeepPartial<StartRun_Request>, I>>(base?: I): StartRun_Request {
    return StartRun_Request.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<StartRun_Request>, I>>(object: I): StartRun_Request {
    const message = createBaseStartRun_Request();
    message.runConfig = object.runConfig ?? new Uint8Array();
    return message;
  },
};

function createBaseStartRun_Response(): StartRun_Response {
  return { diagnostics: [], runId: "" };
}

export const StartRun_Response = {
  encode(message: StartRun_Response, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.diagnostics) {
      Diagnostic.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.runId !== "") {
      writer.uint32(18).string(message.runId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): StartRun_Response {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseStartRun_Response();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.diagnostics.push(Diagnostic.decode(reader, reader.uint32()));
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.runId = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): StartRun_Response {
    return {
      diagnostics: Array.isArray(object?.diagnostics) ? object.diagnostics.map((e: any) => Diagnostic.fromJSON(e)) : [],
      runId: isSet(object.runId) ? String(object.runId) : "",
    };
  },

  toJSON(message: StartRun_Response): unknown {
    const obj: any = {};
    if (message.diagnostics) {
      obj.diagnostics = message.diagnostics.map((e) => e ? Diagnostic.toJSON(e) : undefined);
    } else {
      obj.diagnostics = [];
    }
    message.runId !== undefined && (obj.runId = message.runId);
    return obj;
  },

  create<I extends Exact<DeepPartial<StartRun_Response>, I>>(base?: I): StartRun_Response {
    return StartRun_Response.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<StartRun_Response>, I>>(object: I): StartRun_Response {
    const message = createBaseStartRun_Response();
    message.diagnostics = object.diagnostics?.map((e) => Diagnostic.fromPartial(e)) || [];
    message.runId = object.runId ?? "";
    return message;
  },
};

function createBaseGetLogStream(): GetLogStream {
  return {};
}

export const GetLogStream = {
  encode(_: GetLogStream, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetLogStream {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetLogStream();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(_: any): GetLogStream {
    return {};
  },

  toJSON(_: GetLogStream): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<GetLogStream>, I>>(base?: I): GetLogStream {
    return GetLogStream.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetLogStream>, I>>(_: I): GetLogStream {
    const message = createBaseGetLogStream();
    return message;
  },
};

function createBaseGetLogStream_Request(): GetLogStream_Request {
  return { runId: "" };
}

export const GetLogStream_Request = {
  encode(message: GetLogStream_Request, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.runId !== "") {
      writer.uint32(10).string(message.runId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetLogStream_Request {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetLogStream_Request();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.runId = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetLogStream_Request {
    return { runId: isSet(object.runId) ? String(object.runId) : "" };
  },

  toJSON(message: GetLogStream_Request): unknown {
    const obj: any = {};
    message.runId !== undefined && (obj.runId = message.runId);
    return obj;
  },

  create<I extends Exact<DeepPartial<GetLogStream_Request>, I>>(base?: I): GetLogStream_Request {
    return GetLogStream_Request.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetLogStream_Request>, I>>(object: I): GetLogStream_Request {
    const message = createBaseGetLogStream_Request();
    message.runId = object.runId ?? "";
    return message;
  },
};

function createBaseLog(): Log {
  return { level: "", dateTime: "", name: "", message: "" };
}

export const Log = {
  encode(message: Log, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.level !== "") {
      writer.uint32(10).string(message.level);
    }
    if (message.dateTime !== "") {
      writer.uint32(18).string(message.dateTime);
    }
    if (message.name !== "") {
      writer.uint32(26).string(message.name);
    }
    if (message.message !== "") {
      writer.uint32(34).string(message.message);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Log {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLog();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.level = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.dateTime = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.name = reader.string();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.message = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Log {
    return {
      level: isSet(object.level) ? String(object.level) : "",
      dateTime: isSet(object.dateTime) ? String(object.dateTime) : "",
      name: isSet(object.name) ? String(object.name) : "",
      message: isSet(object.message) ? String(object.message) : "",
    };
  },

  toJSON(message: Log): unknown {
    const obj: any = {};
    message.level !== undefined && (obj.level = message.level);
    message.dateTime !== undefined && (obj.dateTime = message.dateTime);
    message.name !== undefined && (obj.name = message.name);
    message.message !== undefined && (obj.message = message.message);
    return obj;
  },

  create<I extends Exact<DeepPartial<Log>, I>>(base?: I): Log {
    return Log.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<Log>, I>>(object: I): Log {
    const message = createBaseLog();
    message.level = object.level ?? "";
    message.dateTime = object.dateTime ?? "";
    message.name = object.name ?? "";
    message.message = object.message ?? "";
    return message;
  },
};

function createBaseGetRunStatus(): GetRunStatus {
  return {};
}

export const GetRunStatus = {
  encode(_: GetRunStatus, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetRunStatus {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetRunStatus();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(_: any): GetRunStatus {
    return {};
  },

  toJSON(_: GetRunStatus): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<GetRunStatus>, I>>(base?: I): GetRunStatus {
    return GetRunStatus.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetRunStatus>, I>>(_: I): GetRunStatus {
    const message = createBaseGetRunStatus();
    return message;
  },
};

function createBaseGetRunStatus_Request(): GetRunStatus_Request {
  return { runId: "" };
}

export const GetRunStatus_Request = {
  encode(message: GetRunStatus_Request, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.runId !== "") {
      writer.uint32(10).string(message.runId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetRunStatus_Request {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetRunStatus_Request();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.runId = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetRunStatus_Request {
    return { runId: isSet(object.runId) ? String(object.runId) : "" };
  },

  toJSON(message: GetRunStatus_Request): unknown {
    const obj: any = {};
    message.runId !== undefined && (obj.runId = message.runId);
    return obj;
  },

  create<I extends Exact<DeepPartial<GetRunStatus_Request>, I>>(base?: I): GetRunStatus_Request {
    return GetRunStatus_Request.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetRunStatus_Request>, I>>(object: I): GetRunStatus_Request {
    const message = createBaseGetRunStatus_Request();
    message.runId = object.runId ?? "";
    return message;
  },
};

function createBaseGetRunStatus_Status(): GetRunStatus_Status {
  return { name: "" };
}

export const GetRunStatus_Status = {
  encode(message: GetRunStatus_Status, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetRunStatus_Status {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetRunStatus_Status();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.name = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetRunStatus_Status {
    return { name: isSet(object.name) ? String(object.name) : "" };
  },

  toJSON(message: GetRunStatus_Status): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    return obj;
  },

  create<I extends Exact<DeepPartial<GetRunStatus_Status>, I>>(base?: I): GetRunStatus_Status {
    return GetRunStatus_Status.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetRunStatus_Status>, I>>(object: I): GetRunStatus_Status {
    const message = createBaseGetRunStatus_Status();
    message.name = object.name ?? "";
    return message;
  },
};

function createBaseGetRunStatus_Dependence(): GetRunStatus_Dependence {
  return { dependencies: [] };
}

export const GetRunStatus_Dependence = {
  encode(message: GetRunStatus_Dependence, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.dependencies) {
      writer.uint32(10).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetRunStatus_Dependence {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetRunStatus_Dependence();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.dependencies.push(reader.string());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetRunStatus_Dependence {
    return { dependencies: Array.isArray(object?.dependencies) ? object.dependencies.map((e: any) => String(e)) : [] };
  },

  toJSON(message: GetRunStatus_Dependence): unknown {
    const obj: any = {};
    if (message.dependencies) {
      obj.dependencies = message.dependencies.map((e) => e);
    } else {
      obj.dependencies = [];
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetRunStatus_Dependence>, I>>(base?: I): GetRunStatus_Dependence {
    return GetRunStatus_Dependence.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetRunStatus_Dependence>, I>>(object: I): GetRunStatus_Dependence {
    const message = createBaseGetRunStatus_Dependence();
    message.dependencies = object.dependencies?.map((e) => e) || [];
    return message;
  },
};

function createBaseGetRunStatus_Response(): GetRunStatus_Response {
  return { state: {}, dependencies: {} };
}

export const GetRunStatus_Response = {
  encode(message: GetRunStatus_Response, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    Object.entries(message.state).forEach(([key, value]) => {
      GetRunStatus_Response_StateEntry.encode({ key: key as any, value }, writer.uint32(10).fork()).ldelim();
    });
    Object.entries(message.dependencies).forEach(([key, value]) => {
      GetRunStatus_Response_DependenciesEntry.encode({ key: key as any, value }, writer.uint32(18).fork()).ldelim();
    });
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetRunStatus_Response {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetRunStatus_Response();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          const entry1 = GetRunStatus_Response_StateEntry.decode(reader, reader.uint32());
          if (entry1.value !== undefined) {
            message.state[entry1.key] = entry1.value;
          }
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          const entry2 = GetRunStatus_Response_DependenciesEntry.decode(reader, reader.uint32());
          if (entry2.value !== undefined) {
            message.dependencies[entry2.key] = entry2.value;
          }
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetRunStatus_Response {
    return {
      state: isObject(object.state)
        ? Object.entries(object.state).reduce<{ [key: string]: GetRunStatus_Status }>((acc, [key, value]) => {
          acc[key] = GetRunStatus_Status.fromJSON(value);
          return acc;
        }, {})
        : {},
      dependencies: isObject(object.dependencies)
        ? Object.entries(object.dependencies).reduce<{ [key: string]: GetRunStatus_Dependence }>(
          (acc, [key, value]) => {
            acc[key] = GetRunStatus_Dependence.fromJSON(value);
            return acc;
          },
          {},
        )
        : {},
    };
  },

  toJSON(message: GetRunStatus_Response): unknown {
    const obj: any = {};
    obj.state = {};
    if (message.state) {
      Object.entries(message.state).forEach(([k, v]) => {
        obj.state[k] = GetRunStatus_Status.toJSON(v);
      });
    }
    obj.dependencies = {};
    if (message.dependencies) {
      Object.entries(message.dependencies).forEach(([k, v]) => {
        obj.dependencies[k] = GetRunStatus_Dependence.toJSON(v);
      });
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetRunStatus_Response>, I>>(base?: I): GetRunStatus_Response {
    return GetRunStatus_Response.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetRunStatus_Response>, I>>(object: I): GetRunStatus_Response {
    const message = createBaseGetRunStatus_Response();
    message.state = Object.entries(object.state ?? {}).reduce<{ [key: string]: GetRunStatus_Status }>(
      (acc, [key, value]) => {
        if (value !== undefined) {
          acc[key] = GetRunStatus_Status.fromPartial(value);
        }
        return acc;
      },
      {},
    );
    message.dependencies = Object.entries(object.dependencies ?? {}).reduce<{ [key: string]: GetRunStatus_Dependence }>(
      (acc, [key, value]) => {
        if (value !== undefined) {
          acc[key] = GetRunStatus_Dependence.fromPartial(value);
        }
        return acc;
      },
      {},
    );
    return message;
  },
};

function createBaseGetRunStatus_Response_StateEntry(): GetRunStatus_Response_StateEntry {
  return { key: "", value: undefined };
}

export const GetRunStatus_Response_StateEntry = {
  encode(message: GetRunStatus_Response_StateEntry, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== undefined) {
      GetRunStatus_Status.encode(message.value, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetRunStatus_Response_StateEntry {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetRunStatus_Response_StateEntry();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.key = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.value = GetRunStatus_Status.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetRunStatus_Response_StateEntry {
    return {
      key: isSet(object.key) ? String(object.key) : "",
      value: isSet(object.value) ? GetRunStatus_Status.fromJSON(object.value) : undefined,
    };
  },

  toJSON(message: GetRunStatus_Response_StateEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value ? GetRunStatus_Status.toJSON(message.value) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<GetRunStatus_Response_StateEntry>, I>>(
    base?: I,
  ): GetRunStatus_Response_StateEntry {
    return GetRunStatus_Response_StateEntry.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetRunStatus_Response_StateEntry>, I>>(
    object: I,
  ): GetRunStatus_Response_StateEntry {
    const message = createBaseGetRunStatus_Response_StateEntry();
    message.key = object.key ?? "";
    message.value = (object.value !== undefined && object.value !== null)
      ? GetRunStatus_Status.fromPartial(object.value)
      : undefined;
    return message;
  },
};

function createBaseGetRunStatus_Response_DependenciesEntry(): GetRunStatus_Response_DependenciesEntry {
  return { key: "", value: undefined };
}

export const GetRunStatus_Response_DependenciesEntry = {
  encode(message: GetRunStatus_Response_DependenciesEntry, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== undefined) {
      GetRunStatus_Dependence.encode(message.value, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetRunStatus_Response_DependenciesEntry {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetRunStatus_Response_DependenciesEntry();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.key = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.value = GetRunStatus_Dependence.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetRunStatus_Response_DependenciesEntry {
    return {
      key: isSet(object.key) ? String(object.key) : "",
      value: isSet(object.value) ? GetRunStatus_Dependence.fromJSON(object.value) : undefined,
    };
  },

  toJSON(message: GetRunStatus_Response_DependenciesEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined &&
      (obj.value = message.value ? GetRunStatus_Dependence.toJSON(message.value) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<GetRunStatus_Response_DependenciesEntry>, I>>(
    base?: I,
  ): GetRunStatus_Response_DependenciesEntry {
    return GetRunStatus_Response_DependenciesEntry.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetRunStatus_Response_DependenciesEntry>, I>>(
    object: I,
  ): GetRunStatus_Response_DependenciesEntry {
    const message = createBaseGetRunStatus_Response_DependenciesEntry();
    message.key = object.key ?? "";
    message.value = (object.value !== undefined && object.value !== null)
      ? GetRunStatus_Dependence.fromPartial(object.value)
      : undefined;
    return message;
  },
};

export type DaemonService = typeof DaemonService;
export const DaemonService = {
  startRun: {
    path: "/Daemon/StartRun",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: StartRun_Request) => Buffer.from(StartRun_Request.encode(value).finish()),
    requestDeserialize: (value: Buffer) => StartRun_Request.decode(value),
    responseSerialize: (value: StartRun_Response) => Buffer.from(StartRun_Response.encode(value).finish()),
    responseDeserialize: (value: Buffer) => StartRun_Response.decode(value),
  },
  getLogStream: {
    path: "/Daemon/GetLogStream",
    requestStream: false,
    responseStream: true,
    requestSerialize: (value: GetLogStream_Request) => Buffer.from(GetLogStream_Request.encode(value).finish()),
    requestDeserialize: (value: Buffer) => GetLogStream_Request.decode(value),
    responseSerialize: (value: Log) => Buffer.from(Log.encode(value).finish()),
    responseDeserialize: (value: Buffer) => Log.decode(value),
  },
  getRunStatus: {
    path: "/Daemon/GetRunStatus",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: GetRunStatus_Request) => Buffer.from(GetRunStatus_Request.encode(value).finish()),
    requestDeserialize: (value: Buffer) => GetRunStatus_Request.decode(value),
    responseSerialize: (value: GetRunStatus_Response) => Buffer.from(GetRunStatus_Response.encode(value).finish()),
    responseDeserialize: (value: Buffer) => GetRunStatus_Response.decode(value),
  },
} as const;

export interface DaemonServer extends UntypedServiceImplementation {
  startRun: handleUnaryCall<StartRun_Request, StartRun_Response>;
  getLogStream: handleServerStreamingCall<GetLogStream_Request, Log>;
  getRunStatus: handleUnaryCall<GetRunStatus_Request, GetRunStatus_Response>;
}

export interface DaemonClient extends Client {
  startRun(
    request: StartRun_Request,
    callback: (error: ServiceError | null, response: StartRun_Response) => void,
  ): ClientUnaryCall;
  startRun(
    request: StartRun_Request,
    metadata: Metadata,
    callback: (error: ServiceError | null, response: StartRun_Response) => void,
  ): ClientUnaryCall;
  startRun(
    request: StartRun_Request,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (error: ServiceError | null, response: StartRun_Response) => void,
  ): ClientUnaryCall;
  getLogStream(request: GetLogStream_Request, options?: Partial<CallOptions>): ClientReadableStream<Log>;
  getLogStream(
    request: GetLogStream_Request,
    metadata?: Metadata,
    options?: Partial<CallOptions>,
  ): ClientReadableStream<Log>;
  getRunStatus(
    request: GetRunStatus_Request,
    callback: (error: ServiceError | null, response: GetRunStatus_Response) => void,
  ): ClientUnaryCall;
  getRunStatus(
    request: GetRunStatus_Request,
    metadata: Metadata,
    callback: (error: ServiceError | null, response: GetRunStatus_Response) => void,
  ): ClientUnaryCall;
  getRunStatus(
    request: GetRunStatus_Request,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (error: ServiceError | null, response: GetRunStatus_Response) => void,
  ): ClientUnaryCall;
}

export const DaemonClient = makeGenericClientConstructor(DaemonService, "Daemon") as unknown as {
  new (address: string, credentials: ChannelCredentials, options?: Partial<ClientOptions>): DaemonClient;
  service: typeof DaemonService;
};

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var tsProtoGlobalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

function bytesFromBase64(b64: string): Uint8Array {
  if (tsProtoGlobalThis.Buffer) {
    return Uint8Array.from(tsProtoGlobalThis.Buffer.from(b64, "base64"));
  } else {
    const bin = tsProtoGlobalThis.atob(b64);
    const arr = new Uint8Array(bin.length);
    for (let i = 0; i < bin.length; ++i) {
      arr[i] = bin.charCodeAt(i);
    }
    return arr;
  }
}

function base64FromBytes(arr: Uint8Array): string {
  if (tsProtoGlobalThis.Buffer) {
    return tsProtoGlobalThis.Buffer.from(arr).toString("base64");
  } else {
    const bin: string[] = [];
    arr.forEach((byte) => {
      bin.push(String.fromCharCode(byte));
    });
    return tsProtoGlobalThis.btoa(bin.join(""));
  }
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isObject(value: any): boolean {
  return typeof value === "object" && value !== null;
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
