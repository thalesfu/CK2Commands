package syrmia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 兹韦钱Zvecan_syrmiaBarony struct {
	feud.BaseBarony
}

var BZvecan_syrmia兹韦钱 feud.Barony = &兹韦钱Zvecan_syrmiaBarony{}

func init() {
    f := BZvecan_syrmia兹韦钱.(*兹韦钱Zvecan_syrmiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zvecan_syrmia",
		TitleName: "兹韦钱",
		TitleCode: "b_zvecan_syrmia",
	}
}
