package gwalior

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 讷尔沃尔NarwarBarony struct {
	feud.BaseBarony
}

var BNarwar讷尔沃尔 feud.Barony = &讷尔沃尔NarwarBarony{}

func init() {
	f := BNarwar讷尔沃尔.(*讷尔沃尔NarwarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "narwar",
		TitleName: "讷尔沃尔",
		TitleCode: "b_narwar",
	}
}
