package kazakh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杰孜吉尔ZhezgirBarony struct {
	feud.BaseBarony
}

var BZhezgir杰孜吉尔 feud.Barony = &杰孜吉尔ZhezgirBarony{}

func init() {
	f := BZhezgir杰孜吉尔.(*杰孜吉尔ZhezgirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zhezgir",
		TitleName: "杰孜吉尔",
		TitleCode: "b_zhezgir",
	}
}
