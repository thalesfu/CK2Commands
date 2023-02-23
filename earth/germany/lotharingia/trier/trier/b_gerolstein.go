package trier

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 盖罗尔施泰因GerolsteinBarony struct {
	feud.BaseBarony
}

var BGerolstein盖罗尔施泰因 feud.Barony = &盖罗尔施泰因GerolsteinBarony{}

func init() {
	f := BGerolstein盖罗尔施泰因.(*盖罗尔施泰因GerolsteinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gerolstein",
		TitleName: "盖罗尔施泰因",
		TitleCode: "b_gerolstein",
	}
}
