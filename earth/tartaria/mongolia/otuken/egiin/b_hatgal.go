package egiin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈特嘎勒HatgalBarony struct {
	feud.BaseBarony
}

var BHatgal哈特嘎勒 feud.Barony = &哈特嘎勒HatgalBarony{}

func init() {
	f := BHatgal哈特嘎勒.(*哈特嘎勒HatgalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hatgal",
		TitleName: "哈特嘎勒",
		TitleCode: "b_hatgal",
	}
}
