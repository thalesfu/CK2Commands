package tikhvine

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴巴耶沃BabayevoBarony struct {
	feud.BaseBarony
}

var BBabayevo巴巴耶沃 feud.Barony = &巴巴耶沃BabayevoBarony{}

func init() {
	f := BBabayevo巴巴耶沃.(*巴巴耶沃BabayevoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "babayevo",
		TitleName: "巴巴耶沃",
		TitleCode: "b_babayevo",
	}
}
