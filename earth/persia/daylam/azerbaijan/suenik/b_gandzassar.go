package suenik

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 甘扎萨尔GandzassarBarony struct {
	feud.BaseBarony
}

var BGandzassar甘扎萨尔 feud.Barony = &甘扎萨尔GandzassarBarony{}

func init() {
	f := BGandzassar甘扎萨尔.(*甘扎萨尔GandzassarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gandzassar",
		TitleName: "甘扎萨尔",
		TitleCode: "b_gandzassar",
	}
}
