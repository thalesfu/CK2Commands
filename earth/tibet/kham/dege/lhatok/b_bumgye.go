package lhatok

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 木协BumgyeBarony struct {
	feud.BaseBarony
}

var BBumgye木协 feud.Barony = &木协BumgyeBarony{}

func init() {
	f := BBumgye木协.(*木协BumgyeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bumgye",
		TitleName: "木协",
		TitleCode: "b_bumgye",
	}
}
