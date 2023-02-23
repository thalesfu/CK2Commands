package delhi

import (
	"github.com/thalesfu/CK2Commands/earth/rajastan/delhi/haritanaka"
	"github.com/thalesfu/CK2Commands/earth/rajastan/delhi/kuru"
	"github.com/thalesfu/CK2Commands/earth/rajastan/delhi/mathura"
	"github.com/thalesfu/CK2Commands/earth/rajastan/delhi/vodamayutja"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DelhiKingdom interface {
	feud.Kingdom
	DHaritanaka诃利多那迦() haritanaka.HaritanakaDuke
	DKuru俱卢() kuru.KuruDuke
	DMathura秣菟罗() mathura.MathuraDuke
	DVodamayutja菩陀摩瑜多() vodamayutja.VodamayutjaDuke
}

type 德里DelhiKingdom struct {
	feud.BaseKingdom
	诃利多那迦Haritanaka  haritanaka.HaritanakaDuke
	俱卢Kuru           kuru.KuruDuke
	秣菟罗Mathura       mathura.MathuraDuke
	菩陀摩瑜多Vodamayutja vodamayutja.VodamayutjaDuke
}

func (k *德里DelhiKingdom) DHaritanaka诃利多那迦() haritanaka.HaritanakaDuke {
	return k.诃利多那迦Haritanaka
}

func (k *德里DelhiKingdom) DKuru俱卢() kuru.KuruDuke {
	return k.俱卢Kuru
}

func (k *德里DelhiKingdom) DMathura秣菟罗() mathura.MathuraDuke {
	return k.秣菟罗Mathura
}

func (k *德里DelhiKingdom) DVodamayutja菩陀摩瑜多() vodamayutja.VodamayutjaDuke {
	return k.菩陀摩瑜多Vodamayutja
}

var KDelhi德里 DelhiKingdom = &德里DelhiKingdom{}

func init() {
	f := KDelhi德里.(*德里DelhiKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "delhi",
		TitleName: "德里",
		TitleCode: "k_delhi",
		Dukes:     map[string]feud.Duke{},
	}

	f.诃利多那迦Haritanaka = haritanaka.DHaritanaka诃利多那迦
	f.诃利多那迦Haritanaka.SetParent(f)

	f.俱卢Kuru = kuru.DKuru俱卢
	f.俱卢Kuru.SetParent(f)

	f.秣菟罗Mathura = mathura.DMathura秣菟罗
	f.秣菟罗Mathura.SetParent(f)

	f.菩陀摩瑜多Vodamayutja = vodamayutja.DVodamayutja菩陀摩瑜多
	f.菩陀摩瑜多Vodamayutja.SetParent(f)

}
