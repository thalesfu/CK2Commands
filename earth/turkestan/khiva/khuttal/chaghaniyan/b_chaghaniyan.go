package chaghaniyan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斫汗那ChaghaniyanBarony struct {
	feud.BaseBarony
}

var BChaghaniyan斫汗那 feud.Barony = &斫汗那ChaghaniyanBarony{}

func init() {
	f := BChaghaniyan斫汗那.(*斫汗那ChaghaniyanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chaghaniyan",
		TitleName: "斫汗那",
		TitleCode: "b_chaghaniyan",
	}
}
