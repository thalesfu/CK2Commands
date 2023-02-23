package vaspurakan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏尔布哈奇SurbkhachBarony struct {
	feud.BaseBarony
}

var BSurbkhach苏尔布哈奇 feud.Barony = &苏尔布哈奇SurbkhachBarony{}

func init() {
	f := BSurbkhach苏尔布哈奇.(*苏尔布哈奇SurbkhachBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "surbkhach",
		TitleName: "苏尔布哈奇",
		TitleCode: "b_surbkhach",
	}
}
