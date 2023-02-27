package hradyzk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科别利亚基KobeliakyBarony struct {
	feud.BaseBarony
}

var BKobeliaky科别利亚基 feud.Barony = &科别利亚基KobeliakyBarony{}

func init() {
    f := BKobeliaky科别利亚基.(*科别利亚基KobeliakyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kobeliaky",
		TitleName: "科别利亚基",
		TitleCode: "b_kobeliaky",
	}
}
