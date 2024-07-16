package board

import (
	"chess/pieces"
	"chess/utils"
	"errors"
)

type Vec2 = utils.Vec2

func (b *Board) validPosition(pos Vec2) bool {
	return pos.X >= 0 && pos.X < BOARD_SIZE && pos.Y >= 0 && pos.Y < BOARD_SIZE
}

func (b *Board) GetPiece(pos Vec2) (pieces.Piece, error) {
	if b.validPosition(pos) {
		return b.Nodes[pos.Y][pos.X], nil
	}
	return pieces.NewNone(), errors.New("GetPiece called on an invalid location")
}

func (b *Board) hasCollision(piecePos Vec2, destPos Vec2) (bool, error) {
	delta := utils.GetDelta(piecePos, destPos)

	if delta.X == 0 && delta.Y == 0 {
		return true, errors.New("hasCollision called with the same position")
	}

	destPiece, err := b.GetPiece(destPos)
	if err != nil {
		return true, err
	}

	originalPiece, err := b.GetPiece(piecePos)
	if err != nil {
		return true, err
	}

	// check if the color of the piece is the same and the piece is not a none piece // early return
	if destPiece.Color == originalPiece.Color && destPiece.PieceType != pieces.NONE {
		return true, nil
	}

	if originalPiece.PieceType == pieces.PAWN {
		if delta.X == -1 || delta.X == 1 {
			if destPiece.PieceType != pieces.NONE {
				return false, nil
			} else {
				return true, nil
			}
		} else {
			if delta.Y == 2 || delta.Y == -2 {
				collisionPiece, err := b.GetPiece(Vec2{X: piecePos.X, Y: piecePos.Y + delta.Y/2})
				if err != nil {
					return true, err
				}
				if collisionPiece.PieceType != pieces.NONE {
					return true, nil
				}
			}
			if destPiece.PieceType == pieces.NONE {
				return false, nil
			} else {
				return true, nil
			}
		}
	}

	// assert that move is diagonal or straight
	if (utils.Abs(delta.X) == utils.Abs(delta.Y)) || (delta.X == 0 && delta.Y != 0) || (delta.X != 0 && delta.Y == 0) {

		// check each position in the direction of the move
		// clamp the vectors from -1 to 1
		xOffset := utils.Clamp(delta.X, -1, 1)
		yOffset := utils.Clamp(delta.Y, -1, 1)

		searchPos := piecePos
		searchPos.X += xOffset
		searchPos.Y += yOffset

		for searchPos.X != destPos.X || searchPos.Y != destPos.Y {
			piece, err := b.GetPiece(searchPos)
			if err != nil {
				return true, err
			}

			if piece.PieceType != pieces.NONE {
				return true, nil
			}

			searchPos.X += xOffset
			searchPos.Y += yOffset
		}
	}

	return false, nil
}

func (b *Board) ValidMove(piecePos Vec2, destPos Vec2) bool {
	piece, err := b.GetPiece(piecePos)
	if err != nil {
		return false
	}

	if piece.PieceType == pieces.NONE {
		return false
	}

	if piece.ValidMove(utils.GetDelta(piecePos, destPos)) {
		result, err := b.hasCollision(piecePos, destPos)

		if err != nil {
			// fmt.Println(err)
			return false
		}
		if !result {

			// valid move if it does not put the current teams king in check
			// temp node object
			tempNode := b.Nodes[destPos.Y][destPos.X]
			b.Nodes[destPos.Y][destPos.X] = piece
			b.Nodes[piecePos.Y][piecePos.X] = pieces.NewNone()

			// // if the tempnode is a king, then InCheck returns an error that it cannot find the king
			// if tempNode.PieceType == pieces.KING {
			// 	b.Nodes[destPos.Y][destPos.X] = tempNode
			// 	b.Nodes[piecePos.Y][piecePos.X] = piece
			// 	return false
			// }

			// check if the move puts the current teams king in check
			if b.InCheck(piece.Color) {
				// revert the move
				b.Nodes[piecePos.Y][piecePos.X] = piece
				b.Nodes[destPos.Y][destPos.X] = tempNode
				return false
			} else {
				// revert the move
				b.Nodes[piecePos.Y][piecePos.X] = piece
				b.Nodes[destPos.Y][destPos.X] = tempNode
			}

			return true
		}
	}

	return false
}

