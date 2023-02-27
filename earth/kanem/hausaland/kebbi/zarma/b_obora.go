package zarma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥博拉OboraBarony struct {
	feud.BaseBarony
}

var BObora奥博拉 feud.Barony = &奥博拉OboraBarony{}

func init() {
    f := BObora奥博拉.(*奥博拉OboraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "obora",
		TitleName: "奥博拉",
		TitleCode: "b_obora",
	}
}
