package qaidam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 小柴达木XiaochaidamuBarony struct {
	feud.BaseBarony
}

var BXiaochaidamu小柴达木 feud.Barony = &小柴达木XiaochaidamuBarony{}

func init() {
	f := BXiaochaidamu小柴达木.(*小柴达木XiaochaidamuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "xiaochaidamu",
		TitleName: "小柴达木",
		TitleCode: "b_xiaochaidamu",
	}
}
