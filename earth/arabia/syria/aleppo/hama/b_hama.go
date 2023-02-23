package hama

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈马HamaBarony struct {
	feud.BaseBarony
}

var BHama哈马 feud.Barony = &哈马HamaBarony{}

func init() {
	f := BHama哈马.(*哈马HamaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hama",
		TitleName: "哈马",
		TitleCode: "b_hama",
	}
}
