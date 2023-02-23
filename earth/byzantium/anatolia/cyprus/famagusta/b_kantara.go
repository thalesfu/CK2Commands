package famagusta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 坎塔拉KantaraBarony struct {
	feud.BaseBarony
}

var BKantara坎塔拉 feud.Barony = &坎塔拉KantaraBarony{}

func init() {
	f := BKantara坎塔拉.(*坎塔拉KantaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kantara",
		TitleName: "坎塔拉",
		TitleCode: "b_kantara",
	}
}
