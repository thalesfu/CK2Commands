package oleshye

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 第聂伯DniproBarony struct {
	feud.BaseBarony
}

var BDnipro第聂伯 feud.Barony = &第聂伯DniproBarony{}

func init() {
    f := BDnipro第聂伯.(*第聂伯DniproBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dnipro",
		TitleName: "第聂伯",
		TitleCode: "b_dnipro",
	}
}
