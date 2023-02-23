package bejaija

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉杰尔JijelBarony struct {
	feud.BaseBarony
}

var BJijel吉杰尔 feud.Barony = &吉杰尔JijelBarony{}

func init() {
	f := BJijel吉杰尔.(*吉杰尔JijelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jijel",
		TitleName: "吉杰尔",
		TitleCode: "b_jijel",
	}
}
