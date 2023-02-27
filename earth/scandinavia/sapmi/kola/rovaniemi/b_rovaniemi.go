package rovaniemi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗瓦涅米RovaniemiBarony struct {
	feud.BaseBarony
}

var BRovaniemi罗瓦涅米 feud.Barony = &罗瓦涅米RovaniemiBarony{}

func init() {
    f := BRovaniemi罗瓦涅米.(*罗瓦涅米RovaniemiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rovaniemi",
		TitleName: "罗瓦涅米",
		TitleCode: "b_rovaniemi",
	}
}
