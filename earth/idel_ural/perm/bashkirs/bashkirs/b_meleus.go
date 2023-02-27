package bashkirs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅列乌兹MeleusBarony struct {
	feud.BaseBarony
}

var BMeleus梅列乌兹 feud.Barony = &梅列乌兹MeleusBarony{}

func init() {
    f := BMeleus梅列乌兹.(*梅列乌兹MeleusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "meleus",
		TitleName: "梅列乌兹",
		TitleCode: "b_meleus",
	}
}
