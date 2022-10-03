package folderapi

import (
	"github.com/lowl11/lazyfile/filemodels"
	"io/ioutil"
	"path/filepath"
	"sort"
)

func GetObjects(path string) ([]filemodels.Object, error) {
	objectList := make([]filemodels.Object, 0)
	folderObjects, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, objectItem := range folderObjects {
		objectName := objectItem.Name()
		isFolder := objectItem.IsDir()
		objectPath := buildObjectPath(path, objectName)

		objectList = append(objectList, filemodels.Object{
			Name:        objectName,
			Path:        objectPath,
			IsFolder:    isFolder,
			ObjectCount: GetCount(objectPath),
		})
	}

	// sort by folders & files
	sort.Slice(objectList, func(i, j int) bool {
		return objectList[i].IsFolder
	})

	// sort folder by alphabet
	sort.Slice(objectList, func(i, j int) bool {
		return (objectList[i].Name < objectList[j].Name) && (objectList[i].IsFolder && objectList[j].IsFolder)
	})

	// sort files by alphabet
	sort.Slice(objectList, func(i, j int) bool {
		return (objectList[i].Name < objectList[j].Name) && (!objectList[i].IsFolder && !objectList[j].IsFolder)
	})

	return objectList, nil
}

func GetObjectsWithDepth(path, memoryPath string) ([]filemodels.Object, error) {
	objectList := make([]filemodels.Object, 0)
	folderObjects, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	relativeRootPath := filepath.Dir(path)
	if memoryPath == "" {
		memoryPath = relativeRootPath
	}

	for _, objectItem := range folderObjects {
		// main meta info
		objectName := objectItem.Name()
		isFolder := objectItem.IsDir()
		objectPath := buildObjectPath(path, objectName)
		objectCount := GetCount(objectPath)

		// getting children
		children := make([]filemodels.Object, 0, objectCount)
		children, err = GetObjectsWithDepth(objectPath, memoryPath)
		if err != nil {
			children = make([]filemodels.Object, 0, objectCount)
		}

		// memory path
		//var nextPath string
		//if isFolder {
		//	nextPath = objectName
		//}
		//memoryPath = fmt.Sprintf("%s/%s", memoryPath, nextPath)
		//objectMemoryPath := buildMemoryObjectPath(memoryPath, objectName)
		objectMemoryPath := memoryPath

		objectList = append(objectList, filemodels.Object{
			Name:         objectName,
			Path:         objectPath,
			RelativePath: objectMemoryPath,
			IsFolder:     isFolder,
			ObjectCount:  objectCount,
			Children:     children,
		})
	}

	// sort by folders & files
	sort.Slice(objectList, func(i, j int) bool {
		return objectList[i].IsFolder
	})

	// sort folder by alphabet
	sort.Slice(objectList, func(i, j int) bool {
		return (objectList[i].Name < objectList[j].Name) && (objectList[i].IsFolder && objectList[j].IsFolder)
	})

	// sort files by alphabet
	sort.Slice(objectList, func(i, j int) bool {
		return (objectList[i].Name < objectList[j].Name) && (!objectList[i].IsFolder && !objectList[j].IsFolder)
	})

	return objectList, nil
}

func GetCount(path string) int {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return 0
	}

	return len(files)
}
