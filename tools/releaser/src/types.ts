export type MainReleaseType = 'major' | 'minor' | 'patch'
type PreReleaseType = `prerelease`
export type ReleaseType = MainReleaseType | PreReleaseType
export type Commit = {
  message: string,
  hash: string,
  smallHash: string,
  category?: string,
  messageToDisplay?: string
  affectedProjects?: string[]
}
