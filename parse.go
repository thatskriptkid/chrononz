package main

import (
	"strings"
	"fmt"
	gore "github.com/goretk/gore"
	"time"
	"errors"
)

var goVerIncompatible = "incompatible" 
var noVersionStr = "v0.0.0"
var githubPrefix = "github.com"

type VendorInfo struct {
	PkgName string
	Date time.Time
} 

// strings like v0.0.0-20131221200532-179d4d0c4d8d 
// mean that repo doesnt have releases. We should 
// take 20131221 as a date
func parseNoVerStr(ver string) (time.Time, error) {
	
	var releaseDateStr string = (strings.Split(ver, "-")[1])[0:8]
	if (len(releaseDateStr) != 0) {
		layout := "20060102"
		t, err := time.Parse(layout, releaseDateStr)

		if err != nil {
			return time.Time{}, err
		}
		return t, nil
	}
	return time.Time{}, errors.New("сouldnt get date")
}

func getVersion(filePath string) string {
	i := strings.LastIndex(filePath, "@")
	if i != -1 && i != len(filePath)+1 {
		return strings.Split(filePath[i+1:], "/")[0]
	}
	return ""
}

func GetVendorsInfo(pkgs []*gore.Package) ([]VendorInfo, error) {

	var vsInfo []VendorInfo
	
	for _, p := range pkgs { 

		var vInfo VendorInfo

		ver := getVersion(p.Filepath)

		if len(ver) == 0 || ver == "" {
			return nil, errors.New("failed to get version")
		}

		if strings.HasPrefix(p.Name, githubPrefix) {

			if strings.Contains(ver, goVerIncompatible) {
				ver = strings.Split(ver, "+")[0]
			}
			
			if strings.Contains(ver, noVersionStr) {
				t, err := parseNoVerStr(ver)
				if (err != nil) {
					fmt.Println(err)
					continue
				}
				vInfo = VendorInfo{p.Name, t}
				vsInfo = append(vsInfo, vInfo)
				continue
			}
			
			// get date using GitHub API
			t, err := ResolveReleaseDate(p.Name, ver)
			if err != nil {
				fmt.Println(err)
				continue
			}
			vInfo = VendorInfo{p.Name, t}
			vsInfo = append(vsInfo, vInfo)
			
		} else {
			t, err := parseNoVerStr(ver)
			if (err != nil) {
				fmt.Println(err)
				continue
			}
			vInfo = VendorInfo{p.Name, t}
			vsInfo = append(vsInfo, vInfo)
		}
	}
	return vsInfo, nil
}