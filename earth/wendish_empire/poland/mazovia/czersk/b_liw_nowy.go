package czersk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 新利夫Liw_nowyBarony struct {
	feud.BaseBarony
}

var BLiw_nowy新利夫 feud.Barony = &新利夫Liw_nowyBarony{}

func init() {
    f := BLiw_nowy新利夫.(*新利夫Liw_nowyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "liw_nowy",
		TitleName: "新利夫",
		TitleCode: "b_liw_nowy",
	}
}
