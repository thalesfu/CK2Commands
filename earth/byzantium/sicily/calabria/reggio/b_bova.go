package reggio

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博瓦BovaBarony struct {
	feud.BaseBarony
}

var BBova博瓦 feud.Barony = &博瓦BovaBarony{}

func init() {
    f := BBova博瓦.(*博瓦BovaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bova",
		TitleName: "博瓦",
		TitleCode: "b_bova",
	}
}
