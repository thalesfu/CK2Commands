package alexandretta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗苏斯RhosusBarony struct {
	feud.BaseBarony
}

var BRhosus罗苏斯 feud.Barony = &罗苏斯RhosusBarony{}

func init() {
    f := BRhosus罗苏斯.(*罗苏斯RhosusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rhosus",
		TitleName: "罗苏斯",
		TitleCode: "b_rhosus",
	}
}
