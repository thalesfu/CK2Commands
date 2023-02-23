package fife

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 福克兰FalklandBarony struct {
	feud.BaseBarony
}

var BFalkland福克兰 feud.Barony = &福克兰FalklandBarony{}

func init() {
	f := BFalkland福克兰.(*福克兰FalklandBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "falkland",
		TitleName: "福克兰",
		TitleCode: "b_falkland",
	}
}
