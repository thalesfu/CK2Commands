package keltma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯莫良卡SmolyankaBarony struct {
	feud.BaseBarony
}

var BSmolyanka斯莫良卡 feud.Barony = &斯莫良卡SmolyankaBarony{}

func init() {
    f := BSmolyanka斯莫良卡.(*斯莫良卡SmolyankaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "smolyanka",
		TitleName: "斯莫良卡",
		TitleCode: "b_smolyanka",
	}
}
