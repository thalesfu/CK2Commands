package argyll

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 邓斯塔夫内奇DunstaffnageBarony struct {
	feud.BaseBarony
}

var BDunstaffnage邓斯塔夫内奇 feud.Barony = &邓斯塔夫内奇DunstaffnageBarony{}

func init() {
    f := BDunstaffnage邓斯塔夫内奇.(*邓斯塔夫内奇DunstaffnageBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dunstaffnage",
		TitleName: "邓斯塔夫内奇",
		TitleCode: "b_dunstaffnage",
	}
}
