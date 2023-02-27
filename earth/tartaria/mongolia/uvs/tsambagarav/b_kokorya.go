package tsambagarav

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科科里亚KokoryaBarony struct {
	feud.BaseBarony
}

var BKokorya科科里亚 feud.Barony = &科科里亚KokoryaBarony{}

func init() {
    f := BKokorya科科里亚.(*科科里亚KokoryaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kokorya",
		TitleName: "科科里亚",
		TitleCode: "b_kokorya",
	}
}
