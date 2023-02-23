package madhupur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗波多罗LahpatraBarony struct {
	feud.BaseBarony
}

var BLahpatra罗波多罗 feud.Barony = &罗波多罗LahpatraBarony{}

func init() {
	f := BLahpatra罗波多罗.(*罗波多罗LahpatraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lahpatra",
		TitleName: "罗波多罗",
		TitleCode: "b_lahpatra",
	}
}
