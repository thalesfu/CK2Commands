package lhunzhub

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 林周LhunzhubBarony struct {
	feud.BaseBarony
}

var BLhunzhub林周 feud.Barony = &林周LhunzhubBarony{}

func init() {
	f := BLhunzhub林周.(*林周LhunzhubBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lhunzhub",
		TitleName: "林周",
		TitleCode: "b_lhunzhub",
	}
}
