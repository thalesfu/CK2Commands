package parnakheta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吠都罗BetoraBarony struct {
	feud.BaseBarony
}

var BBetora吠都罗 feud.Barony = &吠都罗BetoraBarony{}

func init() {
    f := BBetora吠都罗.(*吠都罗BetoraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "betora",
		TitleName: "吠都罗",
		TitleCode: "b_betora",
	}
}
