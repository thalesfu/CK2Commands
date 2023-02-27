package angouleme

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉罗什富科LarochefoucauldBarony struct {
	feud.BaseBarony
}

var BLarochefoucauld拉罗什富科 feud.Barony = &拉罗什富科LarochefoucauldBarony{}

func init() {
    f := BLarochefoucauld拉罗什富科.(*拉罗什富科LarochefoucauldBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "larochefoucauld",
		TitleName: "拉罗什富科",
		TitleCode: "b_larochefoucauld",
	}
}
