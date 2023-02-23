package dodugu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纳曼迪吉纳NamandiguinaBarony struct {
	feud.BaseBarony
}

var BNamandiguina纳曼迪吉纳 feud.Barony = &纳曼迪吉纳NamandiguinaBarony{}

func init() {
	f := BNamandiguina纳曼迪吉纳.(*纳曼迪吉纳NamandiguinaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "namandiguina",
		TitleName: "纳曼迪吉纳",
		TitleCode: "b_namandiguina",
	}
}
