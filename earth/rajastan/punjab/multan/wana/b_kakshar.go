package wana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伽叉刹KaksharBarony struct {
	feud.BaseBarony
}

var BKakshar伽叉刹 feud.Barony = &伽叉刹KaksharBarony{}

func init() {
    f := BKakshar伽叉刹.(*伽叉刹KaksharBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kakshar",
		TitleName: "伽叉刹",
		TitleCode: "b_kakshar",
	}
}
