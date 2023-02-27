package gwalior

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 娑婆罗姞利呬SabalgarhBarony struct {
	feud.BaseBarony
}

var BSabalgarh娑婆罗姞利呬 feud.Barony = &娑婆罗姞利呬SabalgarhBarony{}

func init() {
    f := BSabalgarh娑婆罗姞利呬.(*娑婆罗姞利呬SabalgarhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sabalgarh",
		TitleName: "娑婆罗姞利呬",
		TitleCode: "b_sabalgarh",
	}
}
