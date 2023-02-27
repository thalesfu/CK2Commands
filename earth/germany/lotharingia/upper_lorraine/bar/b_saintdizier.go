package bar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣迪济耶SaintdizierBarony struct {
	feud.BaseBarony
}

var BSaintdizier圣迪济耶 feud.Barony = &圣迪济耶SaintdizierBarony{}

func init() {
    f := BSaintdizier圣迪济耶.(*圣迪济耶SaintdizierBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saintdizier",
		TitleName: "圣迪济耶",
		TitleCode: "b_saintdizier",
	}
}
