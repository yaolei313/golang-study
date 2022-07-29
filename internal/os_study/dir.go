package os_study

import (
	log "github.com/sirupsen/logrus"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
)

func CleanMavenRepository() {
	pwd, _ := os.Getwd()
	log.Infof("working directory %s", pwd)
	dirname := "/Users/yaolei/.m2/repository/kuaishou"
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
		if count > 1 {
			break
		}
	}

}

func doCleanVersionDir(dirName string) int {
	var versionList []string
	filepath.WalkDir(dirName, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Errorf("error %+v", err)
			return err
		}
		if d == nil {
			return nil
		}
		if !d.IsDir() {
			return fs.SkipDir
		}
		log.Infof("----%s %s", path, d.Name())
		versionList = append(versionList, d.Name())
		return nil
	})

	log.Infof("--%s have %v", dirName, versionList)

	return 0
}
