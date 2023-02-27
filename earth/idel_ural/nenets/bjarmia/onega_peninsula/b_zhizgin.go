package onega_peninsula

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 日日金ZhizginBarony struct {
	feud.BaseBarony
}

var BZhizgin日日金 feud.Barony = &日日金ZhizginBarony{}

func init() {
    f := BZhizgin日日金.(*日日金ZhizginBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zhizgin",
		TitleName: "日日金",
		TitleCode: "b_zhizgin",
	}
}
