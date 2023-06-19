package pieces

type Move struct {
	x int
	y int
}

func (p Piece) ValidMove(move Move) bool {
	switch p.pieceType {
	case PAWN:
	{
		if p.color == BLACK {
			if move.y == 1 && move.x == 0 {
				return true
			} else if p.firstMove && move.y == 2 && move.x == 0 {
				return true
			}
			return false
		} else {
			if move.y == -1 && move.x == 0 {
				return true
			} else if p.firstMove && move.y == -2 && move.x == 0 {
				return true
			}
			return false
		}
	}
	case ROOK:
	{
		if move.x == 0 && move.y != 0 {
			return true
		}
		if move.x != 0 && move.y == 0 {
			return true
		}
		return false
	}
	case KNIGHT:
	{
		if move.x == 1 && move.y == 2 {
			return true
		}
		if move.x == 2 && move.y == 1 {
			return true
		}
		if move.x == 2 && move.y == -1 {
			return true
		}
		if move.x == 1 && move.y == -2 {
			return true
		}
		if move.x == -1 && move.y == -2 {
			return true
		}
		if move.x == -2 && move.y == -1 {
			return true
		}
		if move.x == -2 && move.y == 1 {
			return true
		}
		if move.x == -1 && move.y == 2 {
			return true
		}
		return false
	}
	case BISHOP:
	{
		return abs(move.x) == abs(move.y) && move.x != 0
	}
	case QUEEN:
	{
		if move.x == 0 && move.y != 0 {
			return true
		}
		if move.x != 0 && move.y == 0 {
			return true
		}
		return abs(move.x) == abs(move.y) && move.x != 0
	}
	case KING:
	{
		if move.x <= 1 && move.x >= -1 && move.y <= 1 && move.y >= -1 && (move.x != 0 || move.y != 0) {
			return true
		}
		return false
	}

	} // end switch

	return false
}

func (p Piece) ListValidMoves() []Move {
	validMoves := []Move{}

	switch p.pieceType {
	case PAWN:
	{
		if p.color == BLACK {
			validMoves = append(validMoves, Move{0, 1})
			if p.firstMove {
				validMoves = append(validMoves, Move{0, 2})
			}
		} else {
			validMoves = append(validMoves, Move{0, -1})
			if p.firstMove {
				validMoves = append(validMoves, Move{0, -2})
			}
		}
		break
	}
	case ROOK:
	{
		for i := 1; i < 8; i++ {
			validMoves = append(validMoves, Move{0, i})
			validMoves = append(validMoves, Move{0, -i})
			validMoves = append(validMoves, Move{i, 0})
			validMoves = append(validMoves, Move{-i, 0})
		}
		break
	}
	case KNIGHT:
	{
		validMoves = append(validMoves, Move{1, 2})
		validMoves = append(validMoves, Move{1, -2})
		validMoves = append(validMoves, Move{2, 1})
		validMoves = append(validMoves, Move{2, -1})
		validMoves = append(validMoves, Move{-1, 2})
		validMoves = append(validMoves, Move{-1, -2})
		validMoves = append(validMoves, Move{-2, 1})
		validMoves = append(validMoves, Move{-2, -1})
		break
	}
	case BISHOP:
	{
		for i := 1; i < 8; i++ {
			validMoves = append(validMoves, Move{i, i})
			validMoves = append(validMoves, Move{i, -i})
			validMoves = append(validMoves, Move{-i, i})
			validMoves = append(validMoves, Move{-i, -i})
		}
		break
	}
	case QUEEN:
	{
		for i := 1; i < 8; i++ {
			validMoves = append(validMoves, Move{0, i})
			validMoves = append(validMoves, Move{0, -i})
			validMoves = append(validMoves, Move{i, 0})
			validMoves = append(validMoves, Move{-i, 0})
			validMoves = append(validMoves, Move{i, i})
			validMoves = append(validMoves, Move{i, -i})
			validMoves = append(validMoves, Move{-i, i})
			validMoves = append(validMoves, Move{-i, -i})
		}
		break
	}
	case KING:
	{
		validMoves = append(validMoves, Move{0, 1})
		validMoves = append(validMoves, Move{0, -1})
		validMoves = append(validMoves, Move{1, 0})
		validMoves = append(validMoves, Move{1, 1})
		validMoves = append(validMoves, Move{1, -1})
		validMoves = append(validMoves, Move{-1, 0})
		validMoves = append(validMoves, Move{-1, 1})
		validMoves = append(validMoves, Move{-1, -1})
		break
	}

	} // end switch
	
	
	return validMoves
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}