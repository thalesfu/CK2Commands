package bhumilka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 菩弥伽BhumilkaBarony struct {
	feud.BaseBarony
}

var BBhumilka菩弥伽 feud.Barony = &菩弥伽BhumilkaBarony{}

func init() {
    f := BBhumilka菩弥伽.(*菩弥伽BhumilkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bhumilka",
		TitleName: "菩弥伽",
		TitleCode: "b_bhumilka",
	}
}
