package venaissin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦纳斯克VenasqueBarony struct {
	feud.BaseBarony
}

var BVenasque韦纳斯克 feud.Barony = &韦纳斯克VenasqueBarony{}

func init() {
	f := BVenasque韦纳斯克.(*韦纳斯克VenasqueBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "venasque",
		TitleName: "韦纳斯克",
		TitleCode: "b_venasque",
	}
}
