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

package era

import (
	"regexp"

	"github.com/renom/go-wesnoth/wesnoth"
	"github.com/renom/go-wml"
)

type Era struct {
	Id       string
	Name     string
	Body     string
	Factions []wml.Data
	Events   []wml.RawData
}

var ErasPath = "/usr/share/wesnoth/data/multiplayer/eras.cfg"
var eras []byte

func Parse(id string) Era {
	if len(eras) == 0 {
		eras = wesnoth.Preprocess(ErasPath, nil)
	}
	e, _ := regexp.Compile(`(?U)\[era\]\n(?:[^\[\]]*\n)*\tid="era_` + id + `"\n(?:.*\n)*\tname=_?"(.*)"\n(?:.*\n)*\[/era\]`)
	body := string(e.Find(eras)) + "\n"
	name := e.FindStringSubmatch(body)[1]

	r, _ := regexp.Compile(`(?U)\[multiplayer_side\](.*\n)*[\t ]*\[/multiplayer_side\]`)
	f := r.FindAll([]byte(body), -1)
	rData, _ := regexp.Compile(`(?U)[\t ]*[0-9a-z_]+[\t ]*=[\t ]*_?"[^"](.|\n)*` + `([^"]"[\t\n ]*\+[\t\n ]*_?"[^"])+` +
		`(.|\n)*[^"]"\n`)
	factions := []wml.Data{}
	for i, v := range f {
		f[i] = rData.ReplaceAll(v, []byte(""))
		factionData := wml.ParseTag(string(f[i])).Data
		if !factionData.Contains("random_faction") || factionData["random_faction"] == false {
			factions = append(factions, factionData)
		}
	}

	rEvents, _ := regexp.Compile(`(?U)\[event\]\n((?:.*\n)*)[\t ]*\[/event\]`)

	events := []wml.RawData{}
	for _, v := range rEvents.FindAllStringSubmatch(body, -1) {
		events = append(events, wml.RawData(v[1]))
	}

	return Era{id, name, body, factions, events}
}
