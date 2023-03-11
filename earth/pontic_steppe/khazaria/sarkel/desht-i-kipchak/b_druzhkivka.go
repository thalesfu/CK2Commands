package desht-i-kipchak

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德鲁日基夫卡DruzhkivkaBarony struct {
	feud.BaseBarony
}

var BDruzhkivka德鲁日基夫卡 feud.Barony = &德鲁日基夫卡DruzhkivkaBarony{}

func init() {
    f := BDruzhkivka德鲁日基夫卡.(*德鲁日基夫卡DruzhkivkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "druzhkivka",
		TitleName: "德鲁日基夫卡",
		TitleCode: "b_druzhkivka",
	}
}
