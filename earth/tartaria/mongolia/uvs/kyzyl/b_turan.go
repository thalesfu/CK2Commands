package kyzyl

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图兰TuranBarony struct {
	feud.BaseBarony
}

var BTuran图兰 feud.Barony = &图兰TuranBarony{}

func init() {
    f := BTuran图兰.(*图兰TuranBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "turan",
		TitleName: "图兰",
		TitleCode: "b_turan",
	}
}
