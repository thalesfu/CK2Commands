package georgia

import (
	"github.com/thalesfu/CK2Commands/earth/byzantium/georgia/abkhazia"
	"github.com/thalesfu/CK2Commands/earth/byzantium/georgia/kakheti"
	"github.com/thalesfu/CK2Commands/earth/byzantium/georgia/kartli"
	"github.com/thalesfu/CK2Commands/earth/byzantium/georgia/tao"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GeorgiaKingdom interface {
	feud.Kingdom
	DAbkhazia阿布哈兹() abkhazia.AbkhaziaDuke
	DKakheti卡赫季() kakheti.KakhetiDuke
	DKartli卡尔特利() kartli.KartliDuke
	DTao陶() tao.TaoDuke
}

type 格鲁吉亚GeorgiaKingdom struct {
	feud.BaseKingdom
	阿布哈兹Abkhazia abkhazia.AbkhaziaDuke
	卡赫季Kakheti   kakheti.KakhetiDuke
	卡尔特利Kartli   kartli.KartliDuke
	陶Tao         tao.TaoDuke
}

func (k *格鲁吉亚GeorgiaKingdom) DAbkhazia阿布哈兹() abkhazia.AbkhaziaDuke {
	return k.阿布哈兹Abkhazia
}

func (k *格鲁吉亚GeorgiaKingdom) DKakheti卡赫季() kakheti.KakhetiDuke {
	return k.卡赫季Kakheti
}

func (k *格鲁吉亚GeorgiaKingdom) DKartli卡尔特利() kartli.KartliDuke {
	return k.卡尔特利Kartli
}

func (k *格鲁吉亚GeorgiaKingdom) DTao陶() tao.TaoDuke {
	return k.陶Tao
}

var KGeorgia格鲁吉亚 GeorgiaKingdom = &格鲁吉亚GeorgiaKingdom{}

func init() {
	f := KGeorgia格鲁吉亚.(*格鲁吉亚GeorgiaKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "georgia",
		TitleName: "格鲁吉亚",
		TitleCode: "k_georgia",
		Dukes:     map[string]feud.Duke{},
	}

	f.阿布哈兹Abkhazia = abkhazia.DAbkhazia阿布哈兹
	f.阿布哈兹Abkhazia.SetParent(f)

	f.卡赫季Kakheti = kakheti.DKakheti卡赫季
	f.卡赫季Kakheti.SetParent(f)

	f.卡尔特利Kartli = kartli.DKartli卡尔特利
	f.卡尔特利Kartli.SetParent(f)

	f.陶Tao = tao.DTao陶
	f.陶Tao.SetParent(f)

}
