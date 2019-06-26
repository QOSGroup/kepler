package types

const (
	pageNo   int = 1
	pageSize int = 20
)

type Page struct {
	PageNo   int `json:"page_no" form:"page_no"`
	PageSize int `json:"page_size form:"page_size""`
}

func DefaultPage() *Page {
	return &Page{pageNo, pageSize}
}

func checkAndFillPage(page *Page) *Page {
	if page == nil {
		return DefaultPage()
	}
	if page.PageNo <= 0 {
		page.PageNo = pageNo
	}
	if page.PageSize <= 0 {
		page.PageSize = pageSize
	}
	return page
}

func (page *Page) Start() int {
	safePage := checkAndFillPage(page)
	return (safePage.PageNo - 1) * safePage.PageSize
}

func (page *Page) Limit() int {
	safePage := checkAndFillPage(page)
	return safePage.PageSize
}
