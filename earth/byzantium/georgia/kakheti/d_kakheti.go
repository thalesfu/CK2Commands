package kakheti

import (
	"github.com/thalesfu/CK2Commands/earth/byzantium/georgia/kakheti/albania"
	"github.com/thalesfu/CK2Commands/earth/byzantium/georgia/kakheti/kakheti"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KakhetiDuke interface {
	feud.Duke
	CAlbania阿尔巴尼亚() albania.AlbaniaCounty
	CKakheti卡赫季() kakheti.KakhetiCounty
}

type 卡赫季KakhetiDuke struct {
	feud.BaseDuke
	阿尔巴尼亚Albania albania.AlbaniaCounty
	卡赫季Kakheti   kakheti.KakhetiCounty
}

func (d *卡赫季KakhetiDuke) CAlbania阿尔巴尼亚() albania.AlbaniaCounty {
	return d.阿尔巴尼亚Albania
}

func (d *卡赫季KakhetiDuke) CKakheti卡赫季() kakheti.KakhetiCounty {
	return d.卡赫季Kakheti
}

var DKakheti卡赫季 KakhetiDuke = &卡赫季KakhetiDuke{}

func init() {
	f := DKakheti卡赫季.(*卡赫季KakhetiDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "kakheti",
		TitleName: "卡赫季",
		TitleCode: "d_kakheti",
		Counties:  map[string]feud.County{},
	}

	f.阿尔巴尼亚Albania = albania.CAlbania阿尔巴尼亚
	f.阿尔巴尼亚Albania.SetParent(f)

	f.卡赫季Kakheti = kakheti.CKakheti卡赫季
	f.卡赫季Kakheti.SetParent(f)

}
