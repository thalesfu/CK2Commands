package racakonda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type RacakondaCounty interface {
    feud.County
    BBhuvanagiri部婆那耆厘() 	feud.Barony
    BManchal曼遮罗() 	feud.Barony
    BManpura曼补罗() 	feud.Barony
    BMara摩罗() 	feud.Barony
    BRacakonda罗遮军荼() 	feud.Barony
    BSobhapur娑婆补罗() 	feud.Barony
    BYadagirigutta耶陀耆厘矩吒() 	feud.Barony
}

type 罗遮军荼RacakondaCounty struct {
	feud.BaseCounty
	部婆那耆厘Bhuvanagiri 	feud.Barony
	曼遮罗Manchal 	feud.Barony
	曼补罗Manpura 	feud.Barony
	摩罗Mara 	feud.Barony
	罗遮军荼Racakonda 	feud.Barony
	娑婆补罗Sobhapur 	feud.Barony
	耶陀耆厘矩吒Yadagirigutta 	feud.Barony
}

func (c *罗遮军荼RacakondaCounty) BBhuvanagiri部婆那耆厘() feud.Barony {
	return c.部婆那耆厘Bhuvanagiri
}
    
func (c *罗遮军荼RacakondaCounty) BManchal曼遮罗() feud.Barony {
	return c.曼遮罗Manchal
}
    
func (c *罗遮军荼RacakondaCounty) BManpura曼补罗() feud.Barony {
	return c.曼补罗Manpura
}
    
func (c *罗遮军荼RacakondaCounty) BMara摩罗() feud.Barony {
	return c.摩罗Mara
}
    
func (c *罗遮军荼RacakondaCounty) BRacakonda罗遮军荼() feud.Barony {
	return c.罗遮军荼Racakonda
}
    
func (c *罗遮军荼RacakondaCounty) BSobhapur娑婆补罗() feud.Barony {
	return c.娑婆补罗Sobhapur
}
    
func (c *罗遮军荼RacakondaCounty) BYadagirigutta耶陀耆厘矩吒() feud.Barony {
	return c.耶陀耆厘矩吒Yadagirigutta
}
    
var CRacakonda罗遮军荼 RacakondaCounty = &罗遮军荼RacakondaCounty{}

func init() {
	f := CRacakonda罗遮军荼.(*罗遮军荼RacakondaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1209",
		Title:     "racakonda",
		TitleName: "罗遮军荼",
		TitleCode: "c_racakonda",
		Baronies:  map[string]feud.Barony{},
	}

	f.部婆那耆厘Bhuvanagiri = BBhuvanagiri部婆那耆厘
	f.部婆那耆厘Bhuvanagiri.SetParent(f)
	
	f.曼遮罗Manchal = BManchal曼遮罗
	f.曼遮罗Manchal.SetParent(f)
	
	f.曼补罗Manpura = BManpura曼补罗
	f.曼补罗Manpura.SetParent(f)
	
	f.摩罗Mara = BMara摩罗
	f.摩罗Mara.SetParent(f)
	
	f.罗遮军荼Racakonda = BRacakonda罗遮军荼
	f.罗遮军荼Racakonda.SetParent(f)
	
	f.娑婆补罗Sobhapur = BSobhapur娑婆补罗
	f.娑婆补罗Sobhapur.SetParent(f)
	
	f.耶陀耆厘矩吒Yadagirigutta = BYadagirigutta耶陀耆厘矩吒
	f.耶陀耆厘矩吒Yadagirigutta.SetParent(f)
	
}
