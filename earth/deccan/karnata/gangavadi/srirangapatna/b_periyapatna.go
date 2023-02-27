package srirangapatna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佩里亚帕提纳PeriyapatnaBarony struct {
	feud.BaseBarony
}

var BPeriyapatna佩里亚帕提纳 feud.Barony = &佩里亚帕提纳PeriyapatnaBarony{}

func init() {
    f := BPeriyapatna佩里亚帕提纳.(*佩里亚帕提纳PeriyapatnaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "periyapatna",
		TitleName: "佩里亚帕提纳",
		TitleCode: "b_periyapatna",
	}
}
