// Alfred Script Filter Json Format: https://www.alfredapp.com/help/workflows/inputs/script-filter/json/

package Alfred

type Icon struct {
	Type string `json:"type"`
	Path string `json:"path"`
}

func NewIcon(ty, path string) Icon {
	return Icon{Type: ty, Path: path}
}

func NewDefaultIcon() Icon {
	return NewIcon("fileicon", "~/Desktop")
}