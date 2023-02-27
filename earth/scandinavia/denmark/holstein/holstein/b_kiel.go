package holstein

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基尔KielBarony struct {
	feud.BaseBarony
}

var BKiel基尔 feud.Barony = &基尔KielBarony{}

func init() {
    f := BKiel基尔.(*基尔KielBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kiel",
		TitleName: "基尔",
		TitleCode: "b_kiel",
	}
}
