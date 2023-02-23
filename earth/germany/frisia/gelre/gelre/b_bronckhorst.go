package gelre

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布隆克霍斯特BronckhorstBarony struct {
	feud.BaseBarony
}

var BBronckhorst布隆克霍斯特 feud.Barony = &布隆克霍斯特BronckhorstBarony{}

func init() {
	f := BBronckhorst布隆克霍斯特.(*布隆克霍斯特BronckhorstBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bronckhorst",
		TitleName: "布隆克霍斯特",
		TitleCode: "b_bronckhorst",
	}
}
