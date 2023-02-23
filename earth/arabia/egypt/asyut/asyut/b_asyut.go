package asyut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾斯尤特AsyutBarony struct {
	feud.BaseBarony
}

var BAsyut艾斯尤特 feud.Barony = &艾斯尤特AsyutBarony{}

func init() {
	f := BAsyut艾斯尤特.(*艾斯尤特AsyutBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "asyut",
		TitleName: "艾斯尤特",
		TitleCode: "b_asyut",
	}
}
