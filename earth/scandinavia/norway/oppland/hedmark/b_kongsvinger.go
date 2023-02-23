package hedmark

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 孔斯温厄尔KongsvingerBarony struct {
	feud.BaseBarony
}

var BKongsvinger孔斯温厄尔 feud.Barony = &孔斯温厄尔KongsvingerBarony{}

func init() {
	f := BKongsvinger孔斯温厄尔.(*孔斯温厄尔KongsvingerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kongsvinger",
		TitleName: "孔斯温厄尔",
		TitleCode: "b_kongsvinger",
	}
}
