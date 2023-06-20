package pieces

import (
	"chess/utils"
)

var abs = utils.Abs
type Vec2 = utils.Vec2

func (p Piece) ValidMove(move Vec2) bool {
	switch p.PieceType {
	case PAWN:
	{
		if p.Color == BLACK {
			if move.Y == 1 && move.X == 0 {
				return true
			} else if p.FirstMove && move.Y == 2 && move.X == 0 {
				return true
			}
			return false
		} else {
			if move.Y == -1 && move.X == 0 {
				return true
			} else if p.FirstMove && move.Y == -2 && move.X == 0 {
				return true
			}
			return false
		}
	}
	case ROOK:
	{
		if move.X == 0 && move.Y != 0 {
			return true
		}
		if move.X != 0 && move.Y == 0 {
			return true
		}
		return false
	}
	case KNIGHT:
	{
		if move.X == 1 && move.Y == 2 {
			return true
		}
		if move.X == 2 && move.Y == 1 {
			return true
		}
		if move.X == 2 && move.Y == -1 {
			return true
		}
		if move.X == 1 && move.Y == -2 {
			return true
		}
		if move.X == -1 && move.Y == -2 {
			return true
		}
		if move.X == -2 && move.Y == -1 {
			return true
		}
		if move.X == -2 && move.Y == 1 {
			return true
		}
		if move.X == -1 && move.Y == 2 {
			return true
		}
		return false
	}
	case BISHOP:
	{
		return abs(move.X) == abs(move.Y) && move.X != 0
	}
	case QUEEN:
	{
		if move.X == 0 && move.Y != 0 {
			return true
		}
		if move.X != 0 && move.Y == 0 {
			return true
		}
		return abs(move.X) == abs(move.Y) && move.X != 0
	}
	case KING:
	{
		if move.X <= 1 && move.X >= -1 && move.Y <= 1 && move.Y >= -1 && (move.X != 0 || move.Y != 0) {
			return true
		}
		return false
	}

	} // end switch

	return false
}

func (p Piece) ListValidMoves() []Vec2 {
	validMoves := []Vec2{}

	switch p.PieceType {
	case PAWN:
	{
		if p.Color == BLACK {
			validMoves = append(validMoves, Vec2{X: 0, Y: 1})
			if p.FirstMove {
				validMoves = append(validMoves, Vec2{X: 0, Y: 2})
			}
		} else {
			validMoves = append(validMoves, Vec2{X: 0, Y: -1})
			if p.FirstMove {
				validMoves = append(validMoves, Vec2{X: 0, Y: -2})
			}
		}
		break
	}
	case ROOK:
	{
		for i := 1; i < 8; i++ {
			validMoves = append(validMoves, Vec2{X: 0, Y: i})
			validMoves = append(validMoves, Vec2{X: 0, Y: -i})
			validMoves = append(validMoves, Vec2{X: i, Y: 0})
			validMoves = append(validMoves, Vec2{X: -i, Y: 0})
		}
		break
	}
	case KNIGHT:
	{
		validMoves = append(validMoves, Vec2{X: 1, Y: 2})
		validMoves = append(validMoves, Vec2{X: 1, Y: -2})
		validMoves = append(validMoves, Vec2{X: 2, Y: 1})
		validMoves = append(validMoves, Vec2{X: 2, Y: -1})
		validMoves = append(validMoves, Vec2{X: -1, Y: 2})
		validMoves = append(validMoves, Vec2{X: -1, Y: -2})
		validMoves = append(validMoves, Vec2{X: -2, Y: 1})
		validMoves = append(validMoves, Vec2{X: -2, Y: -1})
		break
	}
	case BISHOP:
	{
		for i := 1; i < 8; i++ {
			validMoves = append(validMoves, Vec2{X: i, Y: i})
			validMoves = append(validMoves, Vec2{X: i, Y: -i})
			validMoves = append(validMoves, Vec2{X: -i, Y: i})
			validMoves = append(validMoves, Vec2{X: -i, Y: -i})
		}
		break
	}
	case QUEEN:
	{
		for i := 1; i < 8; i++ {
			validMoves = append(validMoves, Vec2{X: 0, Y: i})
			validMoves = append(validMoves, Vec2{X: 0, Y: -i})
			validMoves = append(validMoves, Vec2{X: i, Y: 0})
			validMoves = append(validMoves, Vec2{X: -i, Y: 0})
			validMoves = append(validMoves, Vec2{X: i, Y: i})
			validMoves = append(validMoves, Vec2{X: i, Y: -i})
			validMoves = append(validMoves, Vec2{X: -i, Y: i})
			validMoves = append(validMoves, Vec2{X: -i, Y: -i})
		}
		break
	}
	case KING:
	{
		validMoves = append(validMoves, Vec2{X: 0, Y: 1})
		validMoves = append(validMoves, Vec2{X: 0, Y: -1})
		validMoves = append(validMoves, Vec2{X: 1, Y: 0})
		validMoves = append(validMoves, Vec2{X: 1, Y: 1})
		validMoves = append(validMoves, Vec2{X: 1, Y: -1})
		validMoves = append(validMoves, Vec2{X: -1, Y: 0})
		validMoves = append(validMoves, Vec2{X: -1, Y: 1})
		validMoves = append(validMoves, Vec2{X: -1, Y: -1})
		break
	}

	} // end switch
	
	
	return validMoves
}
