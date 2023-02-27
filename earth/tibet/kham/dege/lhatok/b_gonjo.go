package lhatok

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 馆觉GonjoBarony struct {
	feud.BaseBarony
}

var BGonjo馆觉 feud.Barony = &馆觉GonjoBarony{}

func init() {
    f := BGonjo馆觉.(*馆觉GonjoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gonjo",
		TitleName: "馆觉",
		TitleCode: "b_gonjo",
	}
}
