package vendome

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗米伊RomillyBarony struct {
	feud.BaseBarony
}

var BRomilly罗米伊 feud.Barony = &罗米伊RomillyBarony{}

func init() {
	f := BRomilly罗米伊.(*罗米伊RomillyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "romilly",
		TitleName: "罗米伊",
		TitleCode: "b_romilly",
	}
}
