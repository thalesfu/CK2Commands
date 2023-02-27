package amiens

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 亚眠AmiensBarony struct {
	feud.BaseBarony
}

var BAmiens亚眠 feud.Barony = &亚眠AmiensBarony{}

func init() {
    f := BAmiens亚眠.(*亚眠AmiensBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "amiens",
		TitleName: "亚眠",
		TitleCode: "b_amiens",
	}
}
