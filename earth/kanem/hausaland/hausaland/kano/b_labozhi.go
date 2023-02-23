package kano

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉博日LabozhiBarony struct {
	feud.BaseBarony
}

var BLabozhi拉博日 feud.Barony = &拉博日LabozhiBarony{}

func init() {
	f := BLabozhi拉博日.(*拉博日LabozhiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "labozhi",
		TitleName: "拉博日",
		TitleCode: "b_labozhi",
	}
}
