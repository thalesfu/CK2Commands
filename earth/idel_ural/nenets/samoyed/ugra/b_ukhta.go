package ugra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌赫塔UkhtaBarony struct {
	feud.BaseBarony
}

var BUkhta乌赫塔 feud.Barony = &乌赫塔UkhtaBarony{}

func init() {
    f := BUkhta乌赫塔.(*乌赫塔UkhtaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ukhta",
		TitleName: "乌赫塔",
		TitleCode: "b_ukhta",
	}
}
