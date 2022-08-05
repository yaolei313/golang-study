package os_study

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/fs"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"regexp"
)

func CleanMavenRepository() {
	pwd, _ := os.Getwd()
	log.Infof("working directory %s", pwd)
	u, err := user.Current()
	if err != nil {
		log.Fatal("current error %t", err)
	}
	fmt.Println("user home dir:", u.HomeDir)
	// /Users/yaolei/.m2/repository/com/thoughtworks
	dirname := path.Join(u.HomeDir, "/.m2/repository/com/thoughtworks")
	// fileList, err := ioutil.ReadDir(dirname)
	doCleanVersionDir(dirname)
}

func doCleanVersionDir(dirName string) int {
	versionMap := make(map[string][]string)
	filepath.WalkDir(dirName, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Errorf("error %+v", err)
			return err
		}
		if d == nil {
			return nil
		}
		if d.Name() == ".DS_Store" {
			return nil
		}

		if !d.IsDir() {
			return nil
		}

		parentDir := filepath.Dir(path)
		// log.Infof("----%s %s", path, d.Name())
		if isVersionDir(d.Name()) {
			versionSlice := versionMap[parentDir]
			newVersionSlice := append(versionSlice, d.Name())
			versionMap[parentDir] = newVersionSlice
			// 到version这一层就可以了，不处理内部
			return fs.SkipDir
		}
		return nil
	})

	log.Infof("%s have %v", dirName, versionMap)
	vlen := len(versionMap)
	if vlen > 2 {
		// 保留2个最新的version目录
		needDeleteList := versionList[:vlen-2]
		/*for idx, _ := range needDeleteList {
			deleteDir := path.Join(rootDir, needDeleteList[idx])
			os.RemoveAll(deleteDir)
		}*/
		log.Infof("delete dir %v", needDeleteList)
	}

	return vlen - 2
}

var reg1 = regexp.MustCompile(`^(?P<major>0|[1-9]\d*)\.(?P<minor>0|[1-9]\d*)\.(?P<patch>0|[1-9]\d*)(?:-(?P<prerelease>(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+(?P<buildmetadata>[0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`)

func isVersionDir(name string) bool {
	return reg1.MatchString(name)
}
