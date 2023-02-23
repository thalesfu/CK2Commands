package somerset

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布里奇沃特BridgwaterBarony struct {
	feud.BaseBarony
}

var BBridgwater布里奇沃特 feud.Barony = &布里奇沃特BridgwaterBarony{}

func init() {
	f := BBridgwater布里奇沃特.(*布里奇沃特BridgwaterBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bridgwater",
		TitleName: "布里奇沃特",
		TitleCode: "b_bridgwater",
	}
}
