package lhunze

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 准巴ZhunbaBarony struct {
	feud.BaseBarony
}

var BZhunba准巴 feud.Barony = &准巴ZhunbaBarony{}

func init() {
	f := BZhunba准巴.(*准巴ZhunbaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zhunba",
		TitleName: "准巴",
		TitleCode: "b_zhunba",
	}
}
