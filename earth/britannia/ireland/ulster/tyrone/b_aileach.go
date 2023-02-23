package tyrone

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾丽赫AileachBarony struct {
	feud.BaseBarony
}

var BAileach艾丽赫 feud.Barony = &艾丽赫AileachBarony{}

func init() {
	f := BAileach艾丽赫.(*艾丽赫AileachBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aileach",
		TitleName: "艾丽赫",
		TitleCode: "b_aileach",
	}
}
