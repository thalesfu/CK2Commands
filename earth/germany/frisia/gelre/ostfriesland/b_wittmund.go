package ostfriesland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维特蒙德WittmundBarony struct {
	feud.BaseBarony
}

var BWittmund维特蒙德 feud.Barony = &维特蒙德WittmundBarony{}

func init() {
    f := BWittmund维特蒙德.(*维特蒙德WittmundBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wittmund",
		TitleName: "维特蒙德",
		TitleCode: "b_wittmund",
	}
}
