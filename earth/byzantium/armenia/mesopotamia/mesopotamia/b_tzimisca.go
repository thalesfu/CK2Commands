package mesopotamia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 兹米沙TzimiscaBarony struct {
	feud.BaseBarony
}

var BTzimisca兹米沙 feud.Barony = &兹米沙TzimiscaBarony{}

func init() {
	f := BTzimisca兹米沙.(*兹米沙TzimiscaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tzimisca",
		TitleName: "兹米沙",
		TitleCode: "b_tzimisca",
	}
}
