package shahrazur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙赫里祖尔ShahrazurBarony struct {
	feud.BaseBarony
}

var BShahrazur沙赫里祖尔 feud.Barony = &沙赫里祖尔ShahrazurBarony{}

func init() {
	f := BShahrazur沙赫里祖尔.(*沙赫里祖尔ShahrazurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shahrazur",
		TitleName: "沙赫里祖尔",
		TitleCode: "b_shahrazur",
	}
}
