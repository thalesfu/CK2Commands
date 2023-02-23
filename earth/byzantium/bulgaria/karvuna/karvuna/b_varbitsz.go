package karvuna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃尔比察VarbitszBarony struct {
	feud.BaseBarony
}

var BVarbitsz沃尔比察 feud.Barony = &沃尔比察VarbitszBarony{}

func init() {
	f := BVarbitsz沃尔比察.(*沃尔比察VarbitszBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "varbitsz",
		TitleName: "沃尔比察",
		TitleCode: "b_varbitsz",
	}
}
