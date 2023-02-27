package dauphine_viennois

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔邦AlbonBarony struct {
	feud.BaseBarony
}

var BAlbon阿尔邦 feud.Barony = &阿尔邦AlbonBarony{}

func init() {
    f := BAlbon阿尔邦.(*阿尔邦AlbonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "albon",
		TitleName: "阿尔邦",
		TitleCode: "b_albon",
	}
}
