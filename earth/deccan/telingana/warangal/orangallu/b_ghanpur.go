package orangallu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 汉普尔GhanpurBarony struct {
	feud.BaseBarony
}

var BGhanpur汉普尔 feud.Barony = &汉普尔GhanpurBarony{}

func init() {
	f := BGhanpur汉普尔.(*汉普尔GhanpurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ghanpur",
		TitleName: "汉普尔",
		TitleCode: "b_ghanpur",
	}
}
