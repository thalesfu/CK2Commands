package kamatapur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 毗多姞利呬BhitagarhBarony struct {
	feud.BaseBarony
}

var BBhitagarh毗多姞利呬 feud.Barony = &毗多姞利呬BhitagarhBarony{}

func init() {
	f := BBhitagarh毗多姞利呬.(*毗多姞利呬BhitagarhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bhitagarh",
		TitleName: "毗多姞利呬",
		TitleCode: "b_bhitagarh",
	}
}
