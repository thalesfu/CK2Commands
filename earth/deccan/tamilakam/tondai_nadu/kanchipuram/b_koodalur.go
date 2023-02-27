package kanchipuram

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 句陀楼罗KoodalurBarony struct {
	feud.BaseBarony
}

var BKoodalur句陀楼罗 feud.Barony = &句陀楼罗KoodalurBarony{}

func init() {
    f := BKoodalur句陀楼罗.(*句陀楼罗KoodalurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "koodalur",
		TitleName: "句陀楼罗",
		TitleCode: "b_koodalur",
	}
}
