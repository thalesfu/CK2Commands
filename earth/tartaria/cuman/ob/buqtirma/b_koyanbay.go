package buqtirma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科扬拜KoyanbayBarony struct {
	feud.BaseBarony
}

var BKoyanbay科扬拜 feud.Barony = &科扬拜KoyanbayBarony{}

func init() {
    f := BKoyanbay科扬拜.(*科扬拜KoyanbayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "koyanbay",
		TitleName: "科扬拜",
		TitleCode: "b_koyanbay",
	}
}
