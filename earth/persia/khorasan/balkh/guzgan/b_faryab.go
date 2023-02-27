package guzgan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 法里亚布FaryabBarony struct {
	feud.BaseBarony
}

var BFaryab法里亚布 feud.Barony = &法里亚布FaryabBarony{}

func init() {
    f := BFaryab法里亚布.(*法里亚布FaryabBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "faryab",
		TitleName: "法里亚布",
		TitleCode: "b_faryab",
	}
}
