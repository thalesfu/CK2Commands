package crimea

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡拉苏夫巴扎尔QarasuvbazarBarony struct {
	feud.BaseBarony
}

var BQarasuvbazar卡拉苏夫巴扎尔 feud.Barony = &卡拉苏夫巴扎尔QarasuvbazarBarony{}

func init() {
    f := BQarasuvbazar卡拉苏夫巴扎尔.(*卡拉苏夫巴扎尔QarasuvbazarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qarasuvbazar",
		TitleName: "卡拉苏夫巴扎尔",
		TitleCode: "b_qarasuvbazar",
	}
}