func (b *Board) ListValidMoves(boardPos Vec2) []Vec2 {
	piece, err := b.GetPiece(boardPos)
	if err != nil {
		panic(err)
	}

	moves := piece.ListValidMoves(boardPos)
	var validMoves []Vec2

	for _, move := range moves {
		if b.ValidMove(boardPos, move) {
			validMoves = append(validMoves, move)
		}
	}
	// fmt.Println(validMoves)

	return validMoves
}

func (b *Board) MovePiece(boardPos Vec2, destPos Vec2) error {
	if b.ValidMove(boardPos, destPos) {
		piece, err := b.GetPiece(boardPos)
		if err != nil {
			return err
		}

		if piece.PieceType == pieces.PAWN {
			piece.FirstMove = false
		}

		b.Nodes[destPos.Y][destPos.X] = piece
		b.Nodes[boardPos.Y][boardPos.X] = pieces.NewNone()

		b.Turn = (b.Turn + 1) % 2

		return nil
	}
	return errors.New("invalid Move")
}

func (b *Board) ValidSelection(boardPos Vec2) bool {
	piece, err := b.GetPiece(boardPos)
	if err != nil {
		return false
	}
	if piece.PieceType == pieces.NONE {
		return false
	}
	if piece.Color != b.Turn {
		return false
	}
	return true
}

func (b *Board) getKingPosition(color int) (Vec2, error) {
	for y := 0; y < BOARD_SIZE; y++ {
		for x := 0; x < BOARD_SIZE; x++ {
			piece, err := b.GetPiece(Vec2{X: x, Y: y})
			if err != nil {
				panic(err)
			}
			if piece.PieceType == pieces.KING && piece.Color == color {
				return Vec2{X: x, Y: y}, nil
			}
		}
	}
	return Vec2{X: -1, Y: -1}, errors.New("king not found")
}

func (b *Board) InCheck(color int) bool {
	kingPos, err := b.getKingPosition(color)
	if err != nil {
		panic(err)
	}

	for y := 0; y < BOARD_SIZE; y++ {
		for x := 0; x < BOARD_SIZE; x++ {
			piece, err := b.GetPiece(Vec2{X: x, Y: y})
			if err != nil {
				panic(err)
			}
			if piece.Color != b.Turn {
				if b.ValidMove(Vec2{X: x, Y: y}, kingPos) {
					return true
				}
			}
		}
	}

	return false
}

func (b *Board) IsGameOver() bool {
	// need to check every single possible move and see if it can save the king
	// if not, then it is checkmate and game is over

	var oldPiece pieces.Piece

	if b.InCheck(b.Turn) {
		for y := 0; y < BOARD_SIZE; y++ {
			for x := 0; x < BOARD_SIZE; x++ {
				piece, err := b.GetPiece(Vec2{X: x, Y: y})
				if err != nil {
					panic(err)
				}
				if piece.Color == b.Turn {
					validMoves := b.ListValidMoves(Vec2{X: x, Y: y})
					for _, move := range validMoves {
						oldPiece, err = b.GetPiece(move)
						if err != nil {
							panic(err)
						}
						err := b.MovePiece(Vec2{X: x, Y: y}, move)
						if err == nil {
							b.Nodes[move.Y][move.X] = oldPiece
							b.Nodes[y][x] = piece
							b.Turn = (b.Turn + 1) % 2
							return false
						}
					}
				}
			}
		}
		return true
	}
	return false
}
