package kanara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伽陀利迦KadarikaBarony struct {
	feud.BaseBarony
}

var BKadarika伽陀利迦 feud.Barony = &伽陀利迦KadarikaBarony{}

func init() {
    f := BKadarika伽陀利迦.(*伽陀利迦KadarikaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kadarika",
		TitleName: "伽陀利迦",
		TitleCode: "b_kadarika",
	}
}
