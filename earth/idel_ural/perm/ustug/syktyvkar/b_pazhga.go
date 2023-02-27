package syktyvkar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕日加PazhgaBarony struct {
	feud.BaseBarony
}

var BPazhga帕日加 feud.Barony = &帕日加PazhgaBarony{}

func init() {
    f := BPazhga帕日加.(*帕日加PazhgaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pazhga",
		TitleName: "帕日加",
		TitleCode: "b_pazhga",
	}
}
