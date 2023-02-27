package dadhipadra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DadhipadraCounty interface {
    feud.County
    BBamanlapalli婆曼罗波梨() 	feud.Barony
    BChampaner尚庞() 	feud.Barony
    BDadhipadra陀提波陀罗() 	feud.Barony
    BGodhra戈持罗() 	feud.Barony
    BJawad奢瓦德() 	feud.Barony
    BKalika_mata伽利伽摩多() 	feud.Barony
    BPavagadh波伐伽德() 	feud.Barony
}

type 陀提波陀罗DadhipadraCounty struct {
	feud.BaseCounty
	婆曼罗波梨Bamanlapalli 	feud.Barony
	尚庞Champaner 	feud.Barony
	陀提波陀罗Dadhipadra 	feud.Barony
	戈持罗Godhra 	feud.Barony
	奢瓦德Jawad 	feud.Barony
	伽利伽摩多Kalika_mata 	feud.Barony
	波伐伽德Pavagadh 	feud.Barony
}

func (c *陀提波陀罗DadhipadraCounty) BBamanlapalli婆曼罗波梨() feud.Barony {
	return c.婆曼罗波梨Bamanlapalli
}
    
func (c *陀提波陀罗DadhipadraCounty) BChampaner尚庞() feud.Barony {
	return c.尚庞Champaner
}
    
func (c *陀提波陀罗DadhipadraCounty) BDadhipadra陀提波陀罗() feud.Barony {
	return c.陀提波陀罗Dadhipadra
}
    
func (c *陀提波陀罗DadhipadraCounty) BGodhra戈持罗() feud.Barony {
	return c.戈持罗Godhra
}
    
func (c *陀提波陀罗DadhipadraCounty) BJawad奢瓦德() feud.Barony {
	return c.奢瓦德Jawad
}
    
func (c *陀提波陀罗DadhipadraCounty) BKalika_mata伽利伽摩多() feud.Barony {
	return c.伽利伽摩多Kalika_mata
}
    
func (c *陀提波陀罗DadhipadraCounty) BPavagadh波伐伽德() feud.Barony {
	return c.波伐伽德Pavagadh
}
    
var CDadhipadra陀提波陀罗 DadhipadraCounty = &陀提波陀罗DadhipadraCounty{}

func init() {
	f := CDadhipadra陀提波陀罗.(*陀提波陀罗DadhipadraCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1289",
		Title:     "dadhipadra",
		TitleName: "陀提波陀罗",
		TitleCode: "c_dadhipadra",
		Baronies:  map[string]feud.Barony{},
	}

	f.婆曼罗波梨Bamanlapalli = BBamanlapalli婆曼罗波梨
	f.婆曼罗波梨Bamanlapalli.SetParent(f)
	
	f.尚庞Champaner = BChampaner尚庞
	f.尚庞Champaner.SetParent(f)
	
	f.陀提波陀罗Dadhipadra = BDadhipadra陀提波陀罗
	f.陀提波陀罗Dadhipadra.SetParent(f)
	
	f.戈持罗Godhra = BGodhra戈持罗
	f.戈持罗Godhra.SetParent(f)
	
	f.奢瓦德Jawad = BJawad奢瓦德
	f.奢瓦德Jawad.SetParent(f)
	
	f.伽利伽摩多Kalika_mata = BKalika_mata伽利伽摩多
	f.伽利伽摩多Kalika_mata.SetParent(f)
	
	f.波伐伽德Pavagadh = BPavagadh波伐伽德
	f.波伐伽德Pavagadh.SetParent(f)
	
}
