package sarangpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SarangpurCounty interface {
    feud.County
    BAshta阿什多() 	feud.Barony
    BBhoilymbong菩梨耶佛俱() 	feud.Barony
    BGagraun伽罗恩() 	feud.Barony
    BJhalawar阇罗婆罗() 	feud.Barony
    BRajgarh曷罗阇姞利呬() 	feud.Barony
    BSarangpur萨浪伽补罗() 	feud.Barony
    BSidhhapur悉陀补罗() 	feud.Barony
}

type 萨浪伽补罗SarangpurCounty struct {
	feud.BaseCounty
	阿什多Ashta 	feud.Barony
	菩梨耶佛俱Bhoilymbong 	feud.Barony
	伽罗恩Gagraun 	feud.Barony
	阇罗婆罗Jhalawar 	feud.Barony
	曷罗阇姞利呬Rajgarh 	feud.Barony
	萨浪伽补罗Sarangpur 	feud.Barony
	悉陀补罗Sidhhapur 	feud.Barony
}

func (c *萨浪伽补罗SarangpurCounty) BAshta阿什多() feud.Barony {
	return c.阿什多Ashta
}
    
func (c *萨浪伽补罗SarangpurCounty) BBhoilymbong菩梨耶佛俱() feud.Barony {
	return c.菩梨耶佛俱Bhoilymbong
}
    
func (c *萨浪伽补罗SarangpurCounty) BGagraun伽罗恩() feud.Barony {
	return c.伽罗恩Gagraun
}
    
func (c *萨浪伽补罗SarangpurCounty) BJhalawar阇罗婆罗() feud.Barony {
	return c.阇罗婆罗Jhalawar
}
    
func (c *萨浪伽补罗SarangpurCounty) BRajgarh曷罗阇姞利呬() feud.Barony {
	return c.曷罗阇姞利呬Rajgarh
}
    
func (c *萨浪伽补罗SarangpurCounty) BSarangpur萨浪伽补罗() feud.Barony {
	return c.萨浪伽补罗Sarangpur
}
    
func (c *萨浪伽补罗SarangpurCounty) BSidhhapur悉陀补罗() feud.Barony {
	return c.悉陀补罗Sidhhapur
}
    
var CSarangpur萨浪伽补罗 SarangpurCounty = &萨浪伽补罗SarangpurCounty{}

func init() {
	f := CSarangpur萨浪伽补罗.(*萨浪伽补罗SarangpurCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1150",
		Title:     "sarangpur",
		TitleName: "萨浪伽补罗",
		TitleCode: "c_sarangpur",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿什多Ashta = BAshta阿什多
	f.阿什多Ashta.SetParent(f)
	
	f.菩梨耶佛俱Bhoilymbong = BBhoilymbong菩梨耶佛俱
	f.菩梨耶佛俱Bhoilymbong.SetParent(f)
	
	f.伽罗恩Gagraun = BGagraun伽罗恩
	f.伽罗恩Gagraun.SetParent(f)
	
	f.阇罗婆罗Jhalawar = BJhalawar阇罗婆罗
	f.阇罗婆罗Jhalawar.SetParent(f)
	
	f.曷罗阇姞利呬Rajgarh = BRajgarh曷罗阇姞利呬
	f.曷罗阇姞利呬Rajgarh.SetParent(f)
	
	f.萨浪伽补罗Sarangpur = BSarangpur萨浪伽补罗
	f.萨浪伽补罗Sarangpur.SetParent(f)
	
	f.悉陀补罗Sidhhapur = BSidhhapur悉陀补罗
	f.悉陀补罗Sidhhapur.SetParent(f)
	
}
