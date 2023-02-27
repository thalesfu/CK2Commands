package vairagara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 婆度加罗BhandugaraBarony struct {
	feud.BaseBarony
}

var BBhandugara婆度加罗 feud.Barony = &婆度加罗BhandugaraBarony{}

func init() {
    f := BBhandugara婆度加罗.(*婆度加罗BhandugaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bhandugara",
		TitleName: "婆度加罗",
		TitleCode: "b_bhandugara",
	}
}
