package aquileia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 西斯蒂亚那SistianaBarony struct {
	feud.BaseBarony
}

var BSistiana西斯蒂亚那 feud.Barony = &西斯蒂亚那SistianaBarony{}

func init() {
    f := BSistiana西斯蒂亚那.(*西斯蒂亚那SistianaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sistiana",
		TitleName: "西斯蒂亚那",
		TitleCode: "b_sistiana",
	}
}
