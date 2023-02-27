package dimapur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩夷菩那迦MaibongBarony struct {
	feud.BaseBarony
}

var BMaibong摩夷菩那迦 feud.Barony = &摩夷菩那迦MaibongBarony{}

func init() {
    f := BMaibong摩夷菩那迦.(*摩夷菩那迦MaibongBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "maibong",
		TitleName: "摩夷菩那迦",
		TitleCode: "b_maibong",
	}
}
