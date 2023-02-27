package ravenna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 切塞纳蒂科CesenaticoBarony struct {
	feud.BaseBarony
}

var BCesenatico切塞纳蒂科 feud.Barony = &切塞纳蒂科CesenaticoBarony{}

func init() {
    f := BCesenatico切塞纳蒂科.(*切塞纳蒂科CesenaticoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cesenatico",
		TitleName: "切塞纳蒂科",
		TitleCode: "b_cesenatico",
	}
}
