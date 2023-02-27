package gabes

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉比利KebiliBarony struct {
	feud.BaseBarony
}

var BKebili吉比利 feud.Barony = &吉比利KebiliBarony{}

func init() {
    f := BKebili吉比利.(*吉比利KebiliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kebili",
		TitleName: "吉比利",
		TitleCode: "b_kebili",
	}
}
