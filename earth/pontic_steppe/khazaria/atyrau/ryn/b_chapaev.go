package ryn

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 恰帕耶沃ChapaevBarony struct {
	feud.BaseBarony
}

var BChapaev恰帕耶沃 feud.Barony = &恰帕耶沃ChapaevBarony{}

func init() {
    f := BChapaev恰帕耶沃.(*恰帕耶沃ChapaevBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chapaev",
		TitleName: "恰帕耶沃",
		TitleCode: "b_chapaev",
	}
}
