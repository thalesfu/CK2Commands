package udayagiri

import (
	"github.com/thalesfu/CK2Commands/earth/deccan/andhra/udayagiri/amaravati"
	"github.com/thalesfu/CK2Commands/earth/deccan/andhra/udayagiri/nellore"
	"github.com/thalesfu/CK2Commands/earth/deccan/andhra/udayagiri/penugonda"
	"github.com/thalesfu/CK2Commands/earth/deccan/andhra/udayagiri/udayagiri"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type UdayagiriDuke interface {
    feud.Duke
    CAmaravati阿摩罗伐底() 	amaravati.AmaravatiCounty
    CNellore内卢楼() 	nellore.NelloreCounty
    CPenugonda卑奴郡荼() 	penugonda.PenugondaCounty
    CUdayagiri优陀耶耆厘() 	udayagiri.UdayagiriCounty
}

type 优陀耶耆厘UdayagiriDuke struct {
	feud.BaseDuke
	阿摩罗伐底Amaravati 	amaravati.AmaravatiCounty
	内卢楼Nellore 	nellore.NelloreCounty
	卑奴郡荼Penugonda 	penugonda.PenugondaCounty
	优陀耶耆厘Udayagiri 	udayagiri.UdayagiriCounty
}

func (d *优陀耶耆厘UdayagiriDuke) CAmaravati阿摩罗伐底() amaravati.AmaravatiCounty {
	return d.阿摩罗伐底Amaravati
}
    
func (d *优陀耶耆厘UdayagiriDuke) CNellore内卢楼() nellore.NelloreCounty {
	return d.内卢楼Nellore
}
    
func (d *优陀耶耆厘UdayagiriDuke) CPenugonda卑奴郡荼() penugonda.PenugondaCounty {
	return d.卑奴郡荼Penugonda
}
    
func (d *优陀耶耆厘UdayagiriDuke) CUdayagiri优陀耶耆厘() udayagiri.UdayagiriCounty {
	return d.优陀耶耆厘Udayagiri
}
    
var DUdayagiri优陀耶耆厘 UdayagiriDuke = &优陀耶耆厘UdayagiriDuke{}

func init() {
	f := DUdayagiri优陀耶耆厘.(*优陀耶耆厘UdayagiriDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "udayagiri",
		TitleName: "优陀耶耆厘",
		TitleCode: "d_udayagiri",
		Counties:  map[string]feud.County{},
	}

	f.阿摩罗伐底Amaravati = amaravati.CAmaravati阿摩罗伐底
	f.阿摩罗伐底Amaravati.SetParent(f)
	
	f.内卢楼Nellore = nellore.CNellore内卢楼
	f.内卢楼Nellore.SetParent(f)
	
	f.卑奴郡荼Penugonda = penugonda.CPenugonda卑奴郡荼
	f.卑奴郡荼Penugonda.SetParent(f)
	
	f.优陀耶耆厘Udayagiri = udayagiri.CUdayagiri优陀耶耆厘
	f.优陀耶耆厘Udayagiri.SetParent(f)
	
}
