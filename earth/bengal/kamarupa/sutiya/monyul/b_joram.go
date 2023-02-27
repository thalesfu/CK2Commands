package monyul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 焦鲁姆JoramBarony struct {
	feud.BaseBarony
}

var BJoram焦鲁姆 feud.Barony = &焦鲁姆JoramBarony{}

func init() {
    f := BJoram焦鲁姆.(*焦鲁姆JoramBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "joram",
		TitleName: "焦鲁姆",
		TitleCode: "b_joram",
	}
}
