package sanaa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 先知舒艾卜峰JabalannabishuaybBarony struct {
	feud.BaseBarony
}

var BJabalannabishuayb先知舒艾卜峰 feud.Barony = &先知舒艾卜峰JabalannabishuaybBarony{}

func init() {
    f := BJabalannabishuayb先知舒艾卜峰.(*先知舒艾卜峰JabalannabishuaybBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jabalannabishuayb",
		TitleName: "先知舒艾卜峰",
		TitleCode: "b_jabalannabishuayb",
	}
}
