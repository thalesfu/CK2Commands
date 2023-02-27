package makari

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 甘比拉GambilaBarony struct {
	feud.BaseBarony
}

var BGambila甘比拉 feud.Barony = &甘比拉GambilaBarony{}

func init() {
    f := BGambila甘比拉.(*甘比拉GambilaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gambila",
		TitleName: "甘比拉",
		TitleCode: "b_gambila",
	}
}
