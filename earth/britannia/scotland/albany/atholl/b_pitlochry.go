package atholl

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 皮特洛赫里PitlochryBarony struct {
	feud.BaseBarony
}

var BPitlochry皮特洛赫里 feud.Barony = &皮特洛赫里PitlochryBarony{}

func init() {
    f := BPitlochry皮特洛赫里.(*皮特洛赫里PitlochryBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pitlochry",
		TitleName: "皮特洛赫里",
		TitleCode: "b_pitlochry",
	}
}
