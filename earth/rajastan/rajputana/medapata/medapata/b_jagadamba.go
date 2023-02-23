package medapata

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 社伽淡婆JagadambaBarony struct {
	feud.BaseBarony
}

var BJagadamba社伽淡婆 feud.Barony = &社伽淡婆JagadambaBarony{}

func init() {
	f := BJagadamba社伽淡婆.(*社伽淡婆JagadambaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jagadamba",
		TitleName: "社伽淡婆",
		TitleCode: "b_jagadamba",
	}
}
