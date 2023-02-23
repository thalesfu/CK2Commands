package ennedi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科拜KobbaiBarony struct {
	feud.BaseBarony
}

var BKobbai科拜 feud.Barony = &科拜KobbaiBarony{}

func init() {
	f := BKobbai科拜.(*科拜KobbaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kobbai",
		TitleName: "科拜",
		TitleCode: "b_kobbai",
	}
}
