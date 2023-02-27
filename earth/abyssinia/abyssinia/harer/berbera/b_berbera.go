package berbera

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拨拔力BerberaBarony struct {
	feud.BaseBarony
}

var BBerbera拨拔力 feud.Barony = &拨拔力BerberaBarony{}

func init() {
    f := BBerbera拨拔力.(*拨拔力BerberaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "berbera",
		TitleName: "拨拔力",
		TitleCode: "b_berbera",
	}
}
