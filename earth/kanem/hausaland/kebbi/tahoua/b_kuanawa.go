package tahoua

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 夸纳瓦KuanawaBarony struct {
	feud.BaseBarony
}

var BKuanawa夸纳瓦 feud.Barony = &夸纳瓦KuanawaBarony{}

func init() {
    f := BKuanawa夸纳瓦.(*夸纳瓦KuanawaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kuanawa",
		TitleName: "夸纳瓦",
		TitleCode: "b_kuanawa",
	}
}
