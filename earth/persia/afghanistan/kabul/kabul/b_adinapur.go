package kabul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 那揭罗喝国AdinapurBarony struct {
	feud.BaseBarony
}

var BAdinapur那揭罗喝国 feud.Barony = &那揭罗喝国AdinapurBarony{}

func init() {
    f := BAdinapur那揭罗喝国.(*那揭罗喝国AdinapurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "adinapur",
		TitleName: "那揭罗喝国",
		TitleCode: "b_adinapur",
	}
}
