package ossory

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿哈波AghaboeBarony struct {
	feud.BaseBarony
}

var BAghaboe阿哈波 feud.Barony = &阿哈波AghaboeBarony{}

func init() {
	f := BAghaboe阿哈波.(*阿哈波AghaboeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aghaboe",
		TitleName: "阿哈波",
		TitleCode: "b_aghaboe",
	}
}
