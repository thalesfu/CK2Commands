package kent

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 坎特伯雷CanterburyBarony struct {
	feud.BaseBarony
}

var BCanterbury坎特伯雷 feud.Barony = &坎特伯雷CanterburyBarony{}

func init() {
    f := BCanterbury坎特伯雷.(*坎特伯雷CanterburyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "canterbury",
		TitleName: "坎特伯雷",
		TitleCode: "b_canterbury",
	}
}
