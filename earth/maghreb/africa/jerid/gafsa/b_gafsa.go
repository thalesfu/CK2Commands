package gafsa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加夫萨GafsaBarony struct {
	feud.BaseBarony
}

var BGafsa加夫萨 feud.Barony = &加夫萨GafsaBarony{}

func init() {
	f := BGafsa加夫萨.(*加夫萨GafsaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gafsa",
		TitleName: "加夫萨",
		TitleCode: "b_gafsa",
	}
}
