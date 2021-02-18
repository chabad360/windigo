/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package ui

import (
	"path/filepath"
	"sort"
	"strings"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/win"
)

type _PathT struct{}

// File and folder path functions.
var Path _PathT

// Returns all the file names that match a pattern like "C:\\foo\\*.txt".
func (_PathT) ListFilesInFolder(pathAndPattern string) ([]string, error) {
	retFiles := make([]string, 0)
	dirPath := filepath.Dir(pathAndPattern) // path without file name

	wfd := win.WIN32_FIND_DATA{}
	hFind, found, err := win.FindFirstFile(pathAndPattern, &wfd)
	if err != nil {
		return nil, err
	}
	defer hFind.FindClose()

	for found {
		fileNameFound := win.Str.FromUint16Slice(wfd.CFileName[:])
		if fileNameFound != ".." {
			retFiles = append(retFiles, dirPath+"\\"+fileNameFound)
		}

		found, err = hFind.FindNextFile(&wfd)
		if err != nil {
			return nil, err
		}
	}

	sort.Slice(retFiles, func(i, j int) bool { // case insensitive
		return strings.ToUpper(retFiles[i]) < strings.ToUpper(retFiles[j])
	})
	return retFiles, nil // search finished successfully
}

// Tells if a given file or folder exists.
func (_PathT) Exists(path string) bool {
	attr, _ := win.GetFileAttributes(path)
	return attr != co.FILE_ATTRIBUTE_INVALID
}

// Tells if a given path is a folder, and not a file.
func (_PathT) IsFolder(path string) bool {
	attr, _ := win.GetFileAttributes(path)
	return attr != co.FILE_ATTRIBUTE_INVALID &&
		(attr&co.FILE_ATTRIBUTE_DIRECTORY) != 0
}

// Tells if the given file or folder is hidden.
func (_PathT) IsHidden(path string) bool {
	attr, _ := win.GetFileAttributes(path)
	return attr != co.FILE_ATTRIBUTE_INVALID &&
		(attr&co.FILE_ATTRIBUTE_HIDDEN) != 0
}
