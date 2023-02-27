package sevilla

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡莫纳CarmonaBarony struct {
	feud.BaseBarony
}

var BCarmona卡莫纳 feud.Barony = &卡莫纳CarmonaBarony{}

func init() {
    f := BCarmona卡莫纳.(*卡莫纳CarmonaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "carmona",
		TitleName: "卡莫纳",
		TitleCode: "b_carmona",
	}
}
