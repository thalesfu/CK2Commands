package romny

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维索卡亚VysokayaBarony struct {
	feud.BaseBarony
}

var BVysokaya维索卡亚 feud.Barony = &维索卡亚VysokayaBarony{}

func init() {
    f := BVysokaya维索卡亚.(*维索卡亚VysokayaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vysokaya",
		TitleName: "维索卡亚",
		TitleCode: "b_vysokaya",
	}
}
