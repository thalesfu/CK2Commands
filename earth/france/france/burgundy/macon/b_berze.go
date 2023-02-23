package macon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝尔泽BerzeBarony struct {
	feud.BaseBarony
}

var BBerze贝尔泽 feud.Barony = &贝尔泽BerzeBarony{}

func init() {
	f := BBerze贝尔泽.(*贝尔泽BerzeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "berze",
		TitleName: "贝尔泽",
		TitleCode: "b_berze",
	}
}
