package beersheb

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥法基姆OfakimBarony struct {
	feud.BaseBarony
}

var BOfakim奥法基姆 feud.Barony = &奥法基姆OfakimBarony{}

func init() {
    f := BOfakim奥法基姆.(*奥法基姆OfakimBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ofakim",
		TitleName: "奥法基姆",
		TitleCode: "b_ofakim",
	}
}
