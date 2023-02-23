package vijnot

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 毗者努多VijnotBarony struct {
	feud.BaseBarony
}

var BVijnot毗者努多 feud.Barony = &毗者努多VijnotBarony{}

func init() {
	f := BVijnot毗者努多.(*毗者努多VijnotBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vijnot",
		TitleName: "毗者努多",
		TitleCode: "b_vijnot",
	}
}
