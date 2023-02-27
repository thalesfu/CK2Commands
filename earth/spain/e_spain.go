package spain

import (
	"github.com/thalesfu/CK2Commands/earth/spain/andalusia"
	"github.com/thalesfu/CK2Commands/earth/spain/aragon"
	"github.com/thalesfu/CK2Commands/earth/spain/asturias"
	"github.com/thalesfu/CK2Commands/earth/spain/badajoz"
	"github.com/thalesfu/CK2Commands/earth/spain/castille"
	"github.com/thalesfu/CK2Commands/earth/spain/leon"
	"github.com/thalesfu/CK2Commands/earth/spain/navarra"
	"github.com/thalesfu/CK2Commands/earth/spain/portugal"
	"github.com/thalesfu/CK2Commands/earth/spain/spanish_galicia"
	"github.com/thalesfu/CK2Commands/earth/spain/valencia"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SpainEmpire interface {
    feud.Empire
    KAndalusia安达卢西亚() 	andalusia.AndalusiaKingdom
    KAragon阿拉贡() 	aragon.AragonKingdom
    KAsturias阿斯图里亚斯() 	asturias.AsturiasKingdom
    KBadajoz巴达霍斯() 	badajoz.BadajozKingdom
    KCastille卡斯蒂利亚() 	castille.CastilleKingdom
    KLeon莱昂() 	leon.LeonKingdom
    KNavarra纳瓦拉() 	navarra.NavarraKingdom
    KPortugal葡萄牙() 	portugal.PortugalKingdom
    KSpanish_galicia加利西亚() 	spanish_galicia.Spanish_galiciaKingdom
    KValencia瓦伦西亚() 	valencia.ValenciaKingdom
}

type 西班牙SpainEmpire struct {
	feud.BaseEmpire
	安达卢西亚Andalusia 	andalusia.AndalusiaKingdom
	阿拉贡Aragon 	aragon.AragonKingdom
	阿斯图里亚斯Asturias 	asturias.AsturiasKingdom
	巴达霍斯Badajoz 	badajoz.BadajozKingdom
	卡斯蒂利亚Castille 	castille.CastilleKingdom
	莱昂Leon 	leon.LeonKingdom
	纳瓦拉Navarra 	navarra.NavarraKingdom
	葡萄牙Portugal 	portugal.PortugalKingdom
	加利西亚Spanish_galicia 	spanish_galicia.Spanish_galiciaKingdom
	瓦伦西亚Valencia 	valencia.ValenciaKingdom
}

func (e *西班牙SpainEmpire) KAndalusia安达卢西亚() andalusia.AndalusiaKingdom {
	return e.安达卢西亚Andalusia
}
    
func (e *西班牙SpainEmpire) KAragon阿拉贡() aragon.AragonKingdom {
	return e.阿拉贡Aragon
}
    
func (e *西班牙SpainEmpire) KAsturias阿斯图里亚斯() asturias.AsturiasKingdom {
	return e.阿斯图里亚斯Asturias
}
    
func (e *西班牙SpainEmpire) KBadajoz巴达霍斯() badajoz.BadajozKingdom {
	return e.巴达霍斯Badajoz
}
    
func (e *西班牙SpainEmpire) KCastille卡斯蒂利亚() castille.CastilleKingdom {
	return e.卡斯蒂利亚Castille
}
    
func (e *西班牙SpainEmpire) KLeon莱昂() leon.LeonKingdom {
	return e.莱昂Leon
}
    
func (e *西班牙SpainEmpire) KNavarra纳瓦拉() navarra.NavarraKingdom {
	return e.纳瓦拉Navarra
}
    
func (e *西班牙SpainEmpire) KPortugal葡萄牙() portugal.PortugalKingdom {
	return e.葡萄牙Portugal
}
    
func (e *西班牙SpainEmpire) KSpanish_galicia加利西亚() spanish_galicia.Spanish_galiciaKingdom {
	return e.加利西亚Spanish_galicia
}
    
func (e *西班牙SpainEmpire) KValencia瓦伦西亚() valencia.ValenciaKingdom {
	return e.瓦伦西亚Valencia
}
    
var ESpain西班牙 SpainEmpire = &西班牙SpainEmpire{}

func init() {
	f := ESpain西班牙.(*西班牙SpainEmpire)
	f.BaseEmpire = feud.BaseEmpire{
		Title:     "spain",
		TitleName: "西班牙",
		TitleCode: "e_spain",
		Kingdoms:  map[string]feud.Kingdom{},
	}
	f.安达卢西亚Andalusia = andalusia.KAndalusia安达卢西亚
	f.安达卢西亚Andalusia.SetParent(f)
	f.阿拉贡Aragon = aragon.KAragon阿拉贡
	f.阿拉贡Aragon.SetParent(f)
	f.阿斯图里亚斯Asturias = asturias.KAsturias阿斯图里亚斯
	f.阿斯图里亚斯Asturias.SetParent(f)
	f.巴达霍斯Badajoz = badajoz.KBadajoz巴达霍斯
	f.巴达霍斯Badajoz.SetParent(f)
	f.卡斯蒂利亚Castille = castille.KCastille卡斯蒂利亚
	f.卡斯蒂利亚Castille.SetParent(f)
	f.莱昂Leon = leon.KLeon莱昂
	f.莱昂Leon.SetParent(f)
	f.纳瓦拉Navarra = navarra.KNavarra纳瓦拉
	f.纳瓦拉Navarra.SetParent(f)
	f.葡萄牙Portugal = portugal.KPortugal葡萄牙
	f.葡萄牙Portugal.SetParent(f)
	f.加利西亚Spanish_galicia = spanish_galicia.KSpanish_galicia加利西亚
	f.加利西亚Spanish_galicia.SetParent(f)
	f.瓦伦西亚Valencia = valencia.KValencia瓦伦西亚
	f.瓦伦西亚Valencia.SetParent(f)
}
