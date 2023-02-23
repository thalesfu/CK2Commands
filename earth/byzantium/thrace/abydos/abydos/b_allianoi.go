package abydos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿里亚诺伊AllianoiBarony struct {
	feud.BaseBarony
}

var BAllianoi阿里亚诺伊 feud.Barony = &阿里亚诺伊AllianoiBarony{}

func init() {
	f := BAllianoi阿里亚诺伊.(*阿里亚诺伊AllianoiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "allianoi",
		TitleName: "阿里亚诺伊",
		TitleCode: "b_allianoi",
	}
}
