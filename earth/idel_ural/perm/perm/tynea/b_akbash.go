package tynea

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿克巴什AkbashBarony struct {
	feud.BaseBarony
}

var BAkbash阿克巴什 feud.Barony = &阿克巴什AkbashBarony{}

func init() {
    f := BAkbash阿克巴什.(*阿克巴什AkbashBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "akbash",
		TitleName: "阿克巴什",
		TitleCode: "b_akbash",
	}
}
