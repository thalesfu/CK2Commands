package tummana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 毗恃VijBarony struct {
	feud.BaseBarony
}

var BVij毗恃 feud.Barony = &毗恃VijBarony{}

func init() {
	f := BVij毗恃.(*毗恃VijBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vij",
		TitleName: "毗恃",
		TitleCode: "b_vij",
	}
}
