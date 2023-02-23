package makuria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 栋古拉DongolaBarony struct {
	feud.BaseBarony
}

var BDongola栋古拉 feud.Barony = &栋古拉DongolaBarony{}

func init() {
	f := BDongola栋古拉.(*栋古拉DongolaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dongola",
		TitleName: "栋古拉",
		TitleCode: "b_dongola",
	}
}
