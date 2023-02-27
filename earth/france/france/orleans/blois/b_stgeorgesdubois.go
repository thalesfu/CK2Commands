package blois

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣乔治迪布瓦StgeorgesduboisBarony struct {
	feud.BaseBarony
}

var BStgeorgesdubois圣乔治迪布瓦 feud.Barony = &圣乔治迪布瓦StgeorgesduboisBarony{}

func init() {
    f := BStgeorgesdubois圣乔治迪布瓦.(*圣乔治迪布瓦StgeorgesduboisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stgeorgesdubois",
		TitleName: "圣乔治迪布瓦",
		TitleCode: "b_stgeorgesdubois",
	}
}
