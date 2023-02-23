package bornholm

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯瓦讷克SvanekeBarony struct {
	feud.BaseBarony
}

var BSvaneke斯瓦讷克 feud.Barony = &斯瓦讷克SvanekeBarony{}

func init() {
	f := BSvaneke斯瓦讷克.(*斯瓦讷克SvanekeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "svaneke",
		TitleName: "斯瓦讷克",
		TitleCode: "b_svaneke",
	}
}
