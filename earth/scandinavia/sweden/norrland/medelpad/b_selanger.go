package medelpad

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 谢勒恩耶SelangerBarony struct {
	feud.BaseBarony
}

var BSelanger谢勒恩耶 feud.Barony = &谢勒恩耶SelangerBarony{}

func init() {
    f := BSelanger谢勒恩耶.(*谢勒恩耶SelangerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "selanger",
		TitleName: "谢勒恩耶",
		TitleCode: "b_selanger",
	}
}
