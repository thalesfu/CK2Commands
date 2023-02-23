package gyantse

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 白朗BainangBarony struct {
	feud.BaseBarony
}

var BBainang白朗 feud.Barony = &白朗BainangBarony{}

func init() {
	f := BBainang白朗.(*白朗BainangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bainang",
		TitleName: "白朗",
		TitleCode: "b_bainang",
	}
}
