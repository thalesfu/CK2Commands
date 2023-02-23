package coqen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 措勤CoqenBarony struct {
	feud.BaseBarony
}

var BCoqen措勤 feud.Barony = &措勤CoqenBarony{}

func init() {
	f := BCoqen措勤.(*措勤CoqenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "coqen",
		TitleName: "措勤",
		TitleCode: "b_coqen",
	}
}
