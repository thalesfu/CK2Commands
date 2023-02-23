package karnten

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格赖芬堡GreifenburgBarony struct {
	feud.BaseBarony
}

var BGreifenburg格赖芬堡 feud.Barony = &格赖芬堡GreifenburgBarony{}

func init() {
	f := BGreifenburg格赖芬堡.(*格赖芬堡GreifenburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "greifenburg",
		TitleName: "格赖芬堡",
		TitleCode: "b_greifenburg",
	}
}
