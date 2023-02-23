package wurzburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马林贝格MarienbergBarony struct {
	feud.BaseBarony
}

var BMarienberg马林贝格 feud.Barony = &马林贝格MarienbergBarony{}

func init() {
	f := BMarienberg马林贝格.(*马林贝格MarienbergBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "marienberg",
		TitleName: "马林贝格",
		TitleCode: "b_marienberg",
	}
}
