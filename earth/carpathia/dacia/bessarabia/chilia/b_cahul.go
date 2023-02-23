package chilia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡胡尔CahulBarony struct {
	feud.BaseBarony
}

var BCahul卡胡尔 feud.Barony = &卡胡尔CahulBarony{}

func init() {
	f := BCahul卡胡尔.(*卡胡尔CahulBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cahul",
		TitleName: "卡胡尔",
		TitleCode: "b_cahul",
	}
}
