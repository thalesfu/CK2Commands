package qalqut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尤罗甘UrakamBarony struct {
	feud.BaseBarony
}

var BUrakam尤罗甘 feud.Barony = &尤罗甘UrakamBarony{}

func init() {
    f := BUrakam尤罗甘.(*尤罗甘UrakamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "urakam",
		TitleName: "尤罗甘",
		TitleCode: "b_urakam",
	}
}
