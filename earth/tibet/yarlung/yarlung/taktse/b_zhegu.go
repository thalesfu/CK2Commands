package taktse

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哲古ZheguBarony struct {
	feud.BaseBarony
}

var BZhegu哲古 feud.Barony = &哲古ZheguBarony{}

func init() {
    f := BZhegu哲古.(*哲古ZheguBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zhegu",
		TitleName: "哲古",
		TitleCode: "b_zhegu",
	}
}
