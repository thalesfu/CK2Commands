package dhofar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞迈里特ThamaritBarony struct {
	feud.BaseBarony
}

var BThamarit塞迈里特 feud.Barony = &塞迈里特ThamaritBarony{}

func init() {
	f := BThamarit塞迈里特.(*塞迈里特ThamaritBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "thamarit",
		TitleName: "塞迈里特",
		TitleCode: "b_thamarit",
	}
}
