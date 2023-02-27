package kurmanchal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗耳契吒RanikhetBarony struct {
	feud.BaseBarony
}

var BRanikhet罗耳契吒 feud.Barony = &罗耳契吒RanikhetBarony{}

func init() {
    f := BRanikhet罗耳契吒.(*罗耳契吒RanikhetBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ranikhet",
		TitleName: "罗耳契吒",
		TitleCode: "b_ranikhet",
	}
}
