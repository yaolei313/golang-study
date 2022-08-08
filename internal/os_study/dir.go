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
	dirname := path.Join(u.HomeDir, "/.m2/repository/com/")
	// fileList, err := ioutil.ReadDir(dirname)
	walkAndCleanDir(dirname)
}

func walkAndCleanDir(dirName string) {
	var versionSlice []string
	currentParentDir := ""
	count := 0
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

		// log.Infof("----%s %s", path, d.Name())
		if isVersionDir(d.Name()) {
			parentDir := filepath.Dir(path)
			if parentDir != currentParentDir {
				// 清理已扫描除的version目录
				deleteDuplicatedVersion(currentParentDir, versionSlice)
				count++
				versionSlice = nil
				currentParentDir = parentDir
			}
			versionSlice = append(versionSlice, d.Name())
		}
		return nil
	})
	log.Infof("handle count %v", count)
}

func deleteDuplicatedVersion(dir string, slice []string) {
	log.Infof("dir %s has %v", dir, slice)
	vlen := len(slice)
	if vlen > 2 {
		// 保留2个最新的version目录
		needDeleteList := slice[:vlen-2]
		/*for idx, _ := range needDeleteList {
			deleteDir := path.Join(rootDir, needDeleteList[idx])
			os.RemoveAll(deleteDir)
		}*/
		log.Infof("delete dir %v", needDeleteList)
	}
}

var reg1 = regexp.MustCompile(`^(?P<major>0|[1-9]\d*)\.(?P<minor>0|[1-9]\d*)\.(?P<patch>0|[1-9]\d*)(?:-(?P<prerelease>(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+(?P<buildmetadata>[0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`)

func isVersionDir(name string) bool {
	return reg1.MatchString(name)
}
