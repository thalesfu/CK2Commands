package avranches

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库唐斯CoutancesBarony struct {
	feud.BaseBarony
}

var BCoutances库唐斯 feud.Barony = &库唐斯CoutancesBarony{}

func init() {
    f := BCoutances库唐斯.(*库唐斯CoutancesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "coutances",
		TitleName: "库唐斯",
		TitleCode: "b_coutances",
	}
}
