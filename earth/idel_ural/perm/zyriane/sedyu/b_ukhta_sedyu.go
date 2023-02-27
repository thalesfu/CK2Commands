package sedyu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌赫塔Ukhta_sedyuBarony struct {
	feud.BaseBarony
}

var BUkhta_sedyu乌赫塔 feud.Barony = &乌赫塔Ukhta_sedyuBarony{}

func init() {
    f := BUkhta_sedyu乌赫塔.(*乌赫塔Ukhta_sedyuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ukhta_sedyu",
		TitleName: "乌赫塔",
		TitleCode: "b_ukhta_sedyu",
	}
}
