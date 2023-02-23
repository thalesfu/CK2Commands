package saintonge

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 桑特SaintesBarony struct {
	feud.BaseBarony
}

var BSaintes桑特 feud.Barony = &桑特SaintesBarony{}

func init() {
	f := BSaintes桑特.(*桑特SaintesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saintes",
		TitleName: "桑特",
		TitleCode: "b_saintes",
	}
}
