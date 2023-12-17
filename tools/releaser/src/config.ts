import {MainReleaseType} from "./types";

export const config = {
  major: [], // Major release are disabled until the project exists Beta
  minor: [
    ":boom:",
    ":sparkles:"
  ],
  patch: [
    ":bug:",
    ":ambulance:",
    ":lock:",
    ":lipstick:",
    ":zap:",
    ":globe_with_meridians:",
    ":alien:",
    ":wheelchair:",
    ":loud_sound:",
    ":mute:",
    ":children_crossing:",
    ":speech_balloon:",
    ":iphone:",
    ":pencil2:",
    ":bento:",
    ":green_apple:",
    ":green_heart:",
  ],
} as const

export const releasePower = {
  'major': 2,
  'minor': 1,
  'patch': 0,
} as const satisfies  Record<MainReleaseType, number>

export const gitConfig = {
  username: "Selflow",
  email: "selflow@users.noreply.github.com",
  remoteName: 'origin',
  releaseBranch: 'main'
}

export const githubConfig = {
  repoOwner: "selflow",
  repoName: "selflow-sand",

  token: process.env.GITHUB_TOKEN,
}
