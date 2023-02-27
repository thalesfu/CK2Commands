package danzig

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 但泽DanzigBarony struct {
	feud.BaseBarony
}

var BDanzig但泽 feud.Barony = &但泽DanzigBarony{}

func init() {
    f := BDanzig但泽.(*但泽DanzigBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "danzig",
		TitleName: "但泽",
		TitleCode: "b_danzig",
	}
}
