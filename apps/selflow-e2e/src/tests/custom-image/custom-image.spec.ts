import {startRun} from "../../tools/run";
import {expect} from "vitest";
import {join} from "path"
import {parseLogs} from "../../tools/logParser";
import {matchers} from "../../tools/trace";

expect.extend(matchers);


describe('Custom Images', function () {
  it('step-python should prompt python version and step-node the node version ', async function () {
    const logs = await startRun(join(__dirname, "custom-image.yaml"))
    const trace = parseLogs(logs)

    expect(logs).not.toEqual("")

    expect(trace).toHaveStep("step-python")
    expect(trace).toHaveStep("step-node")

    expect(trace).toHaveStepTerminatedWithStatus(["step-python", "SUCCESS"])
    expect(trace).toHaveStepTerminatedWithStatus(["step-node", "SUCCESS"])

    expect(trace).toHaveStepLogged(["step-python", "Python 3"])
    expect(trace).toHaveStepLogged(["step-node", "v18"])
  });
});
