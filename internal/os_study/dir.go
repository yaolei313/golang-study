package os_study

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/fs"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
)

func CleanMavenRepository() {
	pwd, _ := os.Getwd()
	log.Infof("working directory %s", pwd)
	u, err := user.Current()
	if err != nil {
		log.Fatal("current error %t", err)
	}
	fmt.Println("user home dir:", u.HomeDir)
	dirname := u.HomeDir + "/.m2/repository/kuaishou"
	fileList, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Warnf("read dir fail %s %t", dirname, err)
	}
	log.Infof("dir file count: %d", len(fileList))
	count := 0
	for i, _ := range fileList {
		tmpFile := fileList[i]
		if tmpFile.Name() == ".DS_Store" {
			continue
		}
		if tmpFile.IsDir() {
			log.Infof("handle %s", tmpFile.Name())
			doCleanVersionDir(dirname + "/" + tmpFile.Name())
			count += 1
		}
		if count > 30 {
			break
		}
	}

}

func doCleanVersionDir(dirName string) int {
	rootDir := dirName
	var versionList []string
	filepath.WalkDir(dirName, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Errorf("error %+v", err)
			return err
		}
		if d == nil {
			return nil
		}
		if path == rootDir {
			// 继续处理
			return nil
		}
		if !d.IsDir() {
			return fs.SkipDir
		}

		// log.Infof("----%s %s", path, d.Name())
		versionList = append(versionList, d.Name())

		// 不处理内部目录
		return fs.SkipDir
	})

	log.Infof("--%s have %v", dirName, versionList)

	return 0
}
