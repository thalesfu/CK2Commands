package sarakhs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尤扎拉YuzlarBarony struct {
	feud.BaseBarony
}

var BYuzlar尤扎拉 feud.Barony = &尤扎拉YuzlarBarony{}

func init() {
	f := BYuzlar尤扎拉.(*尤扎拉YuzlarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yuzlar",
		TitleName: "尤扎拉",
		TitleCode: "b_yuzlar",
	}
}
