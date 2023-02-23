package kutch

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿迦尔AnjarBarony struct {
	feud.BaseBarony
}

var BAnjar阿迦尔 feud.Barony = &阿迦尔AnjarBarony{}

func init() {
	f := BAnjar阿迦尔.(*阿迦尔AnjarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "anjar",
		TitleName: "阿迦尔",
		TitleCode: "b_anjar",
	}
}
