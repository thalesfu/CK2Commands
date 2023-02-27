package paraetonium

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕莱托尼翁ParaetoniumBarony struct {
	feud.BaseBarony
}

var BParaetonium帕莱托尼翁 feud.Barony = &帕莱托尼翁ParaetoniumBarony{}

func init() {
    f := BParaetonium帕莱托尼翁.(*帕莱托尼翁ParaetoniumBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "paraetonium",
		TitleName: "帕莱托尼翁",
		TitleCode: "b_paraetonium",
	}
}
