package yatenga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 宗多马ZondomaBarony struct {
	feud.BaseBarony
}

var BZondoma宗多马 feud.Barony = &宗多马ZondomaBarony{}

func init() {
    f := BZondoma宗多马.(*宗多马ZondomaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zondoma",
		TitleName: "宗多马",
		TitleCode: "b_zondoma",
	}
}
