package trans-portage

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Trans-portageCounty interface {
    feud.County
    BKonosha科诺沙() 	feud.Barony
    BNyandoma尼扬多马() 	feud.Barony
    BObozerskiy奥博泽尔斯基() 	feud.Barony
    BPlesetsk普列谢茨克() 	feud.Barony
    BSamoded萨莫杰德() 	feud.Barony
    BShenkursk申库尔斯克() 	feud.Barony
    BSolovetsky索洛韦茨基() 	feud.Barony
}

type 奥涅加Trans-portageCounty struct {
	feud.BaseCounty
	科诺沙Konosha 	feud.Barony
	尼扬多马Nyandoma 	feud.Barony
	奥博泽尔斯基Obozerskiy 	feud.Barony
	普列谢茨克Plesetsk 	feud.Barony
	萨莫杰德Samoded 	feud.Barony
	申库尔斯克Shenkursk 	feud.Barony
	索洛韦茨基Solovetsky 	feud.Barony
}

func (c *奥涅加Trans-portageCounty) BKonosha科诺沙() feud.Barony {
	return c.科诺沙Konosha
}
    
func (c *奥涅加Trans-portageCounty) BNyandoma尼扬多马() feud.Barony {
	return c.尼扬多马Nyandoma
}
    
func (c *奥涅加Trans-portageCounty) BObozerskiy奥博泽尔斯基() feud.Barony {
	return c.奥博泽尔斯基Obozerskiy
}
    
func (c *奥涅加Trans-portageCounty) BPlesetsk普列谢茨克() feud.Barony {
	return c.普列谢茨克Plesetsk
}
    
func (c *奥涅加Trans-portageCounty) BSamoded萨莫杰德() feud.Barony {
	return c.萨莫杰德Samoded
}
    
func (c *奥涅加Trans-portageCounty) BShenkursk申库尔斯克() feud.Barony {
	return c.申库尔斯克Shenkursk
}
    
func (c *奥涅加Trans-portageCounty) BSolovetsky索洛韦茨基() feud.Barony {
	return c.索洛韦茨基Solovetsky
}
    
var CTrans-portage奥涅加 Trans-portageCounty = &奥涅加Trans-portageCounty{}

func init() {
	f := CTrans-portage奥涅加.(*奥涅加Trans-portageCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "394",
		Title:     "trans-portage",
		TitleName: "奥涅加",
		TitleCode: "c_trans-portage",
		Baronies:  map[string]feud.Barony{},
	}

	f.科诺沙Konosha = BKonosha科诺沙
	f.科诺沙Konosha.SetParent(f)
	
	f.尼扬多马Nyandoma = BNyandoma尼扬多马
	f.尼扬多马Nyandoma.SetParent(f)
	
	f.奥博泽尔斯基Obozerskiy = BObozerskiy奥博泽尔斯基
	f.奥博泽尔斯基Obozerskiy.SetParent(f)
	
	f.普列谢茨克Plesetsk = BPlesetsk普列谢茨克
	f.普列谢茨克Plesetsk.SetParent(f)
	
	f.萨莫杰德Samoded = BSamoded萨莫杰德
	f.萨莫杰德Samoded.SetParent(f)
	
	f.申库尔斯克Shenkursk = BShenkursk申库尔斯克
	f.申库尔斯克Shenkursk.SetParent(f)
	
	f.索洛韦茨基Solovetsky = BSolovetsky索洛韦茨基
	f.索洛韦茨基Solovetsky.SetParent(f)
	
}
