package mstislavl

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 切切尔斯克ChacherskBarony struct {
	feud.BaseBarony
}

var BChachersk切切尔斯克 feud.Barony = &切切尔斯克ChacherskBarony{}

func init() {
    f := BChachersk切切尔斯克.(*切切尔斯克ChacherskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chachersk",
		TitleName: "切切尔斯克",
		TitleCode: "b_chachersk",
	}
}
