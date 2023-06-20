package pieces

const (
	BLACK = iota
	WHITE
)

const (
	NONE = iota
	PAWN
	ROOK
	KNIGHT
	BISHOP
	QUEEN
	KING
)

type Piece struct {
	Color int
	PieceType int
	FirstMove bool
}

func NewPiece(color int, pieceType int) Piece {
	return Piece{Color: color, PieceType: pieceType, FirstMove: true}
}

func NewNone() Piece {
	return NewPiece(BLACK, NONE)
}

func NewPawn(color int) Piece {
	return NewPiece(color, PAWN)
}

func NewRook(color int) Piece {
	return NewPiece(color, ROOK)
}

func NewKnight(color int) Piece {
	return NewPiece(color, KNIGHT)
}

func NewBishop(color int) Piece {
	return NewPiece(color, BISHOP)
}

func NewQueen(color int) Piece {
	return NewPiece(color, QUEEN)
}

func NewKing(color int) Piece {
	return NewPiece(color, KING)
}

func (p Piece) String() string {
	str := ""
	if p.Color == WHITE {
		str += "[W]"
	} else {
		str += "[B]"
	}
	switch p.PieceType {
	case NONE:
		str = " "
	case PAWN:
		str += "Pawn"
	case ROOK:
		str += "Rook"
	case KNIGHT:
		str += "Knight"
	case BISHOP:
		str += "Bishop"
	case QUEEN:
		str += "Queen"
	case KING:
		str += "King"
	default:
		str += "Unknown"
	}

	return str
}
