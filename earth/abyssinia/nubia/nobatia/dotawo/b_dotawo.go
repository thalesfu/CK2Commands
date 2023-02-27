package dotawo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多塔沃DotawoBarony struct {
	feud.BaseBarony
}

var BDotawo多塔沃 feud.Barony = &多塔沃DotawoBarony{}

func init() {
    f := BDotawo多塔沃.(*多塔沃DotawoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dotawo",
		TitleName: "多塔沃",
		TitleCode: "b_dotawo",
	}
}
