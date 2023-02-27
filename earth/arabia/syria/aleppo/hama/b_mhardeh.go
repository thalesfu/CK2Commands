package hama

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 玛哈德MhardehBarony struct {
	feud.BaseBarony
}

var BMhardeh玛哈德 feud.Barony = &玛哈德MhardehBarony{}

func init() {
    f := BMhardeh玛哈德.(*玛哈德MhardehBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mhardeh",
		TitleName: "玛哈德",
		TitleCode: "b_mhardeh",
	}
}
