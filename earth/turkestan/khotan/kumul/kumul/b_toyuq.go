package kumul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吐峪沟ToyuqBarony struct {
	feud.BaseBarony
}

var BToyuq吐峪沟 feud.Barony = &吐峪沟ToyuqBarony{}

func init() {
	f := BToyuq吐峪沟.(*吐峪沟ToyuqBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "toyuq",
		TitleName: "吐峪沟",
		TitleCode: "b_toyuq",
	}
}
