package luntai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 且弥QiemiBarony struct {
	feud.BaseBarony
}

var BQiemi且弥 feud.Barony = &且弥QiemiBarony{}

func init() {
	f := BQiemi且弥.(*且弥QiemiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qiemi",
		TitleName: "且弥",
		TitleCode: "b_qiemi",
	}
}
