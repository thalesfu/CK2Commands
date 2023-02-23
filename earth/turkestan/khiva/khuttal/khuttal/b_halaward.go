package khuttal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈拉沃德HalawardBarony struct {
	feud.BaseBarony
}

var BHalaward哈拉沃德 feud.Barony = &哈拉沃德HalawardBarony{}

func init() {
	f := BHalaward哈拉沃德.(*哈拉沃德HalawardBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "halaward",
		TitleName: "哈拉沃德",
		TitleCode: "b_halaward",
	}
}
