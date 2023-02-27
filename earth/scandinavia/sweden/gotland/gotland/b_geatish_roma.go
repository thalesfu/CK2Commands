package gotland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗玛Geatish_romaBarony struct {
	feud.BaseBarony
}

var BGeatish_roma罗玛 feud.Barony = &罗玛Geatish_romaBarony{}

func init() {
    f := BGeatish_roma罗玛.(*罗玛Geatish_romaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "geatish_roma",
		TitleName: "罗玛",
		TitleCode: "b_geatish_roma",
	}
}
