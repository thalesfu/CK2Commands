package varanasi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迦尸毗湿婆那他Kashi_vishwanathBarony struct {
	feud.BaseBarony
}

var BKashi_vishwanath迦尸毗湿婆那他 feud.Barony = &迦尸毗湿婆那他Kashi_vishwanathBarony{}

func init() {
    f := BKashi_vishwanath迦尸毗湿婆那他.(*迦尸毗湿婆那他Kashi_vishwanathBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kashi_vishwanath",
		TitleName: "迦尸毗湿婆那他",
		TitleCode: "b_kashi_vishwanath",
	}
}
