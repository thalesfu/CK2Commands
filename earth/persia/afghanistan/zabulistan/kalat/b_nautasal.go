package kalat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瑙塔萨勒NautasalBarony struct {
	feud.BaseBarony
}

var BNautasal瑙塔萨勒 feud.Barony = &瑙塔萨勒NautasalBarony{}

func init() {
	f := BNautasal瑙塔萨勒.(*瑙塔萨勒NautasalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nautasal",
		TitleName: "瑙塔萨勒",
		TitleCode: "b_nautasal",
	}
}
