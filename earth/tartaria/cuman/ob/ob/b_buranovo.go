package ob

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布拉诺沃BuranovoBarony struct {
	feud.BaseBarony
}

var BBuranovo布拉诺沃 feud.Barony = &布拉诺沃BuranovoBarony{}

func init() {
    f := BBuranovo布拉诺沃.(*布拉诺沃BuranovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "buranovo",
		TitleName: "布拉诺沃",
		TitleCode: "b_buranovo",
	}
}
