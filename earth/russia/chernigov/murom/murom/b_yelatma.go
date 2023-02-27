package murom

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 叶拉季马YelatmaBarony struct {
	feud.BaseBarony
}

var BYelatma叶拉季马 feud.Barony = &叶拉季马YelatmaBarony{}

func init() {
    f := BYelatma叶拉季马.(*叶拉季马YelatmaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yelatma",
		TitleName: "叶拉季马",
		TitleCode: "b_yelatma",
	}
}
