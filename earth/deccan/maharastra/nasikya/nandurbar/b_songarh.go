package nandurbar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 僧姞利呬SongarhBarony struct {
	feud.BaseBarony
}

var BSongarh僧姞利呬 feud.Barony = &僧姞利呬SongarhBarony{}

func init() {
    f := BSongarh僧姞利呬.(*僧姞利呬SongarhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "songarh",
		TitleName: "僧姞利呬",
		TitleCode: "b_songarh",
	}
}
