package simaramapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 计舍梨耶KesariaBarony struct {
	feud.BaseBarony
}

var BKesaria计舍梨耶 feud.Barony = &计舍梨耶KesariaBarony{}

func init() {
    f := BKesaria计舍梨耶.(*计舍梨耶KesariaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kesaria",
		TitleName: "计舍梨耶",
		TitleCode: "b_kesaria",
	}
}
