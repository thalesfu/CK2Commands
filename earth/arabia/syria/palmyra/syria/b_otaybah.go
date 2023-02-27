package syria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥太巴OtaybahBarony struct {
	feud.BaseBarony
}

var BOtaybah奥太巴 feud.Barony = &奥太巴OtaybahBarony{}

func init() {
    f := BOtaybah奥太巴.(*奥太巴OtaybahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "otaybah",
		TitleName: "奥太巴",
		TitleCode: "b_otaybah",
	}
}
