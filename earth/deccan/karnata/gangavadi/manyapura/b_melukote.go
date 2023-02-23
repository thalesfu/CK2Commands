package manyapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫鲁克特MelukoteBarony struct {
	feud.BaseBarony
}

var BMelukote莫鲁克特 feud.Barony = &莫鲁克特MelukoteBarony{}

func init() {
	f := BMelukote莫鲁克特.(*莫鲁克特MelukoteBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "melukote",
		TitleName: "莫鲁克特",
		TitleCode: "b_melukote",
	}
}
