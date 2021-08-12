package salon

type Action int

const (
	Haircut Action = iota
	Styling
	Shaving
	Manicure
)

func (action Action) String() string {
	switch action {
	case Haircut:
		return "Haircut"
	case Styling:
		return "Styling"
	case Shaving:
		return "Shaving"
	case Manicure:
		return "Manicure"
	}

	return ""
}