package dto

type AddUniversityRequest struct {
	Name          string `json:"name"`
	Summary       string `json:"summary"`
	LocationID    int    `json:"location_id"`
	Website       string `json:"website"`
	Address       string `json:"address"`
	MinTuition    int    `json:"min_tuition"`
	MaxTuition    int    `json:"max_tuition"`
	TotalMajors   int    `json:"total_majors"`
	Logo          string `json:"logo"`
	Type          string `json:"type"`
	Accreditation string `json:"accreditation"`
}

type UpdateUniversityRequest struct {
	Name          string `json:"name"`
	Summary       string `json:"summary"`
	LocationID    int    `json:"location_id"`
	Website       string `json:"website"`
	Address       string `json:"address"`
	MinTuition    int    `json:"min_tuition"`
	MaxTuition    int    `json:"max_tuition"`
	TotalMajors   int    `json:"total_majors"`
	Logo          string `json:"logo"`
	Type          string `json:"type"`
	Accreditation string `json:"accreditation"`
}

type AddReviewRequest struct {
	UserID       int    `json:"user_id"`
	UniversityID int    `json:"university_id"`
	Text         string `json:"text"`
	Stars        int    `json:"stars"`
}
