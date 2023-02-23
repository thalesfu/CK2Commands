package ossory

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格兰纳GrannaghBarony struct {
	feud.BaseBarony
}

var BGrannagh格兰纳 feud.Barony = &格兰纳GrannaghBarony{}

func init() {
	f := BGrannagh格兰纳.(*格兰纳GrannaghBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "grannagh",
		TitleName: "格兰纳",
		TitleCode: "b_grannagh",
	}
}
