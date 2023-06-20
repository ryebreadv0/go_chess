package board

import (
	"chess/pieces"
	"chess/utils"
	"errors"
)

var abs = utils.Abs
type Vec2 = utils.Vec2

func (b Board) validPosition(pos Vec2) bool {
	return pos.X >= 0 && pos.X < BOARD_SIZE && pos.Y >= 0 && pos.Y < BOARD_SIZE
}

func (b Board) GetPiece(pos Vec2) (pieces.Piece, error) {
	if b.validPosition(pos) {
		return b.board[pos.Y][pos.X], nil
	}
	return pieces.NewNone(), errors.New("GetPiece called on an invalid location")
}

func (b Board) hasCollision(boardPos Vec2, movePos Vec2) (bool, error) {
	// assert that move is diagonal or straight
	if (abs(movePos.X) == abs(movePos.Y)) || (movePos.X == 0 && movePos.Y != 0) || (movePos.X != 0 && movePos.Y == 0) {
		destLocation := Vec2{X: boardPos.X + movePos.X, Y: boardPos.Y + movePos.Y}

		piece, err := b.GetPiece(destLocation)
		if err != nil {
			return true, err 
		}

		originalPiece, err := b.GetPiece(boardPos)
		if err != nil {
			return true, err
		}
		
		
		// check if the color of the piece is the same and the piece is not a none piece // early return
		if piece.Color == originalPiece.Color && piece.PieceType != pieces.NONE {
			return true, nil
		}

		// check each position in the direction of the move
		// clamp the vectors from -1 to 1
		xOffset := utils.Clamp(movePos.X, -1, 1)
		yOffset := utils.Clamp(movePos.Y, -1, 1)


		for boardPos.X != destLocation.X || boardPos.Y != destLocation.Y {
			boardPos.X += xOffset
			boardPos.Y += yOffset

			piece, err := b.GetPiece(destLocation)
			if err != nil {
				return true, err
			}

			if piece.PieceType != pieces.NONE {
				return true, nil
			}
		}
	}
	
	return false, nil
}

func (b Board) ValidMove(boardPos Vec2, movePos Vec2) bool {
	
	move := Vec2{X: movePos.X - boardPos.X, Y: movePos.Y - boardPos.Y}
	
	if move.X == 0 && move.Y == 0 {
		return false
	}

	piece, err := b.GetPiece(boardPos)
	if err != nil {
		return false
	}

	if piece.PieceType == pieces.NONE {
		return false
	}

	if piece.ValidMove(move) {
		result, err := b.hasCollision(boardPos, move)

		if err != nil {
			panic(err)
		}

		if result {
			return false
		}
	}
	

	return false
}

