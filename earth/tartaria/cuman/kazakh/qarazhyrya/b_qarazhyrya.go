package qarazhyrya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 喀拉哲拉QarazhyryaBarony struct {
	feud.BaseBarony
}

var BQarazhyrya喀拉哲拉 feud.Barony = &喀拉哲拉QarazhyryaBarony{}

func init() {
	f := BQarazhyrya喀拉哲拉.(*喀拉哲拉QarazhyryaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qarazhyrya",
		TitleName: "喀拉哲拉",
		TitleCode: "b_qarazhyrya",
	}
}
