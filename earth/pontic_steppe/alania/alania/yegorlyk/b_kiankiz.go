package yegorlyk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基安基兹KiankizBarony struct {
	feud.BaseBarony
}

var BKiankiz基安基兹 feud.Barony = &基安基兹KiankizBarony{}

func init() {
    f := BKiankiz基安基兹.(*基安基兹KiankizBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kiankiz",
		TitleName: "基安基兹",
		TitleCode: "b_kiankiz",
	}
}
