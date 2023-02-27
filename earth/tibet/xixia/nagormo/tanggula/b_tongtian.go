package tanggula

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 通天TongtianBarony struct {
	feud.BaseBarony
}

var BTongtian通天 feud.Barony = &通天TongtianBarony{}

func init() {
    f := BTongtian通天.(*通天TongtianBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tongtian",
		TitleName: "通天",
		TitleCode: "b_tongtian",
	}
}
