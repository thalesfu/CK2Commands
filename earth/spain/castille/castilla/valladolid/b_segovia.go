package valladolid

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞哥维亚SegoviaBarony struct {
	feud.BaseBarony
}

var BSegovia塞哥维亚 feud.Barony = &塞哥维亚SegoviaBarony{}

func init() {
    f := BSegovia塞哥维亚.(*塞哥维亚SegoviaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "segovia",
		TitleName: "塞哥维亚",
		TitleCode: "b_segovia",
	}
}
