package tsilma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 农布尔NonburBarony struct {
	feud.BaseBarony
}

var BNonbur农布尔 feud.Barony = &农布尔NonburBarony{}

func init() {
    f := BNonbur农布尔.(*农布尔NonburBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nonbur",
		TitleName: "农布尔",
		TitleCode: "b_nonbur",
	}
}
