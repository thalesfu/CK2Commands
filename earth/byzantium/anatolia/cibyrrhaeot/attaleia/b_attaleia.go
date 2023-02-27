package attaleia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿塔利亚AttaleiaBarony struct {
	feud.BaseBarony
}

var BAttaleia阿塔利亚 feud.Barony = &阿塔利亚AttaleiaBarony{}

func init() {
    f := BAttaleia阿塔利亚.(*阿塔利亚AttaleiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "attaleia",
		TitleName: "阿塔利亚",
		TitleCode: "b_attaleia",
	}
}
