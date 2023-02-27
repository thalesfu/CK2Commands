package hebron

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣查尔顿SaintcharltonBarony struct {
	feud.BaseBarony
}

var BSaintcharlton圣查尔顿 feud.Barony = &圣查尔顿SaintcharltonBarony{}

func init() {
    f := BSaintcharlton圣查尔顿.(*圣查尔顿SaintcharltonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saintcharlton",
		TitleName: "圣查尔顿",
		TitleCode: "b_saintcharlton",
	}
}
