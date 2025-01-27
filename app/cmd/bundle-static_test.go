package cmd

import (
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/ngocphuongnb/tetua/app/asset"
	"github.com/ngocphuongnb/tetua/app/config"
	"github.com/stretchr/testify/assert"
)

func fileMTime(name string) (mtime time.Time, err error) {
	file, err := os.Stat(name)
	if err != nil {
		return
	}
	return file.ModTime(), nil
}

func readFile(name string) string {
	b, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func TestBundleAssets(t *testing.T) {
	themeDir := "../themes/default"
	config.ROOT_DIR = "./someinvaliddir"
	asset.Load(themeDir, true)
	assert.Equal(t, true, strings.HasPrefix(BundleStaticAssets().Error(), "open someinvaliddir/"))

	bundledCssMTime, _ := fileMTime("../asset/bundled.css.go")
	bundledJsMTime, _ := fileMTime("../asset/bundled.js.go")
	bundledOtherMTime, _ := fileMTime("../asset/bundled.js.go")

	bundledCssContent := readFile("../asset/bundled.css.go")
	bundledJsContent := readFile("../asset/bundled.js.go")
	bundledOtherContent := readFile("../asset/bundled.other.go")

	_, testFile, _, _ := runtime.Caller(0)
	config.ROOT_DIR = path.Join(path.Dir(testFile), "../../")

	themeDir = path.Join(config.ROOT_DIR, "app/themes/default")
	config.WD = config.ROOT_DIR

	asset.Load(themeDir, true)
	assert.Equal(t, nil, BundleStaticAssets())

	newBundledCssMTime, _ := fileMTime(config.ROOT_DIR + "/app/asset/bundled.css.go")
	newBundledJsMTime, _ := fileMTime(config.ROOT_DIR + "/app/asset/bundled.js.go")
	newBundledOtherMTime, _ := fileMTime(config.ROOT_DIR + "/app/asset/bundled.js.go")

	newBundledCssContent := readFile(config.ROOT_DIR + "/app/asset/bundled.css.go")
	newBundledJsContent := readFile(config.ROOT_DIR + "/app/asset/bundled.js.go")
	newBundledOtherContent := readFile(config.ROOT_DIR + "/app/asset/bundled.other.go")

	assert.Equal(t, true, bundledCssMTime.Before(newBundledCssMTime))
	assert.Equal(t, true, bundledJsMTime.Before(newBundledJsMTime))
	assert.Equal(t, true, bundledOtherMTime.Before(newBundledOtherMTime))

	assert.Equal(t, bundledCssContent, newBundledCssContent)
	assert.Equal(t, bundledJsContent, newBundledJsContent)
	assert.Equal(t, bundledOtherContent, newBundledOtherContent)

	// Test minify js file error
	invalidJsFileContent := `alert(;`
	invalidJsFilePath, err := ioutil.TempFile(config.ROOT_DIR+"/private/tmp", "jsfile-")

	if err != nil {
		panic(err)
	}

	defer os.RemoveAll(invalidJsFilePath.Name())

	if err := os.WriteFile(invalidJsFilePath.Name(), []byte(invalidJsFileContent), 0644); err != nil {
		panic(err)
	}

	asset.AppendAsset(&asset.StaticAsset{
		Name: "invalidjsfile",
		Path: invalidJsFilePath.Name(),
		Type: "js",
	})

	errContent := BundleStaticAssets().Error()
	assert.Equal(t, true, strings.HasPrefix(errContent, "unexpected"))
	assert.Equal(t, true, strings.Contains(errContent, "alert"))

	asset.Load(themeDir, true)
	cssBundledContent = "package asset\npackage asset2"
	assert.Equal(t, "2:1: expected declaration, found 'package'", bundleAssets().Error())

	reset()
	jsBundledContent = "package asset\npackage asset2"
	assert.Equal(t, "2:1: expected declaration, found 'package'", bundleAssets().Error())

	reset()
	otherBundledContent = "package asset\npackage asset2"
	assert.Equal(t, "2:1: expected declaration, found 'package'", bundleAssets().Error())
}

func TestWriteBundledAssetsError(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	config.ROOT_DIR = path.Join(path.Dir(testFile), "../../")
	config.WD = config.ROOT_DIR
	themeDir := config.ROOT_DIR + "/app/themes/default"
	asset.Load(themeDir, true)

	reset()
	cssBundledFile = "./invaliddir/invalidFile"
	assert.Equal(t, "open ./invaliddir/invalidFile: no such file or directory", bundleAssets().Error())

	reset()
	jsBundledFile = "./invaliddir/invalidFile"
	assert.Equal(t, "open ./invaliddir/invalidFile: no such file or directory", bundleAssets().Error())

	reset()
	otherBundledFile = "./invaliddir/invalidFile"
	assert.Equal(t, "open ./invaliddir/invalidFile: no such file or directory", bundleAssets().Error())
}
