package serializer

import "cofmgr/model"

type Contribution struct {
	ID           string `json:"id"`
	UserID       string `json:"userId"`
	ConferenceID string `json:"conferenceId"`
	UpdateDate   int64  `json:"updateDate"`
	Title        string `json:"title"`
	Abstract     string `json:"abstract"`
	Status       int    `json:"status"`
	Comment      string `json:"comment"`
	FileName     string `json:"fileName"`
}

func BuildContribution(ctb *model.Contribution) *Contribution {
	if ctb == nil {
		return nil
	}
	return &Contribution{
		ID:           ctb.ID,
		UserID:       ctb.UserID,
		ConferenceID: ctb.ConferenceID,
		UpdateDate:   ctb.UpdateDate,
		Title:        ctb.Title,
		Abstract:     ctb.Abstract,
		Status:       int(ctb.Status),
		Comment:      ctb.Comment,
		FileName:     ctb.FileName,
	}
}

func BuildContributions(ctbs []model.Contribution) []*Contribution {
	res := make([]*Contribution, len(ctbs))
	for i := range ctbs {
		res[i] = BuildContribution(&ctbs[i])
	}
	return res
}
