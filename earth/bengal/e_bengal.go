package bengal

import (
	"github.com/thalesfu/CK2Commands/earth/bengal/bengal"
	"github.com/thalesfu/CK2Commands/earth/bengal/bihar"
	"github.com/thalesfu/CK2Commands/earth/bengal/gondwana"
	"github.com/thalesfu/CK2Commands/earth/bengal/kamarupa"
	"github.com/thalesfu/CK2Commands/earth/bengal/orissa"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BengalEmpire interface {
    feud.Empire
    KBengal旁伽罗() 	bengal.BengalKingdom
    KBihar毗诃罗() 	bihar.BiharKingdom
    KGondwana共陀婆那() 	gondwana.GondwanaKingdom
    KKamarupa迦摩缕波() 	kamarupa.KamarupaKingdom
    KOrissa乌里舍() 	orissa.OrissaKingdom
}

type 旁伽罗帝国BengalEmpire struct {
	feud.BaseEmpire
	旁伽罗Bengal 	bengal.BengalKingdom
	毗诃罗Bihar 	bihar.BiharKingdom
	共陀婆那Gondwana 	gondwana.GondwanaKingdom
	迦摩缕波Kamarupa 	kamarupa.KamarupaKingdom
	乌里舍Orissa 	orissa.OrissaKingdom
}

func (e *旁伽罗帝国BengalEmpire) KBengal旁伽罗() bengal.BengalKingdom {
	return e.旁伽罗Bengal
}
    
func (e *旁伽罗帝国BengalEmpire) KBihar毗诃罗() bihar.BiharKingdom {
	return e.毗诃罗Bihar
}
    
func (e *旁伽罗帝国BengalEmpire) KGondwana共陀婆那() gondwana.GondwanaKingdom {
	return e.共陀婆那Gondwana
}
    
func (e *旁伽罗帝国BengalEmpire) KKamarupa迦摩缕波() kamarupa.KamarupaKingdom {
	return e.迦摩缕波Kamarupa
}
    
func (e *旁伽罗帝国BengalEmpire) KOrissa乌里舍() orissa.OrissaKingdom {
	return e.乌里舍Orissa
}
    
var EBengal旁伽罗帝国 BengalEmpire = &旁伽罗帝国BengalEmpire{}

func init() {
	f := EBengal旁伽罗帝国.(*旁伽罗帝国BengalEmpire)
	f.BaseEmpire = feud.BaseEmpire{
		Title:     "bengal",
		TitleName: "旁伽罗帝国",
		TitleCode: "e_bengal",
		Kingdoms:  map[string]feud.Kingdom{},
	}
	f.旁伽罗Bengal = bengal.KBengal旁伽罗
	f.旁伽罗Bengal.SetParent(f)
	f.毗诃罗Bihar = bihar.KBihar毗诃罗
	f.毗诃罗Bihar.SetParent(f)
	f.共陀婆那Gondwana = gondwana.KGondwana共陀婆那
	f.共陀婆那Gondwana.SetParent(f)
	f.迦摩缕波Kamarupa = kamarupa.KKamarupa迦摩缕波
	f.迦摩缕波Kamarupa.SetParent(f)
	f.乌里舍Orissa = orissa.KOrissa乌里舍
	f.乌里舍Orissa.SetParent(f)
}
