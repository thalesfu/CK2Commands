package votyaki

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿奇特AchitBarony struct {
	feud.BaseBarony
}

var BAchit阿奇特 feud.Barony = &阿奇特AchitBarony{}

func init() {
    f := BAchit阿奇特.(*阿奇特AchitBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "achit",
		TitleName: "阿奇特",
		TitleCode: "b_achit",
	}
}
