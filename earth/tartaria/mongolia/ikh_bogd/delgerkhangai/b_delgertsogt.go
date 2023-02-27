package delgerkhangai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德勒格尔朝格特DelgertsogtBarony struct {
	feud.BaseBarony
}

var BDelgertsogt德勒格尔朝格特 feud.Barony = &德勒格尔朝格特DelgertsogtBarony{}

func init() {
    f := BDelgertsogt德勒格尔朝格特.(*德勒格尔朝格特DelgertsogtBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "delgertsogt",
		TitleName: "德勒格尔朝格特",
		TitleCode: "b_delgertsogt",
	}
}
