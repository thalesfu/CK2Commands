package taradavadi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伽蓝婆俱提KarambakkudiBarony struct {
	feud.BaseBarony
}

var BKarambakkudi伽蓝婆俱提 feud.Barony = &伽蓝婆俱提KarambakkudiBarony{}

func init() {
    f := BKarambakkudi伽蓝婆俱提.(*伽蓝婆俱提KarambakkudiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karambakkudi",
		TitleName: "伽蓝婆俱提",
		TitleCode: "b_karambakkudi",
	}
}
