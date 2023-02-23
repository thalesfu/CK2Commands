package quena

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉夫特QiftBarony struct {
	feud.BaseBarony
}

var BQift吉夫特 feud.Barony = &吉夫特QiftBarony{}

func init() {
	f := BQift吉夫特.(*吉夫特QiftBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qift",
		TitleName: "吉夫特",
		TitleCode: "b_qift",
	}
}
