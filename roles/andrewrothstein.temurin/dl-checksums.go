package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// https://github.com/adoptium/temurin21-binaries/releases/download/jdk-21%2B35/OpenJDK21U-jre_x64_linux_hotspot_21_35.tar.gz
// https://github.com/adoptium/temurin21-binaries/releases/download/jdk-21.0.1%2B12/OpenJDK21U-jdk_x64_windows_hotspot_21.0.1_12.zip

type Ver struct {
	Major int
	Minor int
	Patch string
	BVer  string
}

func (v *Ver) Fmt() string {
	if v.Major >= 9 {
		if v.Minor == 0 && v.Patch == "" {
			// 21_35
			return fmt.Sprintf("%v_%v", v.Major, v.BVer)
		} else {
			// 21.0.1_12
			return fmt.Sprintf("%v.%v.%v_%v", v.Major, v.Minor, v.Patch, v.BVer)

		}
	} else {
		return fmt.Sprintf("%vu%vb%v", v.Major, v.Minor, v.BVer)
	}
}

func (v *Ver) lastRPath() string {
	if v.Major >= 9 {
		if v.Minor == 0 && v.Patch == ""  {
			// jdk-21%2B35
			return fmt.Sprintf(
				"jdk-%v%%2B%v",
				v.Major, v.BVer,
			)
		} else {
			// jdk-21.0.1%2B12
			return fmt.Sprintf(
				"jdk-%v.%v.%v%%2B%v",
				v.Major, v.Minor, v.Patch, v.BVer,
			)
		}
	} else {
		return fmt.Sprintf(
			"jdk%vu%v-b%v",
			v.Major, v.Minor, v.BVer,
		)
	}
}

func (v *Ver) RPath() string {
	return fmt.Sprintf(
		"temurin%d-binaries/releases/download/%s",
		v.Major, v.lastRPath(),
	)
}

type Platform struct {
	Os          string
	Arch        string
	ArchiveType string
}

func NewPlatform(os string, arch string) Platform {
	return Platform{Os: os, Arch: arch, ArchiveType: "tar.gz"}
}

func NewZipPlatform(os string, arch string) Platform {
	return Platform{Os: os, Arch: arch, ArchiveType: "zip"}
}

func (s *Platform) Fmt() string {
	return fmt.Sprintf("%s%s%s", s.Os, "_", s.Arch)
}

func (s *Platform) FmtReverse() string {
	return fmt.Sprintf("%s%s%s", s.Arch, "_", s.Os)
}

type Params struct {
	Mirror string
}

func dl_url(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", errors.New("not found")
	}
	defer resp.Body.Close()
	if b, err := io.ReadAll(resp.Body); err == nil {
		return string(b), nil
	}
	return "", err
}

func dl_checksum(checksum_url string, f string) (string, error) {
	s, err := dl_url(checksum_url)
	if err == nil {
		lines := strings.Split(s, "\n")
		for _, line := range lines {
			sums := strings.Fields(line)
			if len(sums) > 1 && strings.HasSuffix(sums[1], f) {
				return sums[0], nil
			}
		}
	}
	return "", err
}

func indent(i uint64) string {
	var b strings.Builder
	for j := uint64(0); j < i; j++ {
		fmt.Fprintf(&b, "  ")
	}
	return b.String()
}

func dl_app(
	i uint64,
	params *Params,
	app string,
	v *Ver,
	platforms []Platform,
) {
	fmt.Printf("%s%s:\n", indent(i), app)
	for _, p := range platforms {
		file := fmt.Sprintf(
			"OpenJDK%dU-%s_%s_hotspot_%s.%s",
			v.Major, app, p.FmtReverse(), v.Fmt(), p.ArchiveType,
		)
		checksumsurl := fmt.Sprintf(
			"%s/%s/%s.sha256.txt",
			params.Mirror, v.RPath(), file,
		)
		if checksum, err := dl_checksum(checksumsurl, file); err == nil {
			fmt.Printf("%s# %s\n", indent(i+1), checksumsurl)
			fmt.Printf("%s%s: sha256:%s\n", indent(i+1), p.Fmt(), checksum)
		}
	}
}

