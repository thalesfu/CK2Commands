package kara_khoja

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 狐胡HuhuBarony struct {
	feud.BaseBarony
}

var BHuhu狐胡 feud.Barony = &狐胡HuhuBarony{}

func init() {
    f := BHuhu狐胡.(*狐胡HuhuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "huhu",
		TitleName: "狐胡",
		TitleCode: "b_huhu",
	}
}
