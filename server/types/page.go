package types

const (
	offset int = 1
	limit  int = 20
)

type Page struct {
	PageNo   int `json:"page_no" form:"offset"`
	PageSize int `json:"page_size form:"limit""`
}

func DefaultPage() *Page {
	return &Page{offset, limit}
}

func checkAndFillPage(page *Page) *Page {
	if page == nil {
		return DefaultPage()
	}
	if page.PageNo <= 0 {
		page.PageNo = offset
	}
	if page.PageSize <= 0 {
		page.PageSize = limit
	}
	return page
}

func (page *Page) Offset() int {
	safePage := checkAndFillPage(page)
	return (safePage.PageNo - 1) * safePage.PageSize
}

func (page *Page) Limit() int {
	safePage := checkAndFillPage(page)
	return safePage.PageSize
}
