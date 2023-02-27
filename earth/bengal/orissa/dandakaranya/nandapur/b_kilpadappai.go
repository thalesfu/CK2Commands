package nandapur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉波陀吠KilpadappaiBarony struct {
	feud.BaseBarony
}

var BKilpadappai吉波陀吠 feud.Barony = &吉波陀吠KilpadappaiBarony{}

func init() {
    f := BKilpadappai吉波陀吠.(*吉波陀吠KilpadappaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kilpadappai",
		TitleName: "吉波陀吠",
		TitleCode: "b_kilpadappai",
	}
}
