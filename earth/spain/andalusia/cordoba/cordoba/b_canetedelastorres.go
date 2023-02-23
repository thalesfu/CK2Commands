package cordoba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡涅特德拉斯托雷斯CanetedelastorresBarony struct {
	feud.BaseBarony
}

var BCanetedelastorres卡涅特德拉斯托雷斯 feud.Barony = &卡涅特德拉斯托雷斯CanetedelastorresBarony{}

func init() {
	f := BCanetedelastorres卡涅特德拉斯托雷斯.(*卡涅特德拉斯托雷斯CanetedelastorresBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "canetedelastorres",
		TitleName: "卡涅特德拉斯托雷斯",
		TitleCode: "b_canetedelastorres",
	}
}
