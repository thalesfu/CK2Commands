package bedford

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 邓斯特布尔DunstableBarony struct {
	feud.BaseBarony
}

var BDunstable邓斯特布尔 feud.Barony = &邓斯特布尔DunstableBarony{}

func init() {
    f := BDunstable邓斯特布尔.(*邓斯特布尔DunstableBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dunstable",
		TitleName: "邓斯特布尔",
		TitleCode: "b_dunstable",
	}
}
