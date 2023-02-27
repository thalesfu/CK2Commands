package al_djazair

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔及尔AlgiersBarony struct {
	feud.BaseBarony
}

var BAlgiers阿尔及尔 feud.Barony = &阿尔及尔AlgiersBarony{}

func init() {
    f := BAlgiers阿尔及尔.(*阿尔及尔AlgiersBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "algiers",
		TitleName: "阿尔及尔",
		TitleCode: "b_algiers",
	}
}
