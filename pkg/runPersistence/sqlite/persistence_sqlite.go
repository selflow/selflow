package sqlite

import (
	"database/sql"
	"encoding/json"
	_ "github.com/mattn/go-sqlite3"
	"github.com/selflow/selflow/internal/config"
	"github.com/selflow/selflow/pkg/workflow"
)

type StepDbRecord struct {
	id         string
	runId      string
	statusName string
}

const CreateRunTableQuery = `
CREATE TABLE IF NOT EXISTS t_run
(
    id         varchar primary key,
    start_time text,
    end_time   text
);
`

const CreateStatusTableQuery = `
  CREATE TABLE IF NOT EXISTS t_status (
      code int primary key,
      name varchar,
      is_cancellable int,
      is_finished int
  );
`

const CreateStepTableQuery = `
  CREATE TABLE IF NOT EXISTS t_step
  (
      id          varchar,
      run_id      varchar references t_run (id),
      status      int references t_status (code),
      json_schema text,
      primary key (id, run_id)
  );
`

const CreateDependenceTableQuery = `
  CREATE TABLE IF NOT EXISTS t_dependence
  (
      run_id     varchar,
      step_id    varchar,
      depends_on varchar,
      primary key (run_id, step_id, depends_on)
  );
`

const UpsertDependenceQuery = `
  INSERT INTO t_dependence (run_id, step_id, depends_on) VALUES ($1, $2, $3)
    on conflict do nothing;
`

const SetStepStatus = `
  INSERT INTO t_step (id, status, run_id) VALUES ($1, $2, $3)
    on conflict do update set status = $2;
`

const InsertStepQuery = `
 INSERT INTO t_step (run_id, id, json_schema) VALUES ($1, $2, $3);
`

const UpsertStatusQuery = `
  INSERT INTO t_status (code, name, is_cancellable, is_finished) VALUES ($1, $2, $3, $4)
    on conflict do nothing;
`

const GetRunStateQuery = `
  SELECT
      step.id,
      status.code,
      status.name,
      status.is_cancellable,
      status.is_finished
  FROM t_step step left join t_status status on status.code = step.status WHERE run_id = $1;
`

const GetRunStepDefinitionsQuery = `
  SELECT
      id,
      json_schema
  FROM t_step step WHERE run_id = $1;
`

const GetDependenciesQuery = `
  SELECT
      STEP_ID, DEPENDS_ON
  FROM t_dependence WHERE run_id = $1;
`

type RunPersistence struct {
	db *sql.DB
}

func NewSqliteRunPersistence(fileName string) (*RunPersistence, error) {
	db, err := sql.Open("sqlite3", fileName)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(CreateRunTableQuery)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(CreateStatusTableQuery)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(CreateStepTableQuery)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(CreateDependenceTableQuery)
	if err != nil {
		return nil, err
	}

	return &RunPersistence{db}, nil
}

func (rp *RunPersistence) upsertStatus(status workflow.Status) error {
	_, err := rp.db.Exec(UpsertStatusQuery, status.GetCode(), status.GetName(), status.IsCancellable(), status.IsFinished())
	return err
}

func (rp *RunPersistence) upsertDependence(runId string, stepId string, dependenceId string) error {
	_, err := rp.db.Exec(UpsertDependenceQuery, runId, stepId, dependenceId)
	return err
}

func (rp *RunPersistence) upsertStep(runId string, stepId string, status workflow.Status) error {

	err := rp.upsertStatus(status)
	if err != nil {
		return nil
	}

	_, err = rp.db.Exec(SetStepStatus, stepId, status.GetCode(), runId)
	return err
}

func (rp *RunPersistence) SetRunState(runId string, state map[string]workflow.Status) error {
	for stepId, status := range state {
		err := rp.upsertStep(runId, stepId, status)
		if err != nil {
			return err
		}
	}

	return nil
}

func (rp *RunPersistence) SetDependenciesState(runId string, dependencies map[workflow.Step][]workflow.Step) error {
	for step, stepDependencies := range dependencies {
		for _, dependency := range stepDependencies {
			err := rp.upsertDependence(runId, step.GetId(), dependency.GetId())
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (rp *RunPersistence) SetRunStepDefinitions(runId string, stepDefinitions map[string]config.StepDefinition) error {

	for stepId, stepDefinition := range stepDefinitions {

		definitionAsJson, err := json.Marshal(stepDefinition)
		if err != nil {
			return err
		}

		_, err = rp.db.Exec(InsertStepQuery, runId, stepId, string(definitionAsJson))
		if err != nil {
			return nil
		}
	}

	return nil
}

func (rp *RunPersistence) GetRunState(runId string) (map[string]workflow.Status, error) {
	rows, err := rp.db.Query(GetRunStateQuery, runId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	state := map[string]workflow.Status{}

	for rows.Next() {
		var id string
		var statusCode uint
		var statusName string
		var statusIsCancellable int
		var statusIsFinished int
		if err := rows.Scan(&id, &statusCode, &statusName, &statusIsCancellable, &statusIsFinished); err != nil {
			return nil, err
		}

		state[id] = workflow.SimpleStatus{
			Code:        statusCode,
			Name:        statusName,
			Finished:    statusIsFinished != 0,
			Cancellable: statusIsCancellable != 0,
		}
	}

	return state, nil
}

func (rp *RunPersistence) GetRunDependencies(runId string) (map[string][]string, error) {
	rows, err := rp.db.Query(GetDependenciesQuery, runId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	dependencies := map[string][]string{}

	for rows.Next() {
		var stepId string
		var dependenceId string
		if err := rows.Scan(&stepId, &dependenceId); err != nil {
			return nil, err
		}

		dependencies[stepId] = append(dependencies[stepId], dependenceId)
	}

	return dependencies, nil
}

func (rp *RunPersistence) GetRunStepDefinitions(runId string) (map[string][]byte, error) {
	rows, err := rp.db.Query(GetRunStepDefinitionsQuery, runId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	stepDefinitions := map[string][]byte{}

	for rows.Next() {
		var stepId string
		var definitionAsString string
		if err := rows.Scan(&stepId, &definitionAsString); err != nil {
			return nil, err
		}

		stepDefinitions[stepId] = []byte(definitionAsString)
	}

	return stepDefinitions, nil
}
