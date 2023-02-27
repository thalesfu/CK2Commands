package kagha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿特马加AtemagaBarony struct {
	feud.BaseBarony
}

var BAtemaga阿特马加 feud.Barony = &阿特马加AtemagaBarony{}

func init() {
    f := BAtemaga阿特马加.(*阿特马加AtemagaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "atemaga",
		TitleName: "阿特马加",
		TitleCode: "b_atemaga",
	}
}
