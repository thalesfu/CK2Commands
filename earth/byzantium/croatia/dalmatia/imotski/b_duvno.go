package imotski

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杜夫诺DuvnoBarony struct {
	feud.BaseBarony
}

var BDuvno杜夫诺 feud.Barony = &杜夫诺DuvnoBarony{}

func init() {
	f := BDuvno杜夫诺.(*杜夫诺DuvnoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "duvno",
		TitleName: "杜夫诺",
		TitleCode: "b_duvno",
	}
}
