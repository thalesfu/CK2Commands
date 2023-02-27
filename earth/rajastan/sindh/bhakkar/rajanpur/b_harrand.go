package rajanpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈兰德HarrandBarony struct {
	feud.BaseBarony
}

var BHarrand哈兰德 feud.Barony = &哈兰德HarrandBarony{}

func init() {
    f := BHarrand哈兰德.(*哈兰德HarrandBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "harrand",
		TitleName: "哈兰德",
		TitleCode: "b_harrand",
	}
}
