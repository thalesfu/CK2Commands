package taghaza

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿德拉尔AdrarBarony struct {
	feud.BaseBarony
}

var BAdrar阿德拉尔 feud.Barony = &阿德拉尔AdrarBarony{}

func init() {
    f := BAdrar阿德拉尔.(*阿德拉尔AdrarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "adrar",
		TitleName: "阿德拉尔",
		TitleCode: "b_adrar",
	}
}
