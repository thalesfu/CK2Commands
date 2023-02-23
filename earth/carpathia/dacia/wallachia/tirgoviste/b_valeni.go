package tirgoviste

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃莱尼ValeniBarony struct {
	feud.BaseBarony
}

var BValeni沃莱尼 feud.Barony = &沃莱尼ValeniBarony{}

func init() {
	f := BValeni沃莱尼.(*沃莱尼ValeniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "valeni",
		TitleName: "沃莱尼",
		TitleCode: "b_valeni",
	}
}
