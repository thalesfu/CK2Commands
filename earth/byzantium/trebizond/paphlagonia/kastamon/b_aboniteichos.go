package kastamon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿沃努蒂霍斯AboniteichosBarony struct {
	feud.BaseBarony
}

var BAboniteichos阿沃努蒂霍斯 feud.Barony = &阿沃努蒂霍斯AboniteichosBarony{}

func init() {
    f := BAboniteichos阿沃努蒂霍斯.(*阿沃努蒂霍斯AboniteichosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aboniteichos",
		TitleName: "阿沃努蒂霍斯",
		TitleCode: "b_aboniteichos",
	}
}
