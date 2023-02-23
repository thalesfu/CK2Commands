package mozhaysk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙胡尼亚ShakhunyaBarony struct {
	feud.BaseBarony
}

var BShakhunya沙胡尼亚 feud.Barony = &沙胡尼亚ShakhunyaBarony{}

func init() {
	f := BShakhunya沙胡尼亚.(*沙胡尼亚ShakhunyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shakhunya",
		TitleName: "沙胡尼亚",
		TitleCode: "b_shakhunya",
	}
}
