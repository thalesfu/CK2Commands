package assab

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿古脱AgurtoBarony struct {
	feud.BaseBarony
}

var BAgurto阿古脱 feud.Barony = &阿古脱AgurtoBarony{}

func init() {
	f := BAgurto阿古脱.(*阿古脱AgurtoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "agurto",
		TitleName: "阿古脱",
		TitleCode: "b_agurto",
	}
}
