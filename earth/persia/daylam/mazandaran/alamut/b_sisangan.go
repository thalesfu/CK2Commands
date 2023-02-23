package alamut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯桑干SisanganBarony struct {
	feud.BaseBarony
}

var BSisangan斯桑干 feud.Barony = &斯桑干SisanganBarony{}

func init() {
	f := BSisangan斯桑干.(*斯桑干SisanganBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sisangan",
		TitleName: "斯桑干",
		TitleCode: "b_sisangan",
	}
}
