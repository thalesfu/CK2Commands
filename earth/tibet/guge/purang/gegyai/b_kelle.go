package gegyai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉勒KelleBarony struct {
	feud.BaseBarony
}

var BKelle吉勒 feud.Barony = &吉勒KelleBarony{}

func init() {
    f := BKelle吉勒.(*吉勒KelleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kelle",
		TitleName: "吉勒",
		TitleCode: "b_kelle",
	}
}
