package traianopolis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库普塞拉CypselaBarony struct {
	feud.BaseBarony
}

var BCypsela库普塞拉 feud.Barony = &库普塞拉CypselaBarony{}

func init() {
    f := BCypsela库普塞拉.(*库普塞拉CypselaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cypsela",
		TitleName: "库普塞拉",
		TitleCode: "b_cypsela",
	}
}
