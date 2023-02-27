package pisa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉廖GiglioBarony struct {
	feud.BaseBarony
}

var BGiglio吉廖 feud.Barony = &吉廖GiglioBarony{}

func init() {
    f := BGiglio吉廖.(*吉廖GiglioBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "giglio",
		TitleName: "吉廖",
		TitleCode: "b_giglio",
	}
}
