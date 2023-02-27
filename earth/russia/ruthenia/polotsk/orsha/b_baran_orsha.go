package orsha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴兰Baran_orshaBarony struct {
	feud.BaseBarony
}

var BBaran_orsha巴兰 feud.Barony = &巴兰Baran_orshaBarony{}

func init() {
    f := BBaran_orsha巴兰.(*巴兰Baran_orshaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "baran_orsha",
		TitleName: "巴兰",
		TitleCode: "b_baran_orsha",
	}
}
