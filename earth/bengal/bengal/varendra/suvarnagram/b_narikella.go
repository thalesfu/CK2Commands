package suvarnagram

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 那利蓟罗NarikellaBarony struct {
	feud.BaseBarony
}

var BNarikella那利蓟罗 feud.Barony = &那利蓟罗NarikellaBarony{}

func init() {
	f := BNarikella那利蓟罗.(*那利蓟罗NarikellaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "narikella",
		TitleName: "那利蓟罗",
		TitleCode: "b_narikella",
	}
}
