package yaroslavl

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type YaroslavlCounty interface {
    feud.County
    BBorisoglebsky鲍里索格列布斯基() 	feud.Barony
    BPetropavlovsky彼得罗巴甫洛夫斯基() 	feud.Barony
    BRomanov罗曼诺夫() 	feud.Barony
    BRybinsk雷宾斯克() 	feud.Barony
    BTimerevo季梅列沃() 	feud.Barony
    BYaroslavl雅罗斯拉夫尔() 	feud.Barony
}

type 雅罗斯拉夫尔YaroslavlCounty struct {
	feud.BaseCounty
	鲍里索格列布斯基Borisoglebsky 	feud.Barony
	彼得罗巴甫洛夫斯基Petropavlovsky 	feud.Barony
	罗曼诺夫Romanov 	feud.Barony
	雷宾斯克Rybinsk 	feud.Barony
	季梅列沃Timerevo 	feud.Barony
	雅罗斯拉夫尔Yaroslavl 	feud.Barony
}

func (c *雅罗斯拉夫尔YaroslavlCounty) BBorisoglebsky鲍里索格列布斯基() feud.Barony {
	return c.鲍里索格列布斯基Borisoglebsky
}
    
func (c *雅罗斯拉夫尔YaroslavlCounty) BPetropavlovsky彼得罗巴甫洛夫斯基() feud.Barony {
	return c.彼得罗巴甫洛夫斯基Petropavlovsky
}
    
func (c *雅罗斯拉夫尔YaroslavlCounty) BRomanov罗曼诺夫() feud.Barony {
	return c.罗曼诺夫Romanov
}
    
func (c *雅罗斯拉夫尔YaroslavlCounty) BRybinsk雷宾斯克() feud.Barony {
	return c.雷宾斯克Rybinsk
}
    
func (c *雅罗斯拉夫尔YaroslavlCounty) BTimerevo季梅列沃() feud.Barony {
	return c.季梅列沃Timerevo
}
    
func (c *雅罗斯拉夫尔YaroslavlCounty) BYaroslavl雅罗斯拉夫尔() feud.Barony {
	return c.雅罗斯拉夫尔Yaroslavl
}
    
var CYaroslavl雅罗斯拉夫尔 YaroslavlCounty = &雅罗斯拉夫尔YaroslavlCounty{}

func init() {
	f := CYaroslavl雅罗斯拉夫尔.(*雅罗斯拉夫尔YaroslavlCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "572",
		Title:     "yaroslavl",
		TitleName: "雅罗斯拉夫尔",
		TitleCode: "c_yaroslavl",
		Baronies:  map[string]feud.Barony{},
	}

	f.鲍里索格列布斯基Borisoglebsky = BBorisoglebsky鲍里索格列布斯基
	f.鲍里索格列布斯基Borisoglebsky.SetParent(f)
	
	f.彼得罗巴甫洛夫斯基Petropavlovsky = BPetropavlovsky彼得罗巴甫洛夫斯基
	f.彼得罗巴甫洛夫斯基Petropavlovsky.SetParent(f)
	
	f.罗曼诺夫Romanov = BRomanov罗曼诺夫
	f.罗曼诺夫Romanov.SetParent(f)
	
	f.雷宾斯克Rybinsk = BRybinsk雷宾斯克
	f.雷宾斯克Rybinsk.SetParent(f)
	
	f.季梅列沃Timerevo = BTimerevo季梅列沃
	f.季梅列沃Timerevo.SetParent(f)
	
	f.雅罗斯拉夫尔Yaroslavl = BYaroslavl雅罗斯拉夫尔
	f.雅罗斯拉夫尔Yaroslavl.SetParent(f)
	
}
