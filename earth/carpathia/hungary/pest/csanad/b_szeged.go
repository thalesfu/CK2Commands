package csanad

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞盖德SzegedBarony struct {
	feud.BaseBarony
}

var BSzeged塞盖德 feud.Barony = &塞盖德SzegedBarony{}

func init() {
	f := BSzeged塞盖德.(*塞盖德SzegedBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "szeged",
		TitleName: "塞盖德",
		TitleCode: "b_szeged",
	}
}
