package torres

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TorresCounty interface {
	feud.County
	BAlghero阿尔盖罗() feud.Barony
	BArdara阿尔达拉() feud.Barony
	BBosa博萨() feud.Barony
	BOschiri奥斯基里() feud.Barony
	BOttana奥塔纳() feud.Barony
	BPortotorres托雷斯港() feud.Barony
	BSassari萨萨里() feud.Barony
	BValledoria瓦莱多里亚() feud.Barony
}

type 托雷斯TorresCounty struct {
	feud.BaseCounty
	阿尔盖罗Alghero     feud.Barony
	阿尔达拉Ardara      feud.Barony
	博萨Bosa          feud.Barony
	奥斯基里Oschiri     feud.Barony
	奥塔纳Ottana       feud.Barony
	托雷斯港Portotorres feud.Barony
	萨萨里Sassari      feud.Barony
	瓦莱多里亚Valledoria feud.Barony
}

func (c *托雷斯TorresCounty) BAlghero阿尔盖罗() feud.Barony {
	return c.阿尔盖罗Alghero
}

func (c *托雷斯TorresCounty) BArdara阿尔达拉() feud.Barony {
	return c.阿尔达拉Ardara
}

func (c *托雷斯TorresCounty) BBosa博萨() feud.Barony {
	return c.博萨Bosa
}

func (c *托雷斯TorresCounty) BOschiri奥斯基里() feud.Barony {
	return c.奥斯基里Oschiri
}

func (c *托雷斯TorresCounty) BOttana奥塔纳() feud.Barony {
	return c.奥塔纳Ottana
}

func (c *托雷斯TorresCounty) BPortotorres托雷斯港() feud.Barony {
	return c.托雷斯港Portotorres
}

func (c *托雷斯TorresCounty) BSassari萨萨里() feud.Barony {
	return c.萨萨里Sassari
}

func (c *托雷斯TorresCounty) BValledoria瓦莱多里亚() feud.Barony {
	return c.瓦莱多里亚Valledoria
}

var CTorres托雷斯 TorresCounty = &托雷斯TorresCounty{}

func init() {
	f := CTorres托雷斯.(*托雷斯TorresCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1576",
		Title:     "torres",
		TitleName: "托雷斯",
		TitleCode: "c_torres",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔盖罗Alghero = BAlghero阿尔盖罗
	f.阿尔盖罗Alghero.SetParent(f)

	f.阿尔达拉Ardara = BArdara阿尔达拉
	f.阿尔达拉Ardara.SetParent(f)

	f.博萨Bosa = BBosa博萨
	f.博萨Bosa.SetParent(f)

	f.奥斯基里Oschiri = BOschiri奥斯基里
	f.奥斯基里Oschiri.SetParent(f)

	f.奥塔纳Ottana = BOttana奥塔纳
	f.奥塔纳Ottana.SetParent(f)

	f.托雷斯港Portotorres = BPortotorres托雷斯港
	f.托雷斯港Portotorres.SetParent(f)

	f.萨萨里Sassari = BSassari萨萨里
	f.萨萨里Sassari.SetParent(f)

	f.瓦莱多里亚Valledoria = BValledoria瓦莱多里亚
	f.瓦莱多里亚Valledoria.SetParent(f)

}
