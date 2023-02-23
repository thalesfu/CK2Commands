package gondar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德姆比亚DembiyaBarony struct {
	feud.BaseBarony
}

var BDembiya德姆比亚 feud.Barony = &德姆比亚DembiyaBarony{}

func init() {
	f := BDembiya德姆比亚.(*德姆比亚DembiyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dembiya",
		TitleName: "德姆比亚",
		TitleCode: "b_dembiya",
	}
}
