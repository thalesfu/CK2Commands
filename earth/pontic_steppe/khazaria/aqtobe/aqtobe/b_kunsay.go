package aqtobe

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 昆赛KunsayBarony struct {
	feud.BaseBarony
}

var BKunsay昆赛 feud.Barony = &昆赛KunsayBarony{}

func init() {
    f := BKunsay昆赛.(*昆赛KunsayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kunsay",
		TitleName: "昆赛",
		TitleCode: "b_kunsay",
	}
}
