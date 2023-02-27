package chester

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ChesterCounty interface {
    feud.County
    BBeeston比斯顿() 	feud.Barony
    BChester切斯特() 	feud.Barony
    BHalton哈尔顿() 	feud.Barony
    BMacclesfield麦克尔斯菲尔德() 	feud.Barony
    BMalpas莫尔珀斯() 	feud.Barony
    BNantwich楠特威奇() 	feud.Barony
    BNorthwich诺斯维奇() 	feud.Barony
    BSandbach桑德巴奇() 	feud.Barony
}

type 切斯特ChesterCounty struct {
	feud.BaseCounty
	比斯顿Beeston 	feud.Barony
	切斯特Chester 	feud.Barony
	哈尔顿Halton 	feud.Barony
	麦克尔斯菲尔德Macclesfield 	feud.Barony
	莫尔珀斯Malpas 	feud.Barony
	楠特威奇Nantwich 	feud.Barony
	诺斯维奇Northwich 	feud.Barony
	桑德巴奇Sandbach 	feud.Barony
}

func (c *切斯特ChesterCounty) BBeeston比斯顿() feud.Barony {
	return c.比斯顿Beeston
}
    
func (c *切斯特ChesterCounty) BChester切斯特() feud.Barony {
	return c.切斯特Chester
}
    
func (c *切斯特ChesterCounty) BHalton哈尔顿() feud.Barony {
	return c.哈尔顿Halton
}
    
func (c *切斯特ChesterCounty) BMacclesfield麦克尔斯菲尔德() feud.Barony {
	return c.麦克尔斯菲尔德Macclesfield
}
    
func (c *切斯特ChesterCounty) BMalpas莫尔珀斯() feud.Barony {
	return c.莫尔珀斯Malpas
}
    
func (c *切斯特ChesterCounty) BNantwich楠特威奇() feud.Barony {
	return c.楠特威奇Nantwich
}
    
func (c *切斯特ChesterCounty) BNorthwich诺斯维奇() feud.Barony {
	return c.诺斯维奇Northwich
}
    
func (c *切斯特ChesterCounty) BSandbach桑德巴奇() feud.Barony {
	return c.桑德巴奇Sandbach
}
    
var CChester切斯特 ChesterCounty = &切斯特ChesterCounty{}

func init() {
	f := CChester切斯特.(*切斯特ChesterCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "59",
		Title:     "chester",
		TitleName: "切斯特",
		TitleCode: "c_chester",
		Baronies:  map[string]feud.Barony{},
	}

	f.比斯顿Beeston = BBeeston比斯顿
	f.比斯顿Beeston.SetParent(f)
	
	f.切斯特Chester = BChester切斯特
	f.切斯特Chester.SetParent(f)
	
	f.哈尔顿Halton = BHalton哈尔顿
	f.哈尔顿Halton.SetParent(f)
	
	f.麦克尔斯菲尔德Macclesfield = BMacclesfield麦克尔斯菲尔德
	f.麦克尔斯菲尔德Macclesfield.SetParent(f)
	
	f.莫尔珀斯Malpas = BMalpas莫尔珀斯
	f.莫尔珀斯Malpas.SetParent(f)
	
	f.楠特威奇Nantwich = BNantwich楠特威奇
	f.楠特威奇Nantwich.SetParent(f)
	
	f.诺斯维奇Northwich = BNorthwich诺斯维奇
	f.诺斯维奇Northwich.SetParent(f)
	
	f.桑德巴奇Sandbach = BSandbach桑德巴奇
	f.桑德巴奇Sandbach.SetParent(f)
	
}
