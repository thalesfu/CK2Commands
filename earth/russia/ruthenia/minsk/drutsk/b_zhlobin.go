package drutsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 日洛宾ZhlobinBarony struct {
	feud.BaseBarony
}

var BZhlobin日洛宾 feud.Barony = &日洛宾ZhlobinBarony{}

func init() {
    f := BZhlobin日洛宾.(*日洛宾ZhlobinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zhlobin",
		TitleName: "日洛宾",
		TitleCode: "b_zhlobin",
	}
}
