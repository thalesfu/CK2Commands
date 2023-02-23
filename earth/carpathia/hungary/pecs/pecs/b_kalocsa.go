package pecs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 考洛乔KalocsaBarony struct {
	feud.BaseBarony
}

var BKalocsa考洛乔 feud.Barony = &考洛乔KalocsaBarony{}

func init() {
	f := BKalocsa考洛乔.(*考洛乔KalocsaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kalocsa",
		TitleName: "考洛乔",
		TitleCode: "b_kalocsa",
	}
}
