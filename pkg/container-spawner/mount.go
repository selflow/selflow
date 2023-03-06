package container_spawner

import (
	"github.com/docker/docker/api/types/mount"
	"os"
	"path"
	"path/filepath"
)

type Mountable interface {
	ToMount() (mount.Mount, error)
}

type BinaryMount struct {
	FileContent   []byte
	Destination   string
	ReadOnly      bool
	TempDirectory string
}

func (b BinaryMount) ToMount() (mount.Mount, error) {
	tmpFile, err := os.CreateTemp("/etc/selflow", "bind-volume")
	if err != nil {
		return mount.Mount{}, err
	}

	_, err = tmpFile.Write(b.FileContent)
	if err != nil {
		return mount.Mount{}, err
	}

	return mount.Mount{
		Type:     mount.TypeBind,
		Source:   path.Join(b.TempDirectory, filepath.Base(tmpFile.Name())),
		Target:   b.Destination,
		ReadOnly: b.ReadOnly,
	}, nil
}

type FileMount struct {
	SourceFileName string
	Destination    string
	ReadOnly       bool
}

func (f FileMount) ToMount() (mount.Mount, error) {
	return mount.Mount{
		Type:     mount.TypeBind,
		Source:   f.SourceFileName,
		Target:   f.Destination,
		ReadOnly: f.ReadOnly,
	}, nil
}

func ToMountList(mountableConfigs []Mountable) ([]mount.Mount, error) {
	mounts := make([]mount.Mount, len(mountableConfigs))
	for i, mountableConfig := range mountableConfigs {
		m, err := mountableConfig.ToMount()
		if err != nil {
			return nil, err
		}
		mounts[i] = m
	}
	return mounts, nil
}
