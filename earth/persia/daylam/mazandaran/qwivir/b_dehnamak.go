package qwivir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德纳马克DehnamakBarony struct {
	feud.BaseBarony
}

var BDehnamak德纳马克 feud.Barony = &德纳马克DehnamakBarony{}

func init() {
	f := BDehnamak德纳马克.(*德纳马克DehnamakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dehnamak",
		TitleName: "德纳马克",
		TitleCode: "b_dehnamak",
	}
}
