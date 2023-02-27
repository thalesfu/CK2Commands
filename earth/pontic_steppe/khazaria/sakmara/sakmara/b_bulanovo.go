package sakmara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布拉诺沃BulanovoBarony struct {
	feud.BaseBarony
}

var BBulanovo布拉诺沃 feud.Barony = &布拉诺沃BulanovoBarony{}

func init() {
    f := BBulanovo布拉诺沃.(*布拉诺沃BulanovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bulanovo",
		TitleName: "布拉诺沃",
		TitleCode: "b_bulanovo",
	}
}
