package rothas

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伽蓝婆甘KarumbakhamBarony struct {
	feud.BaseBarony
}

var BKarumbakham伽蓝婆甘 feud.Barony = &伽蓝婆甘KarumbakhamBarony{}

func init() {
	f := BKarumbakham伽蓝婆甘.(*伽蓝婆甘KarumbakhamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karumbakham",
		TitleName: "伽蓝婆甘",
		TitleCode: "b_karumbakham",
	}
}
