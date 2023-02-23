package vaspurakan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈达马克特HadamakertBarony struct {
	feud.BaseBarony
}

var BHadamakert哈达马克特 feud.Barony = &哈达马克特HadamakertBarony{}

func init() {
	f := BHadamakert哈达马克特.(*哈达马克特HadamakertBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hadamakert",
		TitleName: "哈达马克特",
		TitleCode: "b_hadamakert",
	}
}