func dlall(
	i uint64,
	params *Params,
	vs []Ver,
	platforms []Platform,
) {
	for _, v := range vs {
		fmt.Printf("%s'%s':\n", indent(i), v.Fmt())
		dl_app(i+1, params, "jdk", &v, platforms)
		dl_app(i+1, params, "jre", &v, platforms)
	}
}

func main() {
	params := Params{
		Mirror: "https://github.com/adoptium",
	}

	platforms := []Platform{
		NewPlatform("aix", "ppc64"),
		NewPlatform("alpine-linux", "x64"),
		NewPlatform("linux", "s390x"),
		NewPlatform("linux", "x64"),
		NewPlatform("linux", "aarch64"),
		NewPlatform("linux", "arm"),
		NewPlatform("linux", "ppc64le"),
		NewPlatform("mac", "aarch64"),
		NewPlatform("mac", "x64"),
		NewZipPlatform("windows", "x64"),
		NewZipPlatform("windows", "x86-32"),
	}

	fmt.Println("---")
	fmt.Println("temurin_checksums:")

	// https://adoptium.net/temurin/releases/
	versions := []Ver{
		{Major: 8, Minor: 312, Patch: "0", BVer: "07"},
		{Major: 8, Minor: 322, Patch: "0", BVer: "06"},
		{Major: 8, Minor: 332, Patch: "0", BVer: "09"},
		{Major: 8, Minor: 342, Patch: "0", BVer: "07"},
		{Major: 8, Minor: 345, Patch: "0", BVer: "01"},
		{Major: 8, Minor: 352, Patch: "0", BVer: "08"},
		{Major: 8, Minor: 362, Patch: "0", BVer: "09"},
		{Major: 8, Minor: 372, Patch: "0", BVer: "07"},
		{Major: 8, Minor: 382, Patch: "0", BVer: "05"},
		{Major: 8, Minor: 392, Patch: "0", BVer: "08"},
		{Major: 11, Minor: 0, Patch: "13", BVer: "8"},
		{Major: 11, Minor: 0, Patch: "14.1", BVer: "1"},
		{Major: 11, Minor: 0, Patch: "15", BVer: "10"},
		{Major: 11, Minor: 0, Patch: "16", BVer: "8"},
		{Major: 11, Minor: 0, Patch: "16.1", BVer: "1"},
		{Major: 11, Minor: 0, Patch: "17", BVer: "8"},
		{Major: 11, Minor: 0, Patch: "18", BVer: "10"},
		{Major: 11, Minor: 0, Patch: "19", BVer: "7"},
		{Major: 11, Minor: 0, Patch: "20.1", BVer: "1"},
		{Major: 11, Minor: 0, Patch: "21", BVer: "9"},
		{Major: 16, Minor: 0, Patch: "2", BVer: "7"},
		{Major: 17, Minor: 0, Patch: "1", BVer: "12"},
		{Major: 17, Minor: 0, Patch: "2", BVer: "8"},
		{Major: 17, Minor: 0, Patch: "3", BVer: "7"},
		{Major: 17, Minor: 0, Patch: "4", BVer: "8"},
		{Major: 17, Minor: 0, Patch: "4.1", BVer: "1"},
		{Major: 17, Minor: 0, Patch: "5", BVer: "8"},
		{Major: 17, Minor: 0, Patch: "6", BVer: "10"},
		{Major: 17, Minor: 0, Patch: "7", BVer: "7"},
		{Major: 17, Minor: 0, Patch: "8.1", BVer: "1"},
		{Major: 17, Minor: 0, Patch: "9", BVer: "9"},
		{Major: 18, Minor: 0, Patch: "1", BVer: "10"},
		{Major: 18, Minor: 0, Patch: "2", BVer: "9"},
		{Major: 18, Minor: 0, Patch: "2.1", BVer: "1"},
		{Major: 19, Minor: 0, Patch: "1", BVer: "10"},
		{Major: 19, Minor: 0, Patch: "2", BVer: "7"},
		{Major: 20, Minor: 0, Patch: "1", BVer: "9"},
		{Major: 20, Minor: 0, Patch: "2", BVer: "9"},
		{Major: 21, Minor: 0, Patch: "", BVer: "35"},
		{Major: 21, Minor: 0, Patch: "1", BVer: "12"},
	}
	dlall(1, &params, versions, platforms)
}
