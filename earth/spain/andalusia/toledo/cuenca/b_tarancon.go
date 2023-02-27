package cuenca

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔兰孔TaranconBarony struct {
	feud.BaseBarony
}

var BTarancon塔兰孔 feud.Barony = &塔兰孔TaranconBarony{}

func init() {
    f := BTarancon塔兰孔.(*塔兰孔TaranconBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tarancon",
		TitleName: "塔兰孔",
		TitleCode: "b_tarancon",
	}
}
