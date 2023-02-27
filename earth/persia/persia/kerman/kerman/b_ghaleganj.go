package kerman

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉哈勒甘吉GhaleganjBarony struct {
	feud.BaseBarony
}

var BGhaleganj吉哈勒甘吉 feud.Barony = &吉哈勒甘吉GhaleganjBarony{}

func init() {
    f := BGhaleganj吉哈勒甘吉.(*吉哈勒甘吉GhaleganjBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ghaleganj",
		TitleName: "吉哈勒甘吉",
		TitleCode: "b_ghaleganj",
	}
}
