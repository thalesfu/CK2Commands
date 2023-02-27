package wolgast

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 凯姆尼茨KemnitzBarony struct {
	feud.BaseBarony
}

var BKemnitz凯姆尼茨 feud.Barony = &凯姆尼茨KemnitzBarony{}

func init() {
    f := BKemnitz凯姆尼茨.(*凯姆尼茨KemnitzBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kemnitz",
		TitleName: "凯姆尼茨",
		TitleCode: "b_kemnitz",
	}
}
