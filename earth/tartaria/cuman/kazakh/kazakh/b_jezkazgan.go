package kazakh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杰兹卡兹干JezkazganBarony struct {
	feud.BaseBarony
}

var BJezkazgan杰兹卡兹干 feud.Barony = &杰兹卡兹干JezkazganBarony{}

func init() {
    f := BJezkazgan杰兹卡兹干.(*杰兹卡兹干JezkazganBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jezkazgan",
		TitleName: "杰兹卡兹干",
		TitleCode: "b_jezkazgan",
	}
}
