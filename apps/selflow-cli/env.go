package main

import "github.com/selflow/selflow/libs/selflow-daemon/envutils"

const (
	//EnvTmpFileDirectory is the environment variable name for the directory where temporary
	//files like docker entrypoints are stored
	EnvTmpFileDirectory     = "SELFLOW_TEMP_FILE_DIRECTORY"
	defaultTmpFileDirectory = "/tmp"
)

var (
	tmpFileDirectory = envutils.GetEnv(EnvTmpFileDirectory, defaultTmpFileDirectory)
)
