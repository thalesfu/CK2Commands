package khovsgol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博托戈尔BotogolBarony struct {
	feud.BaseBarony
}

var BBotogol博托戈尔 feud.Barony = &博托戈尔BotogolBarony{}

func init() {
	f := BBotogol博托戈尔.(*博托戈尔BotogolBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "botogol",
		TitleName: "博托戈尔",
		TitleCode: "b_botogol",
	}
}
