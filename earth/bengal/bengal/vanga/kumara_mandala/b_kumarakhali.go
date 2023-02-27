package kumara_mandala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鸠摩罗佉梨KumarakhaliBarony struct {
	feud.BaseBarony
}

var BKumarakhali鸠摩罗佉梨 feud.Barony = &鸠摩罗佉梨KumarakhaliBarony{}

func init() {
    f := BKumarakhali鸠摩罗佉梨.(*鸠摩罗佉梨KumarakhaliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kumarakhali",
		TitleName: "鸠摩罗佉梨",
		TitleCode: "b_kumarakhali",
	}
}
