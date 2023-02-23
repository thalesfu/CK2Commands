package denia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝尼萨BenissaBarony struct {
	feud.BaseBarony
}

var BBenissa贝尼萨 feud.Barony = &贝尼萨BenissaBarony{}

func init() {
	f := BBenissa贝尼萨.(*贝尼萨BenissaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "benissa",
		TitleName: "贝尼萨",
		TitleCode: "b_benissa",
	}
}
