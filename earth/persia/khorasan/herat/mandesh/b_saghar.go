package mandesh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨格哈尔SagharBarony struct {
	feud.BaseBarony
}

var BSaghar萨格哈尔 feud.Barony = &萨格哈尔SagharBarony{}

func init() {
    f := BSaghar萨格哈尔.(*萨格哈尔SagharBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saghar",
		TitleName: "萨格哈尔",
		TitleCode: "b_saghar",
	}
}
