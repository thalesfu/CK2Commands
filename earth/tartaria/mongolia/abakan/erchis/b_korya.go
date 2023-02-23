package erchis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科尔雅KoryaBarony struct {
	feud.BaseBarony
}

var BKorya科尔雅 feud.Barony = &科尔雅KoryaBarony{}

func init() {
	f := BKorya科尔雅.(*科尔雅KoryaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "korya",
		TitleName: "科尔雅",
		TitleCode: "b_korya",
	}
}
