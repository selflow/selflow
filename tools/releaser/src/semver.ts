import {ReleaseType} from "./types";
import semver from "semver";

export function getNextRelease(lastTag: string, releaseType: ReleaseType) {
  // I know, v0.0.0 is not semver but starting with a number can be painful for a lot of software.
  // a 'v' won't kill us, don't worry
  return "v" + semver.inc(lastTag, releaseType, {}, 'beta')
}
