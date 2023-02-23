package chur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库尔瓦尔登ChurwaldenBarony struct {
	feud.BaseBarony
}

var BChurwalden库尔瓦尔登 feud.Barony = &库尔瓦尔登ChurwaldenBarony{}

func init() {
	f := BChurwalden库尔瓦尔登.(*库尔瓦尔登ChurwaldenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "churwalden",
		TitleName: "库尔瓦尔登",
		TitleCode: "b_churwalden",
	}
}
