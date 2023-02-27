package vladimir

import (
	"github.com/thalesfu/CK2Commands/earth/russia/vladimir/moskva"
	"github.com/thalesfu/CK2Commands/earth/russia/vladimir/rostov"
	"github.com/thalesfu/CK2Commands/earth/russia/vladimir/tver"
	"github.com/thalesfu/CK2Commands/earth/russia/vladimir/vladimir"
	"github.com/thalesfu/CK2Commands/earth/russia/vladimir/yaroslavl"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VladimirKingdom interface {
    feud.Kingdom
    DMoskva莫斯科() 	moskva.MoskvaDuke
    DRostov罗斯托夫() 	rostov.RostovDuke
    DTver特维尔() 	tver.TverDuke
    DVladimir弗拉基米尔() 	vladimir.VladimirDuke
    DYaroslavl沃洛格达() 	yaroslavl.YaroslavlDuke
}

type 弗拉基米尔VladimirKingdom struct {
	feud.BaseKingdom
	莫斯科Moskva 	moskva.MoskvaDuke
	罗斯托夫Rostov 	rostov.RostovDuke
	特维尔Tver 	tver.TverDuke
	弗拉基米尔Vladimir 	vladimir.VladimirDuke
	沃洛格达Yaroslavl 	yaroslavl.YaroslavlDuke
}

func (k *弗拉基米尔VladimirKingdom) DMoskva莫斯科() moskva.MoskvaDuke {
	return k.莫斯科Moskva
}
    
func (k *弗拉基米尔VladimirKingdom) DRostov罗斯托夫() rostov.RostovDuke {
	return k.罗斯托夫Rostov
}
    
func (k *弗拉基米尔VladimirKingdom) DTver特维尔() tver.TverDuke {
	return k.特维尔Tver
}
    
func (k *弗拉基米尔VladimirKingdom) DVladimir弗拉基米尔() vladimir.VladimirDuke {
	return k.弗拉基米尔Vladimir
}
    
func (k *弗拉基米尔VladimirKingdom) DYaroslavl沃洛格达() yaroslavl.YaroslavlDuke {
	return k.沃洛格达Yaroslavl
}
    
var KVladimir弗拉基米尔 VladimirKingdom = &弗拉基米尔VladimirKingdom{}

func init() {
	f := KVladimir弗拉基米尔.(*弗拉基米尔VladimirKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "vladimir",
		TitleName: "弗拉基米尔",
		TitleCode: "k_vladimir",
		Dukes:  map[string]feud.Duke{},
	}

	f.莫斯科Moskva = moskva.DMoskva莫斯科
	f.莫斯科Moskva.SetParent(f)
	
	f.罗斯托夫Rostov = rostov.DRostov罗斯托夫
	f.罗斯托夫Rostov.SetParent(f)
	
	f.特维尔Tver = tver.DTver特维尔
	f.特维尔Tver.SetParent(f)
	
	f.弗拉基米尔Vladimir = vladimir.DVladimir弗拉基米尔
	f.弗拉基米尔Vladimir.SetParent(f)
	
	f.沃洛格达Yaroslavl = yaroslavl.DYaroslavl沃洛格达
	f.沃洛格达Yaroslavl.SetParent(f)
	
}
