package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type AddLicense struct {
	appDir   string
	licenses Licenses
}

type Licenses struct {
	Licenses []License `json:"licenses"`
}

type License struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Permissions []string `json:"permissions,omitempty"`
	Conditions  []string `json:"conditions,omitempty"`
	Limitations []string `json:"limitations,omitempty"`
	SpdxID      string   `json:"spdx-id"`
	LicenseFile string   `json:"license_file"`
}

func NewAddLicense() *AddLicense {
	return &AddLicense{}
}

func (a *AddLicense) ListLicenses() string {
	output := "Supported Licenses:\n\n"

	for i, license := range a.licenses.Licenses {
		output += fmt.Sprintf("%v:\n", license.Title)
		output += fmt.Sprintf(" %v\n", license.Description)
		output += fmt.Sprintf(" To Use: \n")
		output += fmt.Sprintf("   %v -l=\"%v\" \n", os.Args[0], license.ID)
		output += fmt.Sprintf("   %v -l=\"%v\"", os.Args[0], license.Title)
		if i != len(a.licenses.Licenses)-1 {
			output += fmt.Sprintf("\n\n")
		}
	}

	return output
}

func (a *AddLicense) Add(input string, dir string) error {
	dir = strings.TrimRight(dir, "/")

	if dir == "" {
		wd, err := os.Getwd()
		if err != nil {
			return err
		}

		dir = wd
	}

	dirCheck := strings.Split(dir, "/")
	if dirCheck[0] == "." {
		wd, err := os.Getwd()
		if err != nil {
			return err
		}

		dirCheck[0] = wd
	}

	if dirCheck[0] == ".." {
		wd, err := os.Getwd()
		if err != nil {
			return err
		}

		wdArr := strings.Split(wd, "/")
		wdArr = wdArr[:len(wdArr)-1]

		dirCheck[0] = strings.Join(wdArr, "/")
	}

	dir = strings.Join(dirCheck, "/")

	_, err := os.Stat(dir)
	if err != nil && os.IsNotExist(err) {
		return fmt.Errorf("Directory does not exist.")
	}

	dir = fmt.Sprintf("%s/LICENSE", dir)

	input = strings.ToLower(input)

	if input == "-o" {
		return fmt.Errorf("Nothing given to \"-l\"")
	}

	license, err := a.getLicense(input)
	if err != nil {
		return err
	}

	licenseFile := license.LicenseFile

	licensePath := fmt.Sprintf("%s/licenses/%s", a.appDir, licenseFile)

	err = copy(licensePath, dir, 1024)
	if err != nil {
		return err
	}

	return nil
}

func (a *AddLicense) getLicense(input string) (License, error) {
	if input == "" {
		return License{}, fmt.Errorf("Invalid Input.")
	}

	for i := 0; i < len(a.licenses.Licenses); i++ {
		license := a.licenses.Licenses[i]

		if license.ID == input {
			return license, nil
		}

		if strings.ToLower(license.Title) == input {
			return license, nil
		}
	}

	return License{}, fmt.Errorf("License Not Found.")
}

func (a *AddLicense) Init() error {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return err
	}

	repoUrl := "https://raw.githubusercontent.com/Th3-S1lenc3/Add-License/master"

	a.appDir = fmt.Sprintf("%s/Add-License", configDir)

	if IsNotExist(a.appDir) == true {
		err = os.Mkdir(a.appDir, os.FileMode(0755))
		if err != nil {
			return err
		}
	}

	licenseDir := fmt.Sprintf("%s/licenses", a.appDir)

	if IsNotExist(licenseDir) == true {
		err = os.Mkdir(licenseDir, os.FileMode(0755))
		if err != nil {
			return err
		}
	}

	licenseIndex := fmt.Sprintf("%s/index.json", licenseDir)

	licenseDirURL := fmt.Sprintf("%s/licenses", repoUrl)
	licenseIndexURL := fmt.Sprintf("%s/index.json", licenseDirURL)

	if IsNotExist(licenseIndex) == true {
		err = DownloadFile(licenseDir, licenseIndexURL)
		if err != nil {
			return err
		}
	}

	data, err := ioutil.ReadFile(licenseIndex)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &a.licenses)
	if err != nil {
		return err
	}

	for _, license := range a.licenses.Licenses {
		licenseFile := fmt.Sprintf("%s/%s", licenseDir, license.LicenseFile)

		if IsNotExist(licenseFile) == true {
			licenseFileURL := fmt.Sprintf("%s/%s", licenseDirURL, license.LicenseFile)
			err = DownloadFile(licenseDir, licenseFileURL)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
