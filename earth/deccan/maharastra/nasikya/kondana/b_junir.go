package kondana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 朱尼尔JunirBarony struct {
	feud.BaseBarony
}

var BJunir朱尼尔 feud.Barony = &朱尼尔JunirBarony{}

func init() {
	f := BJunir朱尼尔.(*朱尼尔JunirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "junir",
		TitleName: "朱尼尔",
		TitleCode: "b_junir",
	}
}
