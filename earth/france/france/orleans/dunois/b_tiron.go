package dunois

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒂龙TironBarony struct {
	feud.BaseBarony
}

var BTiron蒂龙 feud.Barony = &蒂龙TironBarony{}

func init() {
    f := BTiron蒂龙.(*蒂龙TironBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tiron",
		TitleName: "蒂龙",
		TitleCode: "b_tiron",
	}
}
