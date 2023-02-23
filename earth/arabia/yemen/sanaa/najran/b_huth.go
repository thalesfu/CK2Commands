package najran

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 侯斯HuthBarony struct {
	feud.BaseBarony
}

var BHuth侯斯 feud.Barony = &侯斯HuthBarony{}

func init() {
	f := BHuth侯斯.(*侯斯HuthBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "huth",
		TitleName: "侯斯",
		TitleCode: "b_huth",
	}
}
