package peremyshl

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉西夫RasivBarony struct {
	feud.BaseBarony
}

var BRasiv拉西夫 feud.Barony = &拉西夫RasivBarony{}

func init() {
    f := BRasiv拉西夫.(*拉西夫RasivBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rasiv",
		TitleName: "拉西夫",
		TitleCode: "b_rasiv",
	}
}
