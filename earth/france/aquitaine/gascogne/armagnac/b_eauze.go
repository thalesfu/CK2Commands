package armagnac

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 欧兹EauzeBarony struct {
	feud.BaseBarony
}

var BEauze欧兹 feud.Barony = &欧兹EauzeBarony{}

func init() {
    f := BEauze欧兹.(*欧兹EauzeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "eauze",
		TitleName: "欧兹",
		TitleCode: "b_eauze",
	}
}
