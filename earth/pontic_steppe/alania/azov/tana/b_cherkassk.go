package tana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 切尔卡斯克CherkasskBarony struct {
	feud.BaseBarony
}

var BCherkassk切尔卡斯克 feud.Barony = &切尔卡斯克CherkasskBarony{}

func init() {
    f := BCherkassk切尔卡斯克.(*切尔卡斯克CherkasskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cherkassk",
		TitleName: "切尔卡斯克",
		TitleCode: "b_cherkassk",
	}
}
