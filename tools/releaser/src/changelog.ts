import {Commit} from "./types";
import {readFileSync} from "fs";
import {join} from "path";
import handlebars from "handlebars";
import * as process from "process";

export async function buildChangelog(commits: Commit[], nextRelease: string) {
  const templateAsString = readFileSync(join(__dirname, "..", "..", "..", "assets", "release-note.hbs")).toString()
  handlebars.registerPartial('commitTemplate', '[{{smallHash}}](https://github.com/selflow/selflow/commit/{{hash}}) {{messageToDisplay}} > {{#each affectedProjects}}*{{.}}* {{/each}}')
  const template = handlebars.compile(templateAsString)

  const dateformat = await import("dateformat")

  const templateContext = {
    commits: commits.reduce((acc, c) => ({
      ...acc,
      [c.category]: [...(acc[c.category] ?? []), c]
    }), {}),
    nextRelease,
    releaseDate: dateformat.default(new Date(), "UTC:yyyy-mm-dd")
  }

  const changelog = template(templateContext)

  return changelog.replace(/([^\n])\n\n\n+([^\n])/, '$1\n\n$2').trimEnd()

}
