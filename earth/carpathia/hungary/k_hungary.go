package hungary

import (
	"github.com/thalesfu/CK2Commands/earth/carpathia/hungary/esztergom"
	"github.com/thalesfu/CK2Commands/earth/carpathia/hungary/nyitra"
	"github.com/thalesfu/CK2Commands/earth/carpathia/hungary/pecs"
	"github.com/thalesfu/CK2Commands/earth/carpathia/hungary/pest"
	"github.com/thalesfu/CK2Commands/earth/carpathia/hungary/temes"
	"github.com/thalesfu/CK2Commands/earth/carpathia/hungary/transylvania"
	"github.com/thalesfu/CK2Commands/earth/carpathia/hungary/ungvar"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type HungaryKingdom interface {
    feud.Kingdom
    DEsztergom埃斯泰尔戈姆() 	esztergom.EsztergomDuke
    DNyitra尼特劳() 	nyitra.NyitraDuke
    DPecs佩奇() 	pecs.PecsDuke
    DPest佩斯() 	pest.PestDuke
    DTemes泰梅什() 	temes.TemesDuke
    DTransylvania特兰西瓦尼亚() 	transylvania.TransylvaniaDuke
    DUngvar翁格瓦尔() 	ungvar.UngvarDuke
}

type 匈牙利HungaryKingdom struct {
	feud.BaseKingdom
	埃斯泰尔戈姆Esztergom 	esztergom.EsztergomDuke
	尼特劳Nyitra 	nyitra.NyitraDuke
	佩奇Pecs 	pecs.PecsDuke
	佩斯Pest 	pest.PestDuke
	泰梅什Temes 	temes.TemesDuke
	特兰西瓦尼亚Transylvania 	transylvania.TransylvaniaDuke
	翁格瓦尔Ungvar 	ungvar.UngvarDuke
}

func (k *匈牙利HungaryKingdom) DEsztergom埃斯泰尔戈姆() esztergom.EsztergomDuke {
	return k.埃斯泰尔戈姆Esztergom
}
    
func (k *匈牙利HungaryKingdom) DNyitra尼特劳() nyitra.NyitraDuke {
	return k.尼特劳Nyitra
}
    
func (k *匈牙利HungaryKingdom) DPecs佩奇() pecs.PecsDuke {
	return k.佩奇Pecs
}
    
func (k *匈牙利HungaryKingdom) DPest佩斯() pest.PestDuke {
	return k.佩斯Pest
}
    
func (k *匈牙利HungaryKingdom) DTemes泰梅什() temes.TemesDuke {
	return k.泰梅什Temes
}
    
func (k *匈牙利HungaryKingdom) DTransylvania特兰西瓦尼亚() transylvania.TransylvaniaDuke {
	return k.特兰西瓦尼亚Transylvania
}
    
func (k *匈牙利HungaryKingdom) DUngvar翁格瓦尔() ungvar.UngvarDuke {
	return k.翁格瓦尔Ungvar
}
    
var KHungary匈牙利 HungaryKingdom = &匈牙利HungaryKingdom{}

func init() {
	f := KHungary匈牙利.(*匈牙利HungaryKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "hungary",
		TitleName: "匈牙利",
		TitleCode: "k_hungary",
		Dukes:  map[string]feud.Duke{},
	}

	f.埃斯泰尔戈姆Esztergom = esztergom.DEsztergom埃斯泰尔戈姆
	f.埃斯泰尔戈姆Esztergom.SetParent(f)
	
	f.尼特劳Nyitra = nyitra.DNyitra尼特劳
	f.尼特劳Nyitra.SetParent(f)
	
	f.佩奇Pecs = pecs.DPecs佩奇
	f.佩奇Pecs.SetParent(f)
	
	f.佩斯Pest = pest.DPest佩斯
	f.佩斯Pest.SetParent(f)
	
	f.泰梅什Temes = temes.DTemes泰梅什
	f.泰梅什Temes.SetParent(f)
	
	f.特兰西瓦尼亚Transylvania = transylvania.DTransylvania特兰西瓦尼亚
	f.特兰西瓦尼亚Transylvania.SetParent(f)
	
	f.翁格瓦尔Ungvar = ungvar.DUngvar翁格瓦尔
	f.翁格瓦尔Ungvar.SetParent(f)
	
}
