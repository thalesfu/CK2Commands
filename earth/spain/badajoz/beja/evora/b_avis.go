package evora

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿维什AvisBarony struct {
	feud.BaseBarony
}

var BAvis阿维什 feud.Barony = &阿维什AvisBarony{}

func init() {
	f := BAvis阿维什.(*阿维什AvisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "avis",
		TitleName: "阿维什",
		TitleCode: "b_avis",
	}
}
