package ashli

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 亚纳乌尔JanauylBarony struct {
	feud.BaseBarony
}

var BJanauyl亚纳乌尔 feud.Barony = &亚纳乌尔JanauylBarony{}

func init() {
    f := BJanauyl亚纳乌尔.(*亚纳乌尔JanauylBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "janauyl",
		TitleName: "亚纳乌尔",
		TitleCode: "b_janauyl",
	}
}
