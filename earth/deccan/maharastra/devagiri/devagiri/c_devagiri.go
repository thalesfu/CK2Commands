package devagiri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DevagiriCounty interface {
	feud.County
	BBangra旁揭罗() feud.Barony
	BDevagiri提婆耆厘() feud.Barony
	BEllora医罗补罗() feud.Barony
	BGrishneshwar姞利瑟泥湿伐罗() feud.Barony
	BKhuldabad库尔达巴德() feud.Barony
	BSillod锡洛德() feud.Barony
	BVaijapur吠阇补罗() feud.Barony
}

type 提婆耆厘DevagiriCounty struct {
	feud.BaseCounty
	旁揭罗Bangra           feud.Barony
	提婆耆厘Devagiri        feud.Barony
	医罗补罗Ellora          feud.Barony
	姞利瑟泥湿伐罗Grishneshwar feud.Barony
	库尔达巴德Khuldabad      feud.Barony
	锡洛德Sillod           feud.Barony
	吠阇补罗Vaijapur        feud.Barony
}

func (c *提婆耆厘DevagiriCounty) BBangra旁揭罗() feud.Barony {
	return c.旁揭罗Bangra
}

func (c *提婆耆厘DevagiriCounty) BDevagiri提婆耆厘() feud.Barony {
	return c.提婆耆厘Devagiri
}

func (c *提婆耆厘DevagiriCounty) BEllora医罗补罗() feud.Barony {
	return c.医罗补罗Ellora
}

func (c *提婆耆厘DevagiriCounty) BGrishneshwar姞利瑟泥湿伐罗() feud.Barony {
	return c.姞利瑟泥湿伐罗Grishneshwar
}

func (c *提婆耆厘DevagiriCounty) BKhuldabad库尔达巴德() feud.Barony {
	return c.库尔达巴德Khuldabad
}

func (c *提婆耆厘DevagiriCounty) BSillod锡洛德() feud.Barony {
	return c.锡洛德Sillod
}

func (c *提婆耆厘DevagiriCounty) BVaijapur吠阇补罗() feud.Barony {
	return c.吠阇补罗Vaijapur
}

var CDevagiri提婆耆厘 DevagiriCounty = &提婆耆厘DevagiriCounty{}

func init() {
	f := CDevagiri提婆耆厘.(*提婆耆厘DevagiriCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1145",
		Title:     "devagiri",
		TitleName: "提婆耆厘",
		TitleCode: "c_devagiri",
		Baronies:  map[string]feud.Barony{},
	}

	f.旁揭罗Bangra = BBangra旁揭罗
	f.旁揭罗Bangra.SetParent(f)

	f.提婆耆厘Devagiri = BDevagiri提婆耆厘
	f.提婆耆厘Devagiri.SetParent(f)

	f.医罗补罗Ellora = BEllora医罗补罗
	f.医罗补罗Ellora.SetParent(f)

	f.姞利瑟泥湿伐罗Grishneshwar = BGrishneshwar姞利瑟泥湿伐罗
	f.姞利瑟泥湿伐罗Grishneshwar.SetParent(f)

	f.库尔达巴德Khuldabad = BKhuldabad库尔达巴德
	f.库尔达巴德Khuldabad.SetParent(f)

	f.锡洛德Sillod = BSillod锡洛德
	f.锡洛德Sillod.SetParent(f)

	f.吠阇补罗Vaijapur = BVaijapur吠阇补罗
	f.吠阇补罗Vaijapur.SetParent(f)

}
