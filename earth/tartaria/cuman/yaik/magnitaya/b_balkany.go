package magnitaya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴尔卡内BalkanyBarony struct {
	feud.BaseBarony
}

var BBalkany巴尔卡内 feud.Barony = &巴尔卡内BalkanyBarony{}

func init() {
	f := BBalkany巴尔卡内.(*巴尔卡内BalkanyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "balkany",
		TitleName: "巴尔卡内",
		TitleCode: "b_balkany",
	}
}
