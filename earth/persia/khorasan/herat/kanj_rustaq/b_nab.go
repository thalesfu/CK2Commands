package kanj_rustaq

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纳布NabBarony struct {
	feud.BaseBarony
}

var BNab纳布 feud.Barony = &纳布NabBarony{}

func init() {
    f := BNab纳布.(*纳布NabBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nab",
		TitleName: "纳布",
		TitleCode: "b_nab",
	}
}
