package labourd

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索沃泰尔SauveterreBarony struct {
	feud.BaseBarony
}

var BSauveterre索沃泰尔 feud.Barony = &索沃泰尔SauveterreBarony{}

func init() {
    f := BSauveterre索沃泰尔.(*索沃泰尔SauveterreBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sauveterre",
		TitleName: "索沃泰尔",
		TitleCode: "b_sauveterre",
	}
}
