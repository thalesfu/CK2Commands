package limbuwan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 牧龙MorongBarony struct {
	feud.BaseBarony
}

var BMorong牧龙 feud.Barony = &牧龙MorongBarony{}

func init() {
    f := BMorong牧龙.(*牧龙MorongBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "morong",
		TitleName: "牧龙",
		TitleCode: "b_morong",
	}
}
