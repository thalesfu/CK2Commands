package bira

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 喀拉雷比特QalarebeteBarony struct {
	feud.BaseBarony
}

var BQalarebete喀拉雷比特 feud.Barony = &喀拉雷比特QalarebeteBarony{}

func init() {
    f := BQalarebete喀拉雷比特.(*喀拉雷比特QalarebeteBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qalarebete",
		TitleName: "喀拉雷比特",
		TitleCode: "b_qalarebete",
	}
}
