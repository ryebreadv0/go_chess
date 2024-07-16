package pieces

import (
	"chess/utils"
)

// var abs = utils.Abs
type Vec2 = utils.Vec2

// takes in a delta and returns if the move is valid
func (p Piece) ValidMove(move Vec2) bool {
	switch p.PieceType {
	case PAWN:
		{
			if p.Color == BLACK {
				if move.Y == 1 && move.X == 0 {
					return true
				} else if p.FirstMove && move.Y == 2 && move.X == 0 {
					return true
				} else if move.Y == 1 && (move.X == -1 || move.X == 1) {
					return true // eat case
				}
				return false
			} else {
				if move.Y == -1 && move.X == 0 {
					return true
				} else if p.FirstMove && move.Y == -2 && move.X == 0 {
					return true
				} else if move.Y == -1 && (move.X == -1 || move.X == 1) {
					return true // eat case
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
			return utils.Abs(move.X) == utils.Abs(move.Y) && move.X != 0
		}
	case QUEEN:
		{
			if move.X == 0 && move.Y != 0 {
				return true
			}
			if move.X != 0 && move.Y == 0 {
				return true
			}
			return utils.Abs(move.X) == utils.Abs(move.Y) && move.X != 0
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

func appendValidMove(validMoves *[]Vec2, pos Vec2, move Vec2) {
	*validMoves = append(*validMoves, pos.Add(move))
}

func (p Piece) ListValidMoves(pos Vec2) []Vec2 {
	validMoves := []Vec2{}

	switch p.PieceType {
	case PAWN:
		{
			if p.Color == BLACK {
				appendValidMove(&validMoves, pos, Vec2{X: 0, Y: 1})
				if p.FirstMove {
					appendValidMove(&validMoves, pos, Vec2{X: 0, Y: 2})
				}
				appendValidMove(&validMoves, pos, Vec2{X: -1, Y: 1})
				appendValidMove(&validMoves, pos, Vec2{X: 1, Y: 1})
			} else {
				appendValidMove(&validMoves, pos, Vec2{X: 0, Y: -1})
				if p.FirstMove {
					appendValidMove(&validMoves, pos, Vec2{X: 0, Y: -2})
				}
				appendValidMove(&validMoves, pos, Vec2{X: -1, Y: -1})
				appendValidMove(&validMoves, pos, Vec2{X: 1, Y: -1})
			}
			break
		}
	case ROOK:
		{
			for i := 1; i < 8; i++ {
				appendValidMove(&validMoves, pos, Vec2{X: 0, Y: i})
				appendValidMove(&validMoves, pos, Vec2{X: 0, Y: -i})
				appendValidMove(&validMoves, pos, Vec2{X: i, Y: 0})
				appendValidMove(&validMoves, pos, Vec2{X: -i, Y: 0})
			}
			break
		}
	case KNIGHT:
		{
			appendValidMove(&validMoves, pos, Vec2{X: 1, Y: 2})
			appendValidMove(&validMoves, pos, Vec2{X: 1, Y: -2})
			appendValidMove(&validMoves, pos, Vec2{X: 2, Y: 1})
			appendValidMove(&validMoves, pos, Vec2{X: 2, Y: -1})
			appendValidMove(&validMoves, pos, Vec2{X: -1, Y: 2})
			appendValidMove(&validMoves, pos, Vec2{X: -1, Y: -2})
			appendValidMove(&validMoves, pos, Vec2{X: -2, Y: 1})
			appendValidMove(&validMoves, pos, Vec2{X: -2, Y: -1})
			break
		}
	case BISHOP:
		{
			for i := 1; i < 8; i++ {
				appendValidMove(&validMoves, pos, Vec2{X: i, Y: i})
				appendValidMove(&validMoves, pos, Vec2{X: i, Y: -i})
				appendValidMove(&validMoves, pos, Vec2{X: -i, Y: i})
				appendValidMove(&validMoves, pos, Vec2{X: -i, Y: -i})
			}
			break
		}
	case QUEEN:
		{
			for i := 1; i < 8; i++ {
				appendValidMove(&validMoves, pos, Vec2{X: 0, Y: i})
				appendValidMove(&validMoves, pos, Vec2{X: 0, Y: -i})
				appendValidMove(&validMoves, pos, Vec2{X: i, Y: 0})
				appendValidMove(&validMoves, pos, Vec2{X: -i, Y: 0})
				appendValidMove(&validMoves, pos, Vec2{X: i, Y: i})
				appendValidMove(&validMoves, pos, Vec2{X: i, Y: -i})
				appendValidMove(&validMoves, pos, Vec2{X: -i, Y: i})
				appendValidMove(&validMoves, pos, Vec2{X: -i, Y: -i})
			}
			break
		}
	case KING:
		{
			appendValidMove(&validMoves, pos, Vec2{X: 0, Y: 1})
			appendValidMove(&validMoves, pos, Vec2{X: 0, Y: -1})
			appendValidMove(&validMoves, pos, Vec2{X: 1, Y: 0})
			appendValidMove(&validMoves, pos, Vec2{X: 1, Y: 1})
			appendValidMove(&validMoves, pos, Vec2{X: 1, Y: -1})
			appendValidMove(&validMoves, pos, Vec2{X: -1, Y: 0})
			appendValidMove(&validMoves, pos, Vec2{X: -1, Y: 1})
			appendValidMove(&validMoves, pos, Vec2{X: -1, Y: -1})
			break
		}

	} // end switch

	// fmt.Println("Valid moves: ", validMoves)
	return validMoves
}
