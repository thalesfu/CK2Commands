package gandhara

import (
	"github.com/thalesfu/CK2Commands/earth/rajastan/punjab/gandhara/bannu"
	"github.com/thalesfu/CK2Commands/earth/rajastan/punjab/gandhara/nandana"
	"github.com/thalesfu/CK2Commands/earth/rajastan/punjab/gandhara/purushapura"
	"github.com/thalesfu/CK2Commands/earth/rajastan/punjab/gandhara/udabhanda"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GandharaDuke interface {
    feud.Duke
    CBannu般奴() 	bannu.BannuCounty
    CNandana欢喜园() 	nandana.NandanaCounty
    CPurushapura布路沙布逻() 	purushapura.PurushapuraCounty
    CUdabhanda乌铎迦汉荼() 	udabhanda.UdabhandaCounty
}

type 犍陀罗GandharaDuke struct {
	feud.BaseDuke
	般奴Bannu 	bannu.BannuCounty
	欢喜园Nandana 	nandana.NandanaCounty
	布路沙布逻Purushapura 	purushapura.PurushapuraCounty
	乌铎迦汉荼Udabhanda 	udabhanda.UdabhandaCounty
}

func (d *犍陀罗GandharaDuke) CBannu般奴() bannu.BannuCounty {
	return d.般奴Bannu
}
    
func (d *犍陀罗GandharaDuke) CNandana欢喜园() nandana.NandanaCounty {
	return d.欢喜园Nandana
}
    
func (d *犍陀罗GandharaDuke) CPurushapura布路沙布逻() purushapura.PurushapuraCounty {
	return d.布路沙布逻Purushapura
}
    
func (d *犍陀罗GandharaDuke) CUdabhanda乌铎迦汉荼() udabhanda.UdabhandaCounty {
	return d.乌铎迦汉荼Udabhanda
}
    
var DGandhara犍陀罗 GandharaDuke = &犍陀罗GandharaDuke{}

func init() {
	f := DGandhara犍陀罗.(*犍陀罗GandharaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "gandhara",
		TitleName: "犍陀罗",
		TitleCode: "d_gandhara",
		Counties:  map[string]feud.County{},
	}

	f.般奴Bannu = bannu.CBannu般奴
	f.般奴Bannu.SetParent(f)
	
	f.欢喜园Nandana = nandana.CNandana欢喜园
	f.欢喜园Nandana.SetParent(f)
	
	f.布路沙布逻Purushapura = purushapura.CPurushapura布路沙布逻
	f.布路沙布逻Purushapura.SetParent(f)
	
	f.乌铎迦汉荼Udabhanda = udabhanda.CUdabhanda乌铎迦汉荼
	f.乌铎迦汉荼Udabhanda.SetParent(f)
	
}
