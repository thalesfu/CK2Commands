package sevede

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克伦斯堡KronsborgBarony struct {
	feud.BaseBarony
}

var BKronsborg克伦斯堡 feud.Barony = &克伦斯堡KronsborgBarony{}

func init() {
	f := BKronsborg克伦斯堡.(*克伦斯堡KronsborgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kronsborg",
		TitleName: "克伦斯堡",
		TitleCode: "b_kronsborg",
	}
}
