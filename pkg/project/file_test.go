package project_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/wakatime/wakatime-cli/pkg/project"
	"github.com/yookoala/realpath"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFile_Detect_FileExists(t *testing.T) {
	rp, err := realpath.Realpath("testdata/.wakatime-project")
	require.NoError(t, err)

	f := project.File{
		Filepath: rp,
	}

	result, detected, err := f.Detect()
	require.NoError(t, err)

	expected := project.Result{
		Branch:  "master",
		Folder:  filepath.Dir(rp),
		Project: "wakatime-cli",
	}

	assert.True(t, detected)
	assert.Equal(t, expected, result)
}

func TestFile_Detect_ParentFolderExists(t *testing.T) {
	tmpDir, err := os.MkdirTemp(os.TempDir(), "wakatime")
	require.NoError(t, err)

	defer os.RemoveAll(tmpDir)

	tmpDir, err = realpath.Realpath(tmpDir)
	require.NoError(t, err)

	dir := filepath.Join(tmpDir, "src", "otherfolder")

	err = os.MkdirAll(dir, os.FileMode(int(0700)))
	require.NoError(t, err)

	copyFile(
		t,
		"testdata/.wakatime-project",
		filepath.Join(tmpDir, ".wakatime-project"),
	)

	f := project.File{
		Filepath: dir,
	}

	result, detected, err := f.Detect()
	require.NoError(t, err)

	expected := project.Result{
		Branch:  "master",
		Folder:  tmpDir,
		Project: "wakatime-cli",
	}

	assert.True(t, detected)
	assert.Equal(t, expected, result)
}

func TestFile_Detect_AnyFileFound(t *testing.T) {
	tmpFile, err := os.CreateTemp(os.TempDir(), "wakatime-project")
	require.NoError(t, err)

	defer os.Remove(tmpFile.Name())

	f := project.File{
		Filepath: os.TempDir(),
	}

	result, detected, err := f.Detect()
	require.NoError(t, err)

	expected := project.Result{}

	assert.False(t, detected)
	assert.Equal(t, expected, result)
}

func TestFile_Detect_InvalidPath(t *testing.T) {
	tmpFile, err := os.CreateTemp(os.TempDir(), "non-valid-file")
	require.NoError(t, err)

	defer os.Remove(tmpFile.Name())

	f := project.File{
		Filepath: tmpFile.Name(),
	}

	_, detected, err := f.Detect()
	require.NoError(t, err)

	assert.False(t, detected)
}

func TestFindFileOrDirectory(t *testing.T) {
	tmpDir, err := os.MkdirTemp(os.TempDir(), "wakatime")
	require.NoError(t, err)

	defer os.RemoveAll(tmpDir)

	dir := filepath.Join(tmpDir, "src", "otherfolder")

	err = os.MkdirAll(dir, os.FileMode(int(0700)))
	require.NoError(t, err)

	copyFile(
		t,
		"testdata/.wakatime-project",
		filepath.Join(tmpDir, ".wakatime-project"),
	)

	tests := map[string]struct {
		Filepath string
		Filename string
		Expected string
	}{
		"filename": {
			Filepath: dir,
			Filename: ".wakatime-project",
			Expected: filepath.Join(tmpDir, ".wakatime-project"),
		},
		"directory": {
			Filepath: dir,
			Filename: "src",
			Expected: filepath.Join(tmpDir, "src"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			fp, ok := project.FindFileOrDirectory(test.Filepath, test.Filename)
			require.True(t, ok)

			assert.Equal(t, test.Expected, fp)
		})
	}
}

func TestFile_String(t *testing.T) {
	f := project.File{}

	assert.Equal(t, "project-file-detector", f.String())
}

func copyFile(t *testing.T, source, destination string) {
	input, err := os.ReadFile(source)
	require.NoError(t, err)

	err = os.WriteFile(destination, input, 0600)
	require.NoError(t, err)
}
