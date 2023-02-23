package nikomedeia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克里索波利斯ChrysopolisBarony struct {
	feud.BaseBarony
}

var BChrysopolis克里索波利斯 feud.Barony = &克里索波利斯ChrysopolisBarony{}

func init() {
	f := BChrysopolis克里索波利斯.(*克里索波利斯ChrysopolisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chrysopolis",
		TitleName: "克里索波利斯",
		TitleCode: "b_chrysopolis",
	}
}
