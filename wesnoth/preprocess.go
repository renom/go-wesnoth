// This file is part of Go Wesnoth.
//
// Go Wesnoth is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Go Wesnoth is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Go Wesnoth.  If not, see <https://www.gnu.org/licenses/>.

package wesnoth

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var (
	Wesnoth = "/usr/bin/wesnoth"
	Output  = os.TempDir() + "/go-wesnoth/output"
)

func Preprocess(filePath string, defines []string) []byte {
	defines = append(defines, "MULTIPLAYER")
	if _, err := os.Stat(Output); os.IsNotExist(err) {
		os.MkdirAll(Output, 0755)
	}
	cmd := exec.Command(
		Wesnoth,
		"-p",
		filePath,
		Output,
		"--preprocess-defines="+strings.Join(defines, ","),
	)
	cmd.Run()
	result, _ := ioutil.ReadFile(Output + "/" + filepath.Base(filePath))
	return result
}
