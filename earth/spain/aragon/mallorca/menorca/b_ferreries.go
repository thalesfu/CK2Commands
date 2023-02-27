package menorca

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 费雷列斯FerreriesBarony struct {
	feud.BaseBarony
}

var BFerreries费雷列斯 feud.Barony = &费雷列斯FerreriesBarony{}

func init() {
    f := BFerreries费雷列斯.(*费雷列斯FerreriesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ferreries",
		TitleName: "费雷列斯",
		TitleCode: "b_ferreries",
	}
}
