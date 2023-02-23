package beirut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泰隆窟CavedetyronBarony struct {
	feud.BaseBarony
}

var BCavedetyron泰隆窟 feud.Barony = &泰隆窟CavedetyronBarony{}

func init() {
	f := BCavedetyron泰隆窟.(*泰隆窟CavedetyronBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cavedetyron",
		TitleName: "泰隆窟",
		TitleCode: "b_cavedetyron",
	}
}
