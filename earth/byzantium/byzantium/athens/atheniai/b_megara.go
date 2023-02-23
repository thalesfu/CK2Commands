package atheniai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 墨伽拉MegaraBarony struct {
	feud.BaseBarony
}

var BMegara墨伽拉 feud.Barony = &墨伽拉MegaraBarony{}

func init() {
	f := BMegara墨伽拉.(*墨伽拉MegaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "megara",
		TitleName: "墨伽拉",
		TitleCode: "b_megara",
	}
}
