package helper

type Meta struct {
	CurrentPage int64 `json:"current_page"`
	PerPage     int64 `json:"per_page"`
	From        int64 `json:"from"`
	To          int64 `json:"to"`
	Total       int64 `json:"total"`
	LastPage    int64 `json:"last_page"`
}

const (
	defaultLimit    = 10
	maxLimit        = 100
	defaultLastPage = 0
)

func SetPaginationParameter(page, limit int64) (int64, int64) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 || limit > maxLimit {
		limit = defaultLimit
	}

	return page, limit
}

func (m *Meta) SetPaginationData(page, limit, total int64) {
	m.CurrentPage = page
	m.PerPage = limit
	m.Total = total

	if total > 0 {
		m.LastPage = total / limit
		if total%int64(limit) != 0 {
			m.LastPage += 1
		}

	} else {
		m.LastPage = defaultLastPage
	}

	if page <= m.LastPage {
		m.From = limit*(page-1) + 1
		m.To = m.From + limit - 1
		if m.CurrentPage == m.LastPage && total < m.To {
			m.To -= limit - total%limit
		}
	}
}
