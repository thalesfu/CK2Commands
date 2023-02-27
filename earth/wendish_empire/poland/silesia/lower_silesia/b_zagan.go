package lower_silesia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎甘ZaganBarony struct {
	feud.BaseBarony
}

var BZagan扎甘 feud.Barony = &扎甘ZaganBarony{}

func init() {
    f := BZagan扎甘.(*扎甘ZaganBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zagan",
		TitleName: "扎甘",
		TitleCode: "b_zagan",
	}
}
