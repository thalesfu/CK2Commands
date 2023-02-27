package otuken

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库图利克KutulikBarony struct {
	feud.BaseBarony
}

var BKutulik库图利克 feud.Barony = &库图利克KutulikBarony{}

func init() {
    f := BKutulik库图利克.(*库图利克KutulikBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kutulik",
		TitleName: "库图利克",
		TitleCode: "b_kutulik",
	}
}
