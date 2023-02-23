package teviotdale

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗克斯堡RoxburghBarony struct {
	feud.BaseBarony
}

var BRoxburgh罗克斯堡 feud.Barony = &罗克斯堡RoxburghBarony{}

func init() {
	f := BRoxburgh罗克斯堡.(*罗克斯堡RoxburghBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "roxburgh",
		TitleName: "罗克斯堡",
		TitleCode: "b_roxburgh",
	}
}
