package priluk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗姆内RomnyBarony struct {
	feud.BaseBarony
}

var BRomny罗姆内 feud.Barony = &罗姆内RomnyBarony{}

func init() {
	f := BRomny罗姆内.(*罗姆内RomnyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "romny",
		TitleName: "罗姆内",
		TitleCode: "b_romny",
	}
}
