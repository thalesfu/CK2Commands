package lemdiyya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布利达BlidaBarony struct {
	feud.BaseBarony
}

var BBlida布利达 feud.Barony = &布利达BlidaBarony{}

func init() {
    f := BBlida布利达.(*布利达BlidaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "blida",
		TitleName: "布利达",
		TitleCode: "b_blida",
	}
}
