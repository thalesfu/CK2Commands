package udabhanda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 门吉亚罗MankialaBarony struct {
	feud.BaseBarony
}

var BMankiala门吉亚罗 feud.Barony = &门吉亚罗MankialaBarony{}

func init() {
    f := BMankiala门吉亚罗.(*门吉亚罗MankialaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mankiala",
		TitleName: "门吉亚罗",
		TitleCode: "b_mankiala",
	}
}
