package meath

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奇纳赫塔CiannachtaBarony struct {
	feud.BaseBarony
}

var BCiannachta奇纳赫塔 feud.Barony = &奇纳赫塔CiannachtaBarony{}

func init() {
    f := BCiannachta奇纳赫塔.(*奇纳赫塔CiannachtaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ciannachta",
		TitleName: "奇纳赫塔",
		TitleCode: "b_ciannachta",
	}
}
