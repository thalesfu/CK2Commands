package tarsos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泽胡里尤ZephyriumBarony struct {
	feud.BaseBarony
}

var BZephyrium泽胡里尤 feud.Barony = &泽胡里尤ZephyriumBarony{}

func init() {
    f := BZephyrium泽胡里尤.(*泽胡里尤ZephyriumBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zephyrium",
		TitleName: "泽胡里尤",
		TitleCode: "b_zephyrium",
	}
}
