package rhodos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈拉基HarakiBarony struct {
	feud.BaseBarony
}

var BHaraki哈拉基 feud.Barony = &哈拉基HarakiBarony{}

func init() {
	f := BHaraki哈拉基.(*哈拉基HarakiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "haraki",
		TitleName: "哈拉基",
		TitleCode: "b_haraki",
	}
}
