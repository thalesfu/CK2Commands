package methone

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格里采纳GritzenaBarony struct {
	feud.BaseBarony
}

var BGritzena格里采纳 feud.Barony = &格里采纳GritzenaBarony{}

func init() {
    f := BGritzena格里采纳.(*格里采纳GritzenaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gritzena",
		TitleName: "格里采纳",
		TitleCode: "b_gritzena",
	}
}
