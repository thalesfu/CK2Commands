package asni

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 湿婆罗阇补罗ShivrajpurBarony struct {
	feud.BaseBarony
}

var BShivrajpur湿婆罗阇补罗 feud.Barony = &湿婆罗阇补罗ShivrajpurBarony{}

func init() {
	f := BShivrajpur湿婆罗阇补罗.(*湿婆罗阇补罗ShivrajpurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shivrajpur",
		TitleName: "湿婆罗阇补罗",
		TitleCode: "b_shivrajpur",
	}
}
