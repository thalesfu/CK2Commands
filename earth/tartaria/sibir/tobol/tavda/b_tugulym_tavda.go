package tavda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图古雷姆Tugulym_tavdaBarony struct {
	feud.BaseBarony
}

var BTugulym_tavda图古雷姆 feud.Barony = &图古雷姆Tugulym_tavdaBarony{}

func init() {
    f := BTugulym_tavda图古雷姆.(*图古雷姆Tugulym_tavdaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tugulym_tavda",
		TitleName: "图古雷姆",
		TitleCode: "b_tugulym_tavda",
	}
}
