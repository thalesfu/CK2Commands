package tudgha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 兹拜尔ZbairBarony struct {
	feud.BaseBarony
}

var BZbair兹拜尔 feud.Barony = &兹拜尔ZbairBarony{}

func init() {
    f := BZbair兹拜尔.(*兹拜尔ZbairBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zbair",
		TitleName: "兹拜尔",
		TitleCode: "b_zbair",
	}
}
