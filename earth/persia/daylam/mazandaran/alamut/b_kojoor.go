package alamut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阔佐尔KojoorBarony struct {
	feud.BaseBarony
}

var BKojoor阔佐尔 feud.Barony = &阔佐尔KojoorBarony{}

func init() {
	f := BKojoor阔佐尔.(*阔佐尔KojoorBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kojoor",
		TitleName: "阔佐尔",
		TitleCode: "b_kojoor",
	}
}
