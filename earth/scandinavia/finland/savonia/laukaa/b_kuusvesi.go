package laukaa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库斯韦西KuusvesiBarony struct {
	feud.BaseBarony
}

var BKuusvesi库斯韦西 feud.Barony = &库斯韦西KuusvesiBarony{}

func init() {
	f := BKuusvesi库斯韦西.(*库斯韦西KuusvesiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kuusvesi",
		TitleName: "库斯韦西",
		TitleCode: "b_kuusvesi",
	}
}
