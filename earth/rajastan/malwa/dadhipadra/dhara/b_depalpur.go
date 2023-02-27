package dhara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德波补罗DepalpurBarony struct {
	feud.BaseBarony
}

var BDepalpur德波补罗 feud.Barony = &德波补罗DepalpurBarony{}

func init() {
    f := BDepalpur德波补罗.(*德波补罗DepalpurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "depalpur",
		TitleName: "德波补罗",
		TitleCode: "b_depalpur",
	}
}
