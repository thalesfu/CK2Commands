package orbetello

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥尔贝泰洛OrbetelloBarony struct {
	feud.BaseBarony
}

var BOrbetello奥尔贝泰洛 feud.Barony = &奥尔贝泰洛OrbetelloBarony{}

func init() {
	f := BOrbetello奥尔贝泰洛.(*奥尔贝泰洛OrbetelloBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "orbetello",
		TitleName: "奥尔贝泰洛",
		TitleCode: "b_orbetello",
	}
}
