package orava

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佐约姆ZolyomBarony struct {
	feud.BaseBarony
}

var BZolyom佐约姆 feud.Barony = &佐约姆ZolyomBarony{}

func init() {
	f := BZolyom佐约姆.(*佐约姆ZolyomBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zolyom",
		TitleName: "佐约姆",
		TitleCode: "b_zolyom",
	}
}
