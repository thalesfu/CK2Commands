package qohistan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图尔希兹TorshizBarony struct {
	feud.BaseBarony
}

var BTorshiz图尔希兹 feud.Barony = &图尔希兹TorshizBarony{}

func init() {
    f := BTorshiz图尔希兹.(*图尔希兹TorshizBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "torshiz",
		TitleName: "图尔希兹",
		TitleCode: "b_torshiz",
	}
}
