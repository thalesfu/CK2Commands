package yuryev

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 别列若克BerezhokBarony struct {
	feud.BaseBarony
}

var BBerezhok别列若克 feud.Barony = &别列若克BerezhokBarony{}

func init() {
	f := BBerezhok别列若克.(*别列若克BerezhokBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "berezhok",
		TitleName: "别列若克",
		TitleCode: "b_berezhok",
	}
}
