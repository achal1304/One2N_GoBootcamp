package handler

import (
	"fmt"

	"github.com/achal1304/One2N_GoBootcamp/tree/contract"
	"github.com/achal1304/One2N_GoBootcamp/tree/utils"
)

func ProcessTreeRequest(req contract.TreeRequest) (contract.TreeResponse, error) {
	searchResponse := contract.TreeResponse{}
	var currDirName string
	if req.FolderName != "" {
		isDir := utils.CheckDirectory(req.FolderName)
		fmt.Println("isdir is", isDir, req.FolderName)
		if !isDir {
			return contract.TreeResponse{}, fmt.Errorf("%s [error opening dir]", req.FolderName)
		}
		currDirName = req.FolderName
	} else {
		currentDir, err := utils.GetCurrentDir()
		if err != nil {
			return contract.TreeResponse{}, err
		}
		currDirName = currentDir
	}

	req.FolderName = currDirName
	ProcessDirectory(req, &searchResponse)
	return searchResponse, nil
}
