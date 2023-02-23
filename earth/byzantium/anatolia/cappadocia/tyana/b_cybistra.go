package tyana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基比斯特拉CybistraBarony struct {
	feud.BaseBarony
}

var BCybistra基比斯特拉 feud.Barony = &基比斯特拉CybistraBarony{}

func init() {
	f := BCybistra基比斯特拉.(*基比斯特拉CybistraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cybistra",
		TitleName: "基比斯特拉",
		TitleCode: "b_cybistra",
	}
}
