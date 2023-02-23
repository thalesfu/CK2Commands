package veglia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弗尔博夫斯科VrbovskoBarony struct {
	feud.BaseBarony
}

var BVrbovsko弗尔博夫斯科 feud.Barony = &弗尔博夫斯科VrbovskoBarony{}

func init() {
	f := BVrbovsko弗尔博夫斯科.(*弗尔博夫斯科VrbovskoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vrbovsko",
		TitleName: "弗尔博夫斯科",
		TitleCode: "b_vrbovsko",
	}
}
