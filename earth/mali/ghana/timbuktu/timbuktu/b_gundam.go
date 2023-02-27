package timbuktu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贡达姆GundamBarony struct {
	feud.BaseBarony
}

var BGundam贡达姆 feud.Barony = &贡达姆GundamBarony{}

func init() {
    f := BGundam贡达姆.(*贡达姆GundamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gundam",
		TitleName: "贡达姆",
		TitleCode: "b_gundam",
	}
}
