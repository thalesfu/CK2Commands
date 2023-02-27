package chikoi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 舍尔戈利金Shergol_dzhinBarony struct {
	feud.BaseBarony
}

var BShergol_dzhin舍尔戈利金 feud.Barony = &舍尔戈利金Shergol_dzhinBarony{}

func init() {
    f := BShergol_dzhin舍尔戈利金.(*舍尔戈利金Shergol_dzhinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shergol_dzhin",
		TitleName: "舍尔戈利金",
		TitleCode: "b_shergol_dzhin",
	}
}
