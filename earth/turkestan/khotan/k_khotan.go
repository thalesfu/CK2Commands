package khotan

import (
	"github.com/thalesfu/CK2Commands/earth/turkestan/khotan/karashar"
	"github.com/thalesfu/CK2Commands/earth/turkestan/khotan/kashgar"
	"github.com/thalesfu/CK2Commands/earth/turkestan/khotan/khotan"
	"github.com/thalesfu/CK2Commands/earth/turkestan/khotan/kumul"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KhotanKingdom interface {
    feud.Kingdom
    DKarashar喀喇沙尔() 	karashar.KarasharDuke
    DKashgar可失合儿() 	kashgar.KashgarDuke
    DKhotan于阗() 	khotan.KhotanDuke
    DKumul哈密() 	kumul.KumulDuke
}

type 于阗KhotanKingdom struct {
	feud.BaseKingdom
	喀喇沙尔Karashar 	karashar.KarasharDuke
	可失合儿Kashgar 	kashgar.KashgarDuke
	于阗Khotan 	khotan.KhotanDuke
	哈密Kumul 	kumul.KumulDuke
}

func (k *于阗KhotanKingdom) DKarashar喀喇沙尔() karashar.KarasharDuke {
	return k.喀喇沙尔Karashar
}
    
func (k *于阗KhotanKingdom) DKashgar可失合儿() kashgar.KashgarDuke {
	return k.可失合儿Kashgar
}
    
func (k *于阗KhotanKingdom) DKhotan于阗() khotan.KhotanDuke {
	return k.于阗Khotan
}
    
func (k *于阗KhotanKingdom) DKumul哈密() kumul.KumulDuke {
	return k.哈密Kumul
}
    
var KKhotan于阗 KhotanKingdom = &于阗KhotanKingdom{}

func init() {
	f := KKhotan于阗.(*于阗KhotanKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "khotan",
		TitleName: "于阗",
		TitleCode: "k_khotan",
		Dukes:  map[string]feud.Duke{},
	}

	f.喀喇沙尔Karashar = karashar.DKarashar喀喇沙尔
	f.喀喇沙尔Karashar.SetParent(f)
	
	f.可失合儿Kashgar = kashgar.DKashgar可失合儿
	f.可失合儿Kashgar.SetParent(f)
	
	f.于阗Khotan = khotan.DKhotan于阗
	f.于阗Khotan.SetParent(f)
	
	f.哈密Kumul = kumul.DKumul哈密
	f.哈密Kumul.SetParent(f)
	
}
