package barcelona

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴塞罗那BarcelonaBarony struct {
	feud.BaseBarony
}

var BBarcelona巴塞罗那 feud.Barony = &巴塞罗那BarcelonaBarony{}

func init() {
    f := BBarcelona巴塞罗那.(*巴塞罗那BarcelonaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "barcelona",
		TitleName: "巴塞罗那",
		TitleCode: "b_barcelona",
	}
}
