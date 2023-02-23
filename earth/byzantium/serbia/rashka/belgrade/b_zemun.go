package belgrade

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泽姆林ZemunBarony struct {
	feud.BaseBarony
}

var BZemun泽姆林 feud.Barony = &泽姆林ZemunBarony{}

func init() {
	f := BZemun泽姆林.(*泽姆林ZemunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zemun",
		TitleName: "泽姆林",
		TitleCode: "b_zemun",
	}
}
