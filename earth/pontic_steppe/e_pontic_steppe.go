package pontic_steppe

import (
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/alania"
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/crimea"
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/khazaria"
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/magyar"
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/taurica"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Pontic_steppeEmpire interface {
    feud.Empire
    KAlania阿兰尼亚() 	alania.AlaniaKingdom
    KCrimea克里米亚() 	crimea.CrimeaKingdom
    KKhazaria可萨利亚() 	khazaria.KhazariaKingdom
    KMagyar马扎尔() 	magyar.MagyarKingdom
    KTaurica陶里加() 	taurica.TauricaKingdom
}

type 黑海草原Pontic_steppeEmpire struct {
	feud.BaseEmpire
	阿兰尼亚Alania 	alania.AlaniaKingdom
	克里米亚Crimea 	crimea.CrimeaKingdom
	可萨利亚Khazaria 	khazaria.KhazariaKingdom
	马扎尔Magyar 	magyar.MagyarKingdom
	陶里加Taurica 	taurica.TauricaKingdom
}

func (e *黑海草原Pontic_steppeEmpire) KAlania阿兰尼亚() alania.AlaniaKingdom {
	return e.阿兰尼亚Alania
}
    
func (e *黑海草原Pontic_steppeEmpire) KCrimea克里米亚() crimea.CrimeaKingdom {
	return e.克里米亚Crimea
}
    
func (e *黑海草原Pontic_steppeEmpire) KKhazaria可萨利亚() khazaria.KhazariaKingdom {
	return e.可萨利亚Khazaria
}
    
func (e *黑海草原Pontic_steppeEmpire) KMagyar马扎尔() magyar.MagyarKingdom {
	return e.马扎尔Magyar
}
    
func (e *黑海草原Pontic_steppeEmpire) KTaurica陶里加() taurica.TauricaKingdom {
	return e.陶里加Taurica
}
    
var EPontic_steppe黑海草原 Pontic_steppeEmpire = &黑海草原Pontic_steppeEmpire{}

func init() {
	f := EPontic_steppe黑海草原.(*黑海草原Pontic_steppeEmpire)
	f.BaseEmpire = feud.BaseEmpire{
		Title:     "pontic_steppe",
		TitleName: "黑海草原",
		TitleCode: "e_pontic_steppe",
		Kingdoms:  map[string]feud.Kingdom{},
	}
	f.阿兰尼亚Alania = alania.KAlania阿兰尼亚
	f.阿兰尼亚Alania.SetParent(f)
	f.克里米亚Crimea = crimea.KCrimea克里米亚
	f.克里米亚Crimea.SetParent(f)
	f.可萨利亚Khazaria = khazaria.KKhazaria可萨利亚
	f.可萨利亚Khazaria.SetParent(f)
	f.马扎尔Magyar = magyar.KMagyar马扎尔
	f.马扎尔Magyar.SetParent(f)
	f.陶里加Taurica = taurica.KTaurica陶里加
	f.陶里加Taurica.SetParent(f)
}
