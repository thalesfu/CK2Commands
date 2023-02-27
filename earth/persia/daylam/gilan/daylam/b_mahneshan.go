package daylam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马赫内尚MahneshanBarony struct {
	feud.BaseBarony
}

var BMahneshan马赫内尚 feud.Barony = &马赫内尚MahneshanBarony{}

func init() {
    f := BMahneshan马赫内尚.(*马赫内尚MahneshanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mahneshan",
		TitleName: "马赫内尚",
		TitleCode: "b_mahneshan",
	}
}
