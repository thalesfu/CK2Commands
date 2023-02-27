package sjaelland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SjaellandCounty interface {
    feud.County
    BHelsingor赫尔辛格() 	feud.Barony
    BKalundborg凯隆堡() 	feud.Barony
    BKobenhavn哥本哈根() 	feud.Barony
    BLejre雷勒() 	feud.Barony
    BNaestved奈斯特韦兹() 	feud.Barony
    BRingsted灵斯泰兹() 	feud.Barony
    BRoskilde罗斯基勒() 	feud.Barony
    BSlagelse斯劳厄尔瑟() 	feud.Barony
    BVordingborg沃尔丁堡() 	feud.Barony
}

type 西兰SjaellandCounty struct {
	feud.BaseCounty
	赫尔辛格Helsingor 	feud.Barony
	凯隆堡Kalundborg 	feud.Barony
	哥本哈根Kobenhavn 	feud.Barony
	雷勒Lejre 	feud.Barony
	奈斯特韦兹Naestved 	feud.Barony
	灵斯泰兹Ringsted 	feud.Barony
	罗斯基勒Roskilde 	feud.Barony
	斯劳厄尔瑟Slagelse 	feud.Barony
	沃尔丁堡Vordingborg 	feud.Barony
}

func (c *西兰SjaellandCounty) BHelsingor赫尔辛格() feud.Barony {
	return c.赫尔辛格Helsingor
}
    
func (c *西兰SjaellandCounty) BKalundborg凯隆堡() feud.Barony {
	return c.凯隆堡Kalundborg
}
    
func (c *西兰SjaellandCounty) BKobenhavn哥本哈根() feud.Barony {
	return c.哥本哈根Kobenhavn
}
    
func (c *西兰SjaellandCounty) BLejre雷勒() feud.Barony {
	return c.雷勒Lejre
}
    
func (c *西兰SjaellandCounty) BNaestved奈斯特韦兹() feud.Barony {
	return c.奈斯特韦兹Naestved
}
    
func (c *西兰SjaellandCounty) BRingsted灵斯泰兹() feud.Barony {
	return c.灵斯泰兹Ringsted
}
    
func (c *西兰SjaellandCounty) BRoskilde罗斯基勒() feud.Barony {
	return c.罗斯基勒Roskilde
}
    
func (c *西兰SjaellandCounty) BSlagelse斯劳厄尔瑟() feud.Barony {
	return c.斯劳厄尔瑟Slagelse
}
    
func (c *西兰SjaellandCounty) BVordingborg沃尔丁堡() feud.Barony {
	return c.沃尔丁堡Vordingborg
}
    
var CSjaelland西兰 SjaellandCounty = &西兰SjaellandCounty{}

func init() {
	f := CSjaelland西兰.(*西兰SjaellandCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "266",
		Title:     "sjaelland",
		TitleName: "西兰",
		TitleCode: "c_sjaelland",
		Baronies:  map[string]feud.Barony{},
	}

	f.赫尔辛格Helsingor = BHelsingor赫尔辛格
	f.赫尔辛格Helsingor.SetParent(f)
	
	f.凯隆堡Kalundborg = BKalundborg凯隆堡
	f.凯隆堡Kalundborg.SetParent(f)
	
	f.哥本哈根Kobenhavn = BKobenhavn哥本哈根
	f.哥本哈根Kobenhavn.SetParent(f)
	
	f.雷勒Lejre = BLejre雷勒
	f.雷勒Lejre.SetParent(f)
	
	f.奈斯特韦兹Naestved = BNaestved奈斯特韦兹
	f.奈斯特韦兹Naestved.SetParent(f)
	
	f.灵斯泰兹Ringsted = BRingsted灵斯泰兹
	f.灵斯泰兹Ringsted.SetParent(f)
	
	f.罗斯基勒Roskilde = BRoskilde罗斯基勒
	f.罗斯基勒Roskilde.SetParent(f)
	
	f.斯劳厄尔瑟Slagelse = BSlagelse斯劳厄尔瑟
	f.斯劳厄尔瑟Slagelse.SetParent(f)
	
	f.沃尔丁堡Vordingborg = BVordingborg沃尔丁堡
	f.沃尔丁堡Vordingborg.SetParent(f)
	
}
