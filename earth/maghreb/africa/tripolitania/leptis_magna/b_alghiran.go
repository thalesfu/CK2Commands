package leptis_magna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾格赫拉AlghiranBarony struct {
	feud.BaseBarony
}

var BAlghiran艾格赫拉 feud.Barony = &艾格赫拉AlghiranBarony{}

func init() {
    f := BAlghiran艾格赫拉.(*艾格赫拉AlghiranBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alghiran",
		TitleName: "艾格赫拉",
		TitleCode: "b_alghiran",
	}
}
