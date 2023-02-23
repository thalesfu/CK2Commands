package winchester

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type WinchesterCounty interface {
	feud.County
	BAndover安多弗() feud.Barony
	BPortchester彻斯特港() feud.Barony
	BRomsey拉姆西() feud.Barony
	BSouthampton南安普敦() feud.Barony
	BWherwell惠韦尔() feud.Barony
	BWinchester温彻斯特() feud.Barony
}

type 温彻斯特WinchesterCounty struct {
	feud.BaseCounty
	安多弗Andover      feud.Barony
	彻斯特港Portchester feud.Barony
	拉姆西Romsey       feud.Barony
	南安普敦Southampton feud.Barony
	惠韦尔Wherwell     feud.Barony
	温彻斯特Winchester  feud.Barony
}

func (c *温彻斯特WinchesterCounty) BAndover安多弗() feud.Barony {
	return c.安多弗Andover
}

func (c *温彻斯特WinchesterCounty) BPortchester彻斯特港() feud.Barony {
	return c.彻斯特港Portchester
}

func (c *温彻斯特WinchesterCounty) BRomsey拉姆西() feud.Barony {
	return c.拉姆西Romsey
}

func (c *温彻斯特WinchesterCounty) BSouthampton南安普敦() feud.Barony {
	return c.南安普敦Southampton
}

func (c *温彻斯特WinchesterCounty) BWherwell惠韦尔() feud.Barony {
	return c.惠韦尔Wherwell
}

func (c *温彻斯特WinchesterCounty) BWinchester温彻斯特() feud.Barony {
	return c.温彻斯特Winchester
}

var CWinchester温彻斯特 WinchesterCounty = &温彻斯特WinchesterCounty{}

func init() {
	f := CWinchester温彻斯特.(*温彻斯特WinchesterCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "26",
		Title:     "winchester",
		TitleName: "温彻斯特",
		TitleCode: "c_winchester",
		Baronies:  map[string]feud.Barony{},
	}

	f.安多弗Andover = BAndover安多弗
	f.安多弗Andover.SetParent(f)

	f.彻斯特港Portchester = BPortchester彻斯特港
	f.彻斯特港Portchester.SetParent(f)

	f.拉姆西Romsey = BRomsey拉姆西
	f.拉姆西Romsey.SetParent(f)

	f.南安普敦Southampton = BSouthampton南安普敦
	f.南安普敦Southampton.SetParent(f)

	f.惠韦尔Wherwell = BWherwell惠韦尔
	f.惠韦尔Wherwell.SetParent(f)

	f.温彻斯特Winchester = BWinchester温彻斯特
	f.温彻斯特Winchester.SetParent(f)

}
