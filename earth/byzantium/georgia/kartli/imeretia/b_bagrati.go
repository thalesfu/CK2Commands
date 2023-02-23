package imeretia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴格拉季BagratiBarony struct {
	feud.BaseBarony
}

var BBagrati巴格拉季 feud.Barony = &巴格拉季BagratiBarony{}

func init() {
	f := BBagrati巴格拉季.(*巴格拉季BagratiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bagrati",
		TitleName: "巴格拉季",
		TitleCode: "b_bagrati",
	}
}
