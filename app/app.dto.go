package app

type AppDTO struct {
	ID uint `json:"id,string,omitempty"`
	Cover string `json:"cover"`
	Download string `json:"download,string"`
}

func ToApp(dto AppDTO) App  {
	return App{Cover: dto.Cover, Download:dto.Download}
}

func ToAppDTO(app App) AppDTO  {
	return AppDTO{ID: app.ID, Cover: app.Cover, Download:app.Download}
}

func ToAppDTOs(apps []App) []AppDTO {
	appdtos := make([]AppDTO, len(apps))

	for index, item := range apps {
		appdtos[index] = ToAppDTO(item)
	}

	return appdtos
}

