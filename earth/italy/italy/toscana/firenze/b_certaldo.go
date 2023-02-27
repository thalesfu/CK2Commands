package firenze

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 切塔尔多CertaldoBarony struct {
	feud.BaseBarony
}

var BCertaldo切塔尔多 feud.Barony = &切塔尔多CertaldoBarony{}

func init() {
    f := BCertaldo切塔尔多.(*切塔尔多CertaldoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "certaldo",
		TitleName: "切塔尔多",
		TitleCode: "b_certaldo",
	}
}
