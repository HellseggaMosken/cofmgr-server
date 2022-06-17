package serializer

import "cofmgr/model"

type Conference struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	StartDate   int64  `json:"startDate"`
	EndDate     int64  `json:"endDate"`
	URL         string `json:"url"`
	Location    string `json:"location"`
	Description string `json:"description"`
}

func BuildConferencere(cof *model.Conference) *Conference {
	if cof == nil {
		return nil
	}
	return &Conference{
		ID:          cof.ID,
		Name:        cof.Name,
		StartDate:   cof.StartDate,
		EndDate:     cof.EndDate,
		URL:         cof.URL,
		Location:    cof.Location,
		Description: cof.Description,
	}
}

func BuildConferenceres(cofs []model.Conference) []*Conference {
	res := make([]*Conference, len(cofs))
	for i := range cofs {
		res[i] = BuildConferencere(&cofs[i])
	}
	return res
}
