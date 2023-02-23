package nevers

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 栋济DonzyBarony struct {
	feud.BaseBarony
}

var BDonzy栋济 feud.Barony = &栋济DonzyBarony{}

func init() {
	f := BDonzy栋济.(*栋济DonzyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "donzy",
		TitleName: "栋济",
		TitleCode: "b_donzy",
	}
}
